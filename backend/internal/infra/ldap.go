package infra

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"runtime"
	"strings"

	"github.com/go-ldap/ldap/v3"
	"github.com/kreatif/dms-backend/internal/config"
	"github.com/kreatif/dms-backend/internal/repository"
)

type LDAPService struct {
	cfg  config.Config
	repo repository.Querier
}

func NewLDAPService(cfg config.Config, repo repository.Querier) *LDAPService {
	return &LDAPService{cfg: cfg, repo: repo}
}

type LDAPUser struct {
	Username string
	Email    string
	FullName string
	Groups   []string
}

type ldapNodeConfig struct {
	BaseDN     string `json:"base_dn"`
	BindDN     string `json:"bind_dn"`
	BindPass   string `json:"bind_pass"`
	UserFilter string `json:"user_filter"`
	UseTLS     bool   `json:"use_tls"`
	SkipVerify bool   `json:"skip_verify"`
}

func (s *LDAPService) getLDAPConfig(ctx context.Context) (string, ldapNodeConfig, error) {
	node, err := s.repo.GetIntegrationNodeByType(ctx, "LDAP")
	if err != nil {
		return "", ldapNodeConfig{}, fmt.Errorf("LDAP node not found in database: %v", err)
	}

	if !node.IsActive.Bool {
		return "", ldapNodeConfig{}, fmt.Errorf("LDAP integration is disabled in database")
	}

	nodeCfg := ldapNodeConfig{UseTLS: false} // Default to false, especially for local LLDAP/Docker setups
	if err := json.Unmarshal(node.ConfigJson, &nodeCfg); err != nil {
		return "", ldapNodeConfig{}, fmt.Errorf("failed to parse LDAP config JSON: %v", err)
	}

	// Use environment variables as fallback if JSON fields are empty
	if nodeCfg.BaseDN == "" {
		nodeCfg.BaseDN = s.cfg.LDAPBaseDN
	}
	if nodeCfg.BindDN == "" {
		nodeCfg.BindDN = s.cfg.LDAPBindDN
	}
	if nodeCfg.BindPass == "" {
		nodeCfg.BindPass = s.cfg.LDAPBindPass
	}
	if nodeCfg.UserFilter == "" {
		nodeCfg.UserFilter = s.cfg.LDAPUserFilter
	}

	// Format URL from endpoint
	url := node.Endpoint
	if !strings.HasPrefix(url, "ldap://") && !strings.HasPrefix(url, "ldaps://") {
		url = "ldap://" + url
	}

	return url, nodeCfg, nil
}

