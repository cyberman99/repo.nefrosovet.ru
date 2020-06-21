package ldap

import (
	"crypto/tls"
	"errors"
	"fmt"
	"regexp"

	ldap "gopkg.in/LDAP.v2"
	"repo.nefrosovet.ru/maximus-platform/auth/api/models"
	"repo.nefrosovet.ru/maximus-platform/auth/storage"
)

type ConnectConfig struct {
	Host         string
	Port         int
	UserDN       string
	UserPassword string
	Cipher       string

	BaseDN string
	Filter string
}

type Connect struct {
	Config *ConnectConfig
	Handle *ldap.Conn
}

// GetConnect returns main auth service connection to LDAP.
func GetConnect(backend *storage.Backend) (*Connect, error) {
	config := new(ConnectConfig)
	config.Host = backend.Host
	config.Port = int(backend.Port)
	config.UserDN = backend.BindDN
	config.UserPassword = backend.Password
	config.Cipher = backend.Cipher

	config.BaseDN = backend.BaseDN
	config.Filter = backend.Filter

	var err error
	connect, err := New(config)
	if err != nil {
		return nil, err
	}

	return connect, nil
}

func New(config *ConnectConfig) (*Connect, error) {
	var l *ldap.Conn
	var err error
	switch config.Cipher {
	case models.BackendPatchLdapParamsCipherTLS:
		tlsConfig := &tls.Config{InsecureSkipVerify: true}
		l, err = ldap.DialTLS("tcp", fmt.Sprintf("%s:%d", config.Host, config.Port), tlsConfig)
	case models.BackendLdapParamsCipherNONE:
		l, err = ldap.Dial("tcp", fmt.Sprintf("%s:%d", config.Host, config.Port))
	case models.BackendPatchLdapParamsCipherSTARTTLS:
		l, err = ldap.Dial("tcp", fmt.Sprintf("%s:%d", config.Host, config.Port))
		if err != nil {

			return nil, err
		}
		err = l.StartTLS(&tls.Config{InsecureSkipVerify: true})
	default:
		err = errors.New(fmt.Sprintf("Ciper type: %s not match any case", config.Cipher))
	}
	if err != nil {
		return nil, err
	}

	err = l.Bind(config.UserDN, config.UserPassword)
	if err != nil {
		return nil, err
	}

	connect := &Connect{Config: config, Handle: l}

	return connect, err
}

// Close closes connection to LDAP
func (connect *Connect) Close() {
	go connect.Handle.Close()

	return
}

func (connect *Connect) Search(login string) (*ldap.SearchResult, error) {
	searchRequest := ldap.NewSearchRequest(
		connect.Config.BaseDN,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0, 0, false,
		connect.applyFilter(login),
		[]string{"*"}, // A list attributes to retrieve
		nil,
	)

	searchResult, err := connect.Handle.Search(searchRequest)
	if err != nil {
		if !ldap.IsErrorWithCode(err, ldap.LDAPResultNoSuchObject) {
			return nil, err
		}
	}

	return searchResult, nil
}

func (connect *Connect) SearchWithPaging(login string, pagingSize uint32) (*ldap.SearchResult, error) {
	searchRequest := ldap.NewSearchRequest(
		connect.Config.BaseDN,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0, 0, false,
		connect.applyFilter(login),
		[]string{"*"}, // A list attributes to retrieve
		nil,
	)

	searchResult, err := connect.Handle.SearchWithPaging(searchRequest, pagingSize)
	if err != nil {
		if !ldap.IsErrorWithCode(err, ldap.LDAPResultNoSuchObject) {
			return nil, err
		}
	}

	return searchResult, nil
}

// Compile login filter regexp one time
var loginRegExp = regexp.MustCompile("%u")

// applyFilter converts filter template to ready search filter
func (connect *Connect) applyFilter(login string) string {
	if connect.Config.Filter == "" {
		return ""
	}

	return loginRegExp.ReplaceAllString(connect.Config.Filter, login)
}
