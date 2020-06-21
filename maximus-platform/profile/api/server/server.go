package server

import (
	"github.com/labstack/echo"
	"gopkg.in/hlandau/passlib.v1"

	"repo.nefrosovet.ru/maximus-platform/profile/db/mongo"
	"repo.nefrosovet.ru/maximus-platform/profile/logger"
	"repo.nefrosovet.ru/maximus-platform/profile/storage"
)

var (
	SuccessMessage = "SUCCESS"
)

type Server struct {
	Version string

	storage *storage.Storage

	authLog logger.AuthEntrier

	cli *mongo.Client

	e *echo.Echo
}

func New(version string, storage *storage.Storage, l logger.Logger, cli *mongo.Client) *Server {
	return &Server{
		Version: version,

		storage: storage,

		authLog: l.Auth(),

		cli: cli,
	}
}

// VerifyPassword validates password and regenerate hash
func VerifyPassword(password, hash string) (string, error) {
	passlib.UseDefaults(passlib.DefaultsLatest)

	return passlib.Verify(password, hash)
}

// HashPassword returns password hash
func HashPassword(password string) (string, error) {
	passlib.UseDefaults(passlib.DefaultsLatest)

	return passlib.Hash(password)
}
