package server

import (
	"github.com/labstack/echo"

	"repo.nefrosovet.ru/go-lms/api-video/ent"
	"repo.nefrosovet.ru/go-lms/api-video/logger"
)

type Server struct {
	Version string

	ent *ent.Client

	log logger.APIEntrier
	e   *echo.Echo
}

func New(version string, ent *ent.Client, log logger.Logger) *Server {
	return &Server{
		Version: version,

		ent: ent,
		log: log.Api(),
	}
}

func ptrS(s string) *string {
	return &s
}

func ptrI(i int) *int {
	return &i
}

func ptrB(b bool) *bool {
	return &b
}