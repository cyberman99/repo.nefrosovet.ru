package handlers

import "repo.nefrosovet.ru/maximus-platform/apigw/mongodb"

type Server struct {
	Version string
	repo    mongodb.PolicyRepository
}

func NewServer(version string, repo mongodb.PolicyRepository) *Server {
	return &Server{
		Version: version,
		repo:    repo,
	}
}
