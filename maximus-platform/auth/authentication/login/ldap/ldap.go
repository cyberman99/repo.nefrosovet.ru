package ldap

import (
	"errors"
	"fmt"

	"repo.nefrosovet.ru/maximus-platform/auth/storage"

	"github.com/sirupsen/logrus"
	ldaplib "gopkg.in/LDAP.v2"

	"repo.nefrosovet.ru/maximus-platform/auth/authentication/login/index"
	st "repo.nefrosovet.ru/maximus-platform/auth/storage/storage_accessor"
)

var (
	ErrWrongLogin    = errors.New("wrong login")
	ErrWrongPassword = errors.New("wrong password")
)

type Credentials struct {
	Login    string
	Password string

	Backend *storage.Backend
}

type Result struct {
	EntityID  string
	TempRoles []*storage.Role

	Error error
}

func Auth(credentials *Credentials) *Result {
	connect, err := GetConnect(credentials.Backend)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"context":   "CORE",
			"resource":  "ldap",
			"function":  "AuthByLDAP",
			"backendID": credentials.Backend.ID,
			"addr":      fmt.Sprintf("%s:%d", credentials.Backend.Host, credentials.Backend.Port),
			"status":    "FAILED",
		}).Error("LDAP connect failed")
		logrus.Debug(err)

		return &Result{
			Error: err,
		}
	}

	searchResult, err := connect.Search(credentials.Login)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"context":   "CORE",
			"resource":  "ldap",
			"function":  "AuthByLDAP",
			"backendID": credentials.Backend.ID,
			"addr":      fmt.Sprintf("%s:%d", credentials.Backend.Host, credentials.Backend.Port),
			"status":    "FAILED",
		}).Error("LDAP search request failed")
		logrus.Debug(err)

		return &Result{
			Error: err,
		}
	}

	connect.Close()

	if len(searchResult.Entries) == 0 {
		logrus.WithFields(logrus.Fields{
			"context":   "CORE",
			"resource":  "LDAP",
			"function":  "AuthByLDAP",
			"backendID": credentials.Backend.ID,
			"login":     credentials.Login,
			"status":    "FAILED",
		}).Debug("LDAP login not found")

		return &Result{
			Error: ErrWrongLogin,
		}
	}

	for _, entry := range searchResult.Entries {
		ldapConfig := &ConnectConfig{
			Host:         credentials.Backend.Host,
			Port:         int(credentials.Backend.Port),
			UserDN:       entry.DN,
			UserPassword: credentials.Password,
			Cipher:       credentials.Backend.Cipher,
		}

		loginConnect, err := New(ldapConfig)
		if err != nil {
			logrus.WithField("entry", entry).Debug("LDAP connect by login failed")

			continue
		}

		// Auth success

		loginConnect.Close()

		logrus.WithField("login", credentials.Login).Debug("LDAP connect by login success")

		user, err := getUserByLDAPEntry(credentials.Backend, entry)
		if err != nil {
			return &Result{
				Error: err,
			}
		}

		tempRoles, err := getRolesByLDAPGroups(credentials.Backend, entry)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"context":   "CORE",
				"resource":  "ldap",
				"function":  "AuthByLDAP",
				"backendID": credentials.Backend.ID,
				"login":     credentials.Login,
				"error":     err,
			}).Debug("getRolesByLDAPGroups failed")

			return &Result{
				EntityID: user.ID,
				Error:    err,
			}
		}

		return &Result{
			EntityID:  user.ID,
			TempRoles: tempRoles,
		}
	}

	// Auth failed

	return &Result{
		Error: ErrWrongPassword,
	}
}

