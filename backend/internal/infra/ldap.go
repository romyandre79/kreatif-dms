package infra

import (
	"fmt"
	"log"
	"runtime"
	"strings"

	"github.com/go-ldap/ldap/v3"
	"github.com/kreatif/dms-backend/internal/config"
)

type LDAPService struct {
	cfg config.Config
}

func NewLDAPService(cfg config.Config) *LDAPService {
	return &LDAPService{cfg: cfg}
}

type LDAPUser struct {
	Username string
	Email    string
	FullName string
}

func (s *LDAPService) Authenticate(username, password string) (*LDAPUser, error) {
	if !s.cfg.LDAPEnabled {
		return nil, fmt.Errorf("LDAP authentication is disabled")
	}

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

	// === REAL LDAP MODE ===
	log.Printf("[LDAPService] Authenticating user: %s", username)

	// Detailed logging for connection attempt
	log.Printf("[LDAPService] Connecting to: %s", s.cfg.LDAPURL)

	// Detect if user is trying to use LDAP over IPC (ldapi) on Windows
	if strings.HasPrefix(strings.ToLower(s.cfg.LDAPURL), "ldapi://") && runtime.GOOS == "windows" {
		log.Println("[LDAPService] ERROR: 'ldapi://' (LDAP over IPC) is NOT supported on Windows. Use 'ldap://' or 'ldaps://'.")
		return nil, fmt.Errorf("ldapi:// is not supported on Windows (IPC error)")
	}

	l, err := ldap.DialURL(s.cfg.LDAPURL)
	if err != nil {
		log.Printf("[LDAPService] Connection error: %v", err)
		return nil, fmt.Errorf("failed to connect to LDAP at %s: %v", s.cfg.LDAPURL, err)
	}
	defer l.Close()

	// First bind with a read only user (or admin)
	err = l.Bind(s.cfg.LDAPBindDN, s.cfg.LDAPBindPass)
	if err != nil {
		return nil, fmt.Errorf("failed to bind to LDAP: %v", err)
	}

	// Search for the user
	searchRequest := ldap.NewSearchRequest(
		s.cfg.LDAPBaseDN,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf(s.cfg.LDAPUserFilter, username),
		[]string{"dn", "cn", "mail", "displayName"},
		nil,
	)

	sr, err := l.Search(searchRequest)
	if err != nil {
		return nil, fmt.Errorf("LDAP search failed: %v", err)
	}

	if len(sr.Entries) != 1 {
		return nil, fmt.Errorf("user not found or multiple users found")
	}

	userDN := sr.Entries[0].DN
	userEmail := sr.Entries[0].GetAttributeValue("mail")
	userFullName := sr.Entries[0].GetAttributeValue("displayName")
	if userFullName == "" {
		userFullName = sr.Entries[0].GetAttributeValue("cn")
	}

	// Bind as the user to verify password
	err = l.Bind(userDN, password)
	if err != nil {
		return nil, fmt.Errorf("invalid LDAP credentials: %v", err)
	}

	log.Printf("[LDAPService] Successfully authenticated user: %s", username)

	return &LDAPUser{
		Username: username,
		Email:    userEmail,
		FullName: userFullName,
	}, nil
}

func (s *LDAPService) SearchUsers(query string) ([]LDAPUser, error) {
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

	// TODO: Implement real LDAP search if needed
	return nil, fmt.Errorf("not implemented for real LDAP yet")
}