func (s *LDAPService) Authenticate(ctx context.Context, username, password string) (*LDAPUser, error) {
	// === SIMULATION MODE ===
	if s.cfg.LDAPSimulation {
		log.Printf("[LDAPService] (SIMULATION) Authenticating user: %s", username)
		mockUsers := map[string]*LDAPUser{
			"admin": {Username: "admin", Email: "admin@kreatif.id", FullName: "System Administrator"},
			"user":  {Username: "user", Email: "user@kreatif.id", FullName: "Standard User"},
			"budi":  {Username: "budi", Email: "budi@kreatif.id", FullName: "Budi Qartono"},
			"siti":  {Username: "siti", Email: "siti@kreatif.id", FullName: "Siti Rahma"},
		}

		user, ok := mockUsers[username]
		if !ok {
			return nil, fmt.Errorf("user not found in simulation")
		}

		// In simulation, password "password" or username+"123" is always correct
		if password != "password" && password != username+"123" {
			return nil, fmt.Errorf("invalid simulation credentials")
		}

		return user, nil
	}

	// Fetch config from Database
	ldapURL, nodeCfg, err := s.getLDAPConfig(ctx)
	if err != nil {
		return nil, err
	}

	// === REAL LDAP MODE ===
	log.Printf("[LDAPService] Authenticating user: %s", username)

	// Detailed logging for connection attempt
	log.Printf("[LDAPService] Connecting to: %s", ldapURL)

	// Detect if user is trying to use LDAP over IPC (ldapi) on Windows
	if strings.HasPrefix(strings.ToLower(ldapURL), "ldapi://") && runtime.GOOS == "windows" {
		log.Println("[LDAPService] ERROR: 'ldapi://' (LDAP over IPC) is NOT supported on Windows. Use 'ldap://' or 'ldaps://'.")
		return nil, fmt.Errorf("ldapi:// is not supported on Windows (IPC error)")
	}

	// --- STEP 1: Search for User DN using Admin/Bind credentials ---
	log.Printf("[LDAPService] Step 1: Connecting to %s", ldapURL)
	l, err := ldap.DialURL(ldapURL)
	if err != nil {
		log.Printf("[LDAPService] Step 1: Connection error: %v", err)
		return nil, fmt.Errorf("failed to connect to LDAP: %v", err)
	}

	// Try StartTLS upgrade ONLY if explicitly requested
	if !strings.HasPrefix(strings.ToLower(ldapURL), "ldaps://") && nodeCfg.UseTLS {
		log.Printf("[LDAPService] Step 1: Attempting StartTLS")
		err = l.StartTLS(&tls.Config{InsecureSkipVerify: nodeCfg.SkipVerify})
		if err != nil {
			log.Printf("[LDAPService] Step 1: StartTLS failed: %v", err)
		}
	}

	// Bind as Admin/Bind user
	log.Printf("[LDAPService] Step 1: Binding as admin: %s", nodeCfg.BindDN)
	if err := l.Bind(nodeCfg.BindDN, nodeCfg.BindPass); err != nil {
		log.Printf("[LDAPService] Step 1: Admin bind failed: %v", err)
		l.Close()
		return nil, fmt.Errorf("LDAP admin bind failed: %v", err)
	}

	// Search for the user DN
	log.Printf("[LDAPService] Step 1: Searching for user: %s (Filter: %s)", username, nodeCfg.UserFilter)
	searchRequest := ldap.NewSearchRequest(
		nodeCfg.BaseDN,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf(nodeCfg.UserFilter, username),
		[]string{"dn", "cn", "mail", "displayName", "memberOf"},
		nil,
	)

	sr, err := l.Search(searchRequest)
	if err != nil {
		log.Printf("[LDAPService] Step 1: Search failed: %v", err)
		l.Close()
		return nil, fmt.Errorf("LDAP search failed: %v", err)
	}

	if len(sr.Entries) != 1 {
		log.Printf("[LDAPService] Step 1: User not found or multiple results: %d", len(sr.Entries))
		l.Close()
		return nil, fmt.Errorf("user not found or multiple users found")
	}

	entry := sr.Entries[0]
	userDN := entry.DN
	log.Printf("[LDAPService] Step 1: Found User DN: %s", userDN)
	
	userEmail := entry.GetAttributeValue("mail")
	userFullName := entry.GetAttributeValue("displayName")
	if userFullName == "" {
		userFullName = entry.GetAttributeValue("cn")
	}

	// Extract groups
	rawGroups := entry.GetAttributeValues("memberOf")
	var userGroups []string
	for _, g := range rawGroups {
		parts := strings.Split(g, ",")
		if len(parts) > 0 && strings.HasPrefix(strings.ToLower(parts[0]), "cn=") {
			groupName := parts[0][3:]
			userGroups = append(userGroups, strings.ToLower(groupName))
		}
	}

	// Close the search connection
	l.Close()

	// --- STEP 2: Verify password using a FRESH connection as the User ---
	log.Printf("[LDAPService] Step 2: Connecting again for password verification")
	l2, err := ldap.DialURL(ldapURL)
	if err != nil {
		log.Printf("[LDAPService] Step 2: Connection error: %v", err)
		return nil, fmt.Errorf("failed to connect for password verification: %v", err)
	}
	defer l2.Close()

	// Try StartTLS upgrade again ONLY if explicitly requested
	if !strings.HasPrefix(strings.ToLower(ldapURL), "ldaps://") && nodeCfg.UseTLS {
		_ = l2.StartTLS(&tls.Config{InsecureSkipVerify: nodeCfg.SkipVerify})
	}

	log.Printf("[LDAPService] Step 2: Binding as user: %s", userDN)
	if err := l2.Bind(userDN, password); err != nil {
		log.Printf("[LDAPService] Step 2: User bind failed for %s: %v", username, err)
		return nil, fmt.Errorf("invalid LDAP credentials: %v", err)
	}

	log.Printf("[LDAPService] Step 2: Authentication successful for: %s", username)

	return &LDAPUser{
		Username: username,
		Email:    userEmail,
		FullName: userFullName,
		Groups:   userGroups,
	}, nil
}

func (s *LDAPService) SearchUsers(ctx context.Context, query string) ([]LDAPUser, error) {
	if s.cfg.LDAPSimulation {
		log.Printf("[LDAPService] (SIMULATION) Searching users with query: %s", query)
		return []LDAPUser{
			{Username: "admin", Email: "admin@kreatif.id", FullName: "System Administrator"},
			{Username: "budi", Email: "budi@kreatif.id", FullName: "Budi Qartono"},
			{Username: "siti", Email: "siti@kreatif.id", FullName: "Siti Rahma"},
			{Username: "agus", Email: "agus@kreatif.id", FullName: "Agus Santoso"},
			{Username: "ani", Email: "ani@kreatif.id", FullName: "Ani Wijaya"},
		}, nil
	}

	// TODO: Implement real LDAP search using nodeCfg if needed
	return nil, fmt.Errorf("not implemented for real LDAP yet")
}