func getUserByLDAPEntry(backend *storage.Backend, entry *ldaplib.Entry) (*storage.User, error) {
	if backend.Attributes.GUID.Map == nil {
		logrus.WithFields(logrus.Fields{
			"context":   "CORE",
			"resource":  "ldap",
			"function":  "getUserByLDAPEntry",
			"backendID": backend.ID,
		}).Error("GUID attribute information on backend not found! Check attributes.")

		return nil, errors.New(ErrorTypeInternal)
	}
	entryID := entry.GetAttributeValue(*backend.Attributes.GUID.Map)
	if entryID == "" {
		logrus.WithFields(logrus.Fields{
			"context":   "CORE",
			"resource":  "ldap",
			"function":  "getUserByLDAPEntry",
			"backendID": backend.ID,
		}).Error("can't get backend user GUID from entry. May be GUID attribute is wrong.")

		return nil, errors.New(ErrorTypeInternal)
	}

	us := st.GetStorage().UserStorage

	users, err := us.Get(storage.GetUser{
		BackendID:      &backend.ID,
		BackendEntryID: &entryID,
	})
	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		// Need to find or create user on index and local DB.

		var userID string
		userID, err = syncLDAPEntryWithIndex(backend, entry)
		if err != nil {
			return nil, err
		}
		if userID == "" {
			logrus.WithField("entryID", entryID).Error("no user GUID returned by Index")

			return nil, errors.New(ErrorTypeInternal)
		}

		users, err = us.Get(storage.GetUser{
			ID: &userID,
		})
		if err != nil {
			return nil, err
		}

		if len(users) == 0 {
			user, err := us.Store(storage.StoreUser{
				User: storage.User{
					ID: userID,
				},
			})
			if err != nil {
				return nil, err
			}

			users = append(users, user)
		}
		user := users[0]

		rs := st.GetStorage().RoleStorage
		role, err := rs.Get(backend.Sync)
		if err != nil {
			return nil, err
		}

		if user, err = us.Update(user.ID, storage.UpdateUser{
			Roles: map[string]bool{
				role.ID: true,
			},
			BackendEntryIDs: map[string]string{
				backend.ID: entryID,
			},
		}); err != nil {
			return nil, err
		}

		return user, nil
	}

	logrus.WithFields(logrus.Fields{
		"context":   "CORE",
		"resource":  "ldap",
		"function":  "syncLDAPEntryWithIndex",
		"backendID": backend.ID,
		"userGUID":  users[0].ID,
	}).Debug("Push ALWAYS sync attributes to index")

	go func() {
		params, err := fillIndexUserParamsFromLDAPEntry(backend, entry, false)
		if err != nil {
			return
		}

		switch backend.Sync {
		case storage.BackendSyncEmployee:
			index.PatchEmployee(users[0].ID, params)
		case storage.BackendSyncPatient:
			index.PatchPatient(users[0].ID, params)
		default:
			return
		}

		return
	}()

	return users[0], err
}

// syncLDAPEntryWithIndex
func syncLDAPEntryWithIndex(backend *storage.Backend, entry *ldaplib.Entry) (string, error) {
	var (
		userID string
		err    error
	)

	params, err := fillIndexUserParamsFromLDAPEntry(backend, entry, true)
	if err != nil {
		return userID, err
	}

	if params.Username == "" {
		logrus.WithFields(logrus.Fields{
			"context":   "CORE",
			"resource":  "ldap",
			"function":  "syncLDAPEntryWithIndex",
			"backendID": backend.ID,
		}).Error("username not found")
	}

	switch backend.Sync {
	case storage.BackendSyncEmployee:
		userID, err = index.CreateOrSearchEmployee(params)
	case storage.BackendSyncPatient:
		userID, err = index.CreateOrSearchPatient(params)
	default:
		return userID, errors.New("unknown backend sync type")
	}

	return userID, err
}

