package oauth2

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strconv"

	"repo.nefrosovet.ru/maximus-platform/auth/storage"

	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	googleOAuth2 "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"

	"repo.nefrosovet.ru/maximus-platform/auth/authentication/login/index"
	"repo.nefrosovet.ru/maximus-platform/auth/authentication/oauth2/esia"
	"repo.nefrosovet.ru/maximus-platform/auth/jwt"
	st "repo.nefrosovet.ru/maximus-platform/auth/storage/storage_accessor"
)

var (
	GoogleConfig = &oauth2.Config{
		Scopes:   []string{"openid", "email", "profile"},
		Endpoint: google.Endpoint,
	}

	EsiaConfig = &oauth2.Config{
		Scopes:   []string{"openid", "email", "fullname", "contacts"},
		Endpoint: esia.Endpoint,
	}

	nonNumericRegex = regexp.MustCompile(`\D`)
)

type Credentials struct {
	Code        string
	RedirectURI string

	Backend *storage.Backend
}

type Entry struct {
	// ID - OAuth backend entry ID
	ID string

	Attributes map[string]string
}

type Result struct {
	EntityID string
	JWT      *jwt.JWT

	Error error
}

func Auth(credentials *Credentials) *Result {
	switch credentials.Backend.Provider {
	case storage.BackendOAuth2ProviderGoogle:
		ctx := context.Background()

		token, err := GoogleConfig.Exchange(ctx, credentials.Code)
		if err != nil {
			return &Result{
				Error: err,
			}
		}

		oauth2Service, err := googleOAuth2.NewService(ctx, option.WithTokenSource(GoogleConfig.TokenSource(ctx, token)), option.WithScopes(googleOAuth2.UserinfoEmailScope, googleOAuth2.UserinfoProfileScope))
		if err != nil {
			return &Result{
				Error: err,
			}
		}

		r, err := oauth2Service.Userinfo.Get().Do()
		if err != nil {
			return &Result{
				Error: err,
			}
		}

		entry := Entry{
			ID: r.Id,
			Attributes: map[string]string{
				"email": r.Email,

				"firstName": r.GivenName,
				"lastName":  r.FamilyName,
			},
		}

		user, err := getUserByEntry(credentials.Backend, entry)
		if err != nil {
			return &Result{
				Error: err,
			}
		}

		tokens, err := jwt.GenerateTokens(user.ID, "", []*storage.Role{})
		if err != nil {
			return &Result{
				EntityID: user.ID,
				Error:    err,
			}
		}

		return &Result{
			EntityID: user.ID,
			JWT:      tokens,
		}
	case storage.BackendOAuth2ProviderESIA:
		ctx := context.Background()

		result, err := esia.Exchange(
			ctx,
			credentials.Backend.ClientID,
			credentials.Backend.ClientSecret,
			EsiaConfig.Scopes,
			credentials.Code,
			credentials.RedirectURI,
		)
		if err != nil {
			return &Result{
				Error: err,
			}
		}

		personResponse := struct {
			FirstName  string `json:"firstName"`
			LastName   string `json:"lastName"`
			MiddleName string `json:"middleName"`
		}{}

		if err = esia.GetInfo(ctx, result.AuthToken.UrnEsiaSbjId, result.AccessToken, "/", &personResponse); err != nil {
			return &Result{
				Error: err,
			}
		}

		attributes := map[string]string{
			"firstName":  personResponse.FirstName,
			"lastName":   personResponse.LastName,
			"middleName": personResponse.MiddleName,
		}

		contactsResponse := struct {
			Elements []string `json:"elements"`
		}{}

		if err = esia.GetInfo(ctx, result.AuthToken.UrnEsiaSbjId, result.AccessToken, "/ctts", &contactsResponse); err != nil {
			return &Result{
				Error: err,
			}
		}

		for _, contactURL := range contactsResponse.Elements {
			contactResponse := struct {
				StateFacts []string `json:"stateFacts"`
				ID         int      `json:"id"`
				Type       string   `json:"type"`
				VrfStu     string   `json:"vrfStu"`
				Value      string   `json:"value"`
				ETag       string   `json:"eTag"`
			}{}

			if err = esia.GetInfo(ctx, -1, result.AccessToken, contactURL, &contactResponse); err != nil {
				return &Result{
					Error: err,
				}
			}

			if contactResponse.VrfStu != "VERIFIED" {
				continue
			}

			switch contactResponse.Type {
			case "EML":
				attributes["email"] = contactResponse.Value
			case "MBT":
				attributes["mobile"] = nonNumericRegex.ReplaceAllString(contactResponse.Value, "")
			}
		}

		entry := Entry{
			ID:         strconv.Itoa(int(result.AuthToken.UrnEsiaSbjId)),
			Attributes: attributes,
		}

		user, err := getUserByEntry(credentials.Backend, entry)
		if err != nil {
			return &Result{
				Error: err,
			}
		}

		tokens, err := jwt.GenerateTokens(user.ID, "", []*storage.Role{})

		return &Result{
			EntityID: user.ID,
			JWT:      tokens,
		}
	default:
		return &Result{
			Error: errors.New(fmt.Sprintf("No OAuth2 backend with ID '%v'", credentials.Backend.ID)),
		}
	}
}

