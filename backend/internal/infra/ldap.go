package infra

import (
	"fmt"
	"log"

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

	log.Printf("[LDAPService] Authenticating user: %s", username)

	l, err := ldap.DialURL(s.cfg.LDAPURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to LDAP: %v", err)
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
