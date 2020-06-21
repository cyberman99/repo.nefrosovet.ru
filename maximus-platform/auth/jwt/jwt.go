package jwt

import (
	"errors"
	"fmt"
	"time"

	"repo.nefrosovet.ru/maximus-platform/auth/storage"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"

	st "repo.nefrosovet.ru/maximus-platform/auth/storage/storage_accessor"
)

type JWT struct {
	UserID string

	AccessToken  string
	RefreshToken string
}

// GenerateTokens returns new access and refresh tokens
func GenerateTokens(userID string, parent string, tempRoles []*storage.Role) (*JWT, error) {
	secretKey := []byte(viper.GetString("tokenSecret"))
	creationTime := time.Now()

	if userID == "" {
		return nil, errors.New("no userID provided")
	}

	us := st.GetStorage().UserStorage
	users, err := us.Get(storage.GetUser{
		ID: &userID,
	})
	if err != nil {
		return nil, err
	}

	// FIXME:
	if len(users) == 0 {
		return nil, errors.New("user not found")
	}
	user := users[0]

	rolesMap := user.Roles
	for _, role := range tempRoles {
		rolesMap[role.ID] = true
	}

	rolesSlice := make([]string, 0)
	for key, value := range rolesMap {
		if value {
			rolesSlice = append(rolesSlice, key)
		}
	}

	refreshExpired := creationTime.Add(time.Second * viper.GetDuration("ttl.refreshToken"))
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"ID":    user.ID,
		"roles": rolesSlice,
		"type":  "access",
		"nbf":   creationTime.Unix(),
		"exp":   creationTime.Add(time.Second * viper.GetDuration("ttl.accessToken")).Unix(),
	})
	accessToken.Header["kid"] = "default"
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"ID":    user.ID,
		"roles": rolesSlice,
		"type":  "refresh",
		"nbf":   creationTime.Unix(),
		"exp":   refreshExpired.Unix(),
	})

	accessTokenString, err := accessToken.SignedString(secretKey)
	if err != nil {
		return nil, err
	}

	refreshTokenString, err := refreshToken.SignedString(secretKey)
	if err != nil {
		return nil, err
	}

	tokenDB := storage.Token{
		Access:   accessTokenString,
		Refresh:  refreshTokenString,
		Expired:  &refreshExpired,
		Parent:   parent,
		Username: user.ID,
	}

	ts := st.GetStorage().TokenStorage
	err = ts.Store(tokenDB)
	if err != nil {
		return nil, err
	}

	return &JWT{
		UserID:       userID,
		AccessToken:  tokenDB.Access,
		RefreshToken: tokenDB.Refresh,
	}, nil
}

// ParseToken parses access or refresh token and returns data
func ParseToken(tokenString string) (*jwt.Token, jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// HMACSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(viper.GetString("tokenSecret")), nil
	})

	if err != nil {
		return nil, nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, nil, errors.New("invalid token")
	}

	return token, claims, err
}
