package authentication

import (
	"errors"

	"repo.nefrosovet.ru/maximus-platform/auth/storage"

	"repo.nefrosovet.ru/maximus-platform/auth/authentication/client"
	"repo.nefrosovet.ru/maximus-platform/auth/authentication/login"
	"repo.nefrosovet.ru/maximus-platform/auth/authentication/login/ldap"
	"repo.nefrosovet.ru/maximus-platform/auth/authentication/oauth2"
	"repo.nefrosovet.ru/maximus-platform/auth/jwt"
	st "repo.nefrosovet.ru/maximus-platform/auth/storage/storage_accessor"
)

var (
	ErrAuthMethodNotFound = errors.New("auth method not found")
	ErrNoRefreshToken     = errors.New("no RefreshToken")
	ErrInvalidToken       = errors.New("invalid token")
)

type Result struct {
	EntityID    string
	EntityLogin string
	JWT         *jwt.JWT

	error error
}

func (r *Result) Err() error {
	return r.error
}

func Auth(credentials interface{}) *Result {
	var res Result

	switch credentials := credentials.(type) {
	case *login.Credentials:
		r := login.Auth(credentials)

		res.EntityID = r.EntityID
		res.EntityLogin = r.EntityLogin
		res.JWT = r.JWT
		res.error = r.Error
	case *client.Credentials:
		r := client.Auth(credentials)

		res.EntityID = r.EntityID
		res.JWT = r.JWT
		res.error = r.Error
	case *oauth2.Credentials:
		r := oauth2.Auth(credentials)

		res.EntityID = r.EntityID
		res.JWT = r.JWT
		res.error = r.Error
	default:
		res.error = ErrAuthMethodNotFound
	}

	return &res
}

func Refresh(refreshToken string, tempRoles ...*storage.Role) *Result {
	if refreshToken == "" {
		return &Result{
			error: ErrNoRefreshToken,
		}
	}

	token, claims, err := jwt.ParseToken(refreshToken)
	if err != nil {
		return &Result{
			error: err,
		}
	}

	userID := claims["ID"].(string)

	if claims["type"] != "refresh" {
		return &Result{
			error: ErrInvalidToken,
		}
	}

	ts := st.GetStorage().TokenStorage
	oldToken, err := ts.GetByRefresh(token.Raw)
	if err != nil {
		// Kill used token for security reasons
		refreshedTokenRow, _ := ts.GetByParent(token.Raw)
		if refreshedTokenRow != nil {
			ts.Delete(refreshedTokenRow.Refresh)
		}

		return &Result{
			error: err,
		}
	}

	rs := st.GetStorage().RoleStorage
	// Inherit roles from refresh token (need for group roles)
	tokenRoles := claims["roles"].([]interface{})
	for _, roleIDInt := range tokenRoles {
		role, err := rs.Get(roleIDInt.(string))
		if err != nil {
			if errors.Is(err, storage.ErrNotFound) {
				continue
			}

			return &Result{
				error: err,
			}
		}

		tempRoles = append(tempRoles, role)
	}

	result, err := jwt.GenerateTokens(userID, token.Raw, tempRoles)
	if err != nil {
		return &Result{
			error: err,
		}
	}

	ts.Delete(oldToken.Refresh)

	return &Result{
		EntityID: result.UserID,
		JWT:      result,
	}
}

// TestBackend tests backend
func TestBackend(backend *storage.Backend) error {
	switch backend.Type {
	case storage.BackendTypeLDAP:
		return ldap.TestLDAPBackend(backend)
	}

	return errors.New("unknown backend type")
}