func getUserByEntry(backend *storage.Backend, entry Entry) (*storage.User, error) {
	us := st.GetStorage().UserStorage
	users, err := us.Get(storage.GetUser{
		BackendID:      &backend.ID,
		BackendEntryID: &entry.ID,
	})
	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		// Not found in local DB. Try to find on Index
		userID, err := syncEntryWithIndex(backend, entry)
		if err != nil {
			return nil, err
		}

		if userID == "" {
			logrus.WithField("entryID", entry.ID).Error("no user GUID returned by Index")
			return nil, errors.New(ErrorTypeInternal)
		}

		// Found on Index. Create or update local DB entry
		users, err = us.Get(storage.GetUser{
			ID: &userID,
		})
		if err != nil {
			return nil, err
		}

		var user *storage.User
		if len(users) == 0 {
			user, err = us.Store(storage.StoreUser{
				User: storage.User{
					ID: userID,
				},
			})
			if err != nil {
				return nil, err
			}
		}

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
				backend.ID: entry.ID,
			},
		}); err != nil {
			return nil, err
		}

		return user, nil
	}
	user := users[0]

	go func() {
		params, err := fillIndexUserParamsFromEntry(backend, entry, false)
		if err != nil {
			return
		}
		switch backend.Sync {
		case storage.BackendSyncEmployee:
			_ = index.PatchEmployee(user.ID, params)
		case storage.BackendSyncPatient:
			_ = index.PatchPatient(user.ID, params)
		}
	}()
	return user, err
}

func syncEntryWithIndex(backend *storage.Backend, entry Entry) (userID string, err error) {
	params, err := fillIndexUserParamsFromEntry(backend, entry, true)
	if err != nil {
		return "", err
	}

	if params.Email == "" {
		logrus.WithFields(logrus.Fields{
			"context":   "CORE",
			"resource":  "OAuth2",
			"function":  "syncEntryWithIndex",
			"backendID": backend.ID,
		}).Error("email not found")
		return "", errors.New("empty email")
	}

	switch backend.Sync {
	case storage.BackendSyncEmployee:
		userID, err = index.CreateOrSearchEmployee(params)
	case storage.BackendSyncPatient:
		userID, err = index.CreateOrSearchPatient(params)
	default:
		return "", fmt.Errorf("unknown backend sync type: %s", backend.Sync)
	}

	return userID, err
}

func fillIndexUserParamsFromEntry(backend *storage.Backend, entry Entry, firstLogin bool) (index.UserParams, error) {
	var params index.UserParams

	if backend.Attributes.Email.Map != nil && *backend.Attributes.Email.Map != "" &&
		(firstLogin || *backend.Attributes.Email.Sync == storage.BackendAttrSyncAlways) {
		params.Email = entry.Attributes[*backend.Attributes.Email.Map]
	}

	if backend.Attributes.FirstName.Map != nil && *backend.Attributes.FirstName.Map != "" &&
		(firstLogin || *backend.Attributes.FirstName.Sync == storage.BackendAttrSyncAlways) {
		params.FirstName = entry.Attributes[*backend.Attributes.FirstName.Map]
	}

	if backend.Attributes.LastName.Map != nil && *backend.Attributes.LastName.Map != "" &&
		(firstLogin || *backend.Attributes.LastName.Sync == storage.BackendAttrSyncAlways) {
		params.LastName = entry.Attributes[*backend.Attributes.LastName.Map]
	}

	if backend.Attributes.Patronymic.Map != nil && *backend.Attributes.Patronymic.Map != "" &&
		(firstLogin || *backend.Attributes.Patronymic.Sync == storage.BackendAttrSyncAlways) {
		params.Patronymic = entry.Attributes[*backend.Attributes.Patronymic.Map]
	}

	if backend.Attributes.Mobile.Map != nil && *backend.Attributes.Mobile.Map != "" &&
		(firstLogin || *backend.Attributes.Mobile.Sync == storage.BackendAttrSyncAlways) {
		params.Mobile = entry.Attributes[*backend.Attributes.Mobile.Map]
	}

	return params, nil
}