// fillIndexUserParamsFromLDAPEntry
func fillIndexUserParamsFromLDAPEntry(backend *storage.Backend, entry *ldaplib.Entry, firstLogin bool) (index.UserParams, error) {
	var params index.UserParams

	params.Username = entry.GetAttributeValue(*backend.Attributes.Username.Map)
	if backend.Attributes.Email.Map != nil && *backend.Attributes.Email.Map != "" && (firstLogin || *backend.Attributes.Email.Sync == storage.BackendAttrSyncAlways) {
		params.Email = entry.GetAttributeValue(*backend.Attributes.Email.Map)
	}
	if backend.Attributes.FirstName.Map != nil && *backend.Attributes.FirstName.Map != "" && (firstLogin || *backend.Attributes.FirstName.Sync == storage.BackendAttrSyncAlways) {
		params.FirstName = entry.GetAttributeValue(*backend.Attributes.FirstName.Map)
	}
	if backend.Attributes.LastName.Map != nil && *backend.Attributes.LastName.Map != "" && (firstLogin || *backend.Attributes.LastName.Sync == storage.BackendAttrSyncAlways) {
		params.LastName = entry.GetAttributeValue(*backend.Attributes.LastName.Map)
	}
	if backend.Attributes.Patronymic.Map != nil && *backend.Attributes.Patronymic.Map != "" && (firstLogin || *backend.Attributes.Patronymic.Sync == storage.BackendAttrSyncAlways) {
		params.Patronymic = entry.GetAttributeValue(*backend.Attributes.Patronymic.Map)
	}
	if backend.Attributes.Mobile.Map != nil && *backend.Attributes.Mobile.Map != "" && (firstLogin || *backend.Attributes.Mobile.Sync == storage.BackendAttrSyncAlways) {
		params.Mobile = entry.GetAttributeValue(*backend.Attributes.Mobile.Map)
	}

	return params, nil
}

func getRolesByLDAPGroups(backend *storage.Backend, entry *ldaplib.Entry) ([]*storage.Role, error) {
	groups := entry.GetAttributeValues("memberOf")
	roles := make([]*storage.Role, 0)

	for _, group := range groups {
		bs := st.GetStorage().BackendStorage

		backends, err := bs.Get(storage.GetBackend{
			ID: &backend.ID,
		})
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"context":  "CORE",
				"resource": "ldap",
				"function": "getRolesByLDAPGroups",
				"error":    err,
			}).Error("Can't get role by LDAP group")

			return nil, err
		} else if len(backends) == 0 {
			return nil, storage.ErrBackendNotFound
		}

		var role *storage.Role
		roleID := backends[0].Groups[group]
		if roleID != "" {
			role, err = st.GetStorage().RoleStorage.Get(roleID)
			if err != nil && !errors.Is(err, storage.ErrNotFound){
				return nil, err
			}
		}

		if role != nil {
			roles = append(roles, role)

			logrus.WithFields(logrus.Fields{
				"context":  "CORE",
				"resource": "ldap",
				"function": "getRolesByLDAPGroups",
				"group":    group,
				"role":     role.ID,
			}).Debug("appended role by LDAP group")
		}
	}

	return roles, nil
}

// TestLDAPBackend tests LDAP backend
func TestLDAPBackend(backend *storage.Backend) error {
	connect, err := GetConnect(backend)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"context":   "CORE",
			"resource":  "ldap",
			"function":  "TestLDAPBackend",
			"backendID": backend.ID,
			"addr":      fmt.Sprintf("%s:%d", backend.Host, backend.Port),
			"status":    "FAILED",
		}).Error("LDAP connect failed")
		logrus.Debug(err)

		return err
	}

	searchResult, err := connect.SearchWithPaging("*", 1000)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"context":   "CORE",
			"resource":  "ldap",
			"function":  "TestLDAPBackend",
			"backendID": backend.ID,
			"addr":      fmt.Sprintf("%s:%d", backend.Host, backend.Port),
			"status":    "FAILED",
		}).Error("LDAP search request failed")
		logrus.Debug(err)

		return err
	}

	connect.Close()

	if len(searchResult.Entries) == 0 {
		logrus.WithFields(logrus.Fields{
			"context":   "CORE",
			"resource":  "LDAP",
			"function":  "TestLDAPBackend",
			"backendID": backend.ID,
			"status":    "FAILED",
		}).Debug("no LDAP entries found")

		return errors.New("no LDAP entries found")
	}

	return nil
}
