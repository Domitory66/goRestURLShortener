package service

import (
	authorization "url-shortener"
	"url-shortener/pkg/repository"
)

type Storage interface {
	SaveURL(string, string) error
	GetURL(string) (string, error)
	DeleteURL(string) error
}

type Auth interface {
	CreateUser(user authorization.User) (int, error)
	GenerateToken(username, password string) string
}

type Service struct {
	Storage
	Auth
}

func NewService(repos *repository.Repository) *Service {
	return &Service{Storage: repos.Storage}
}
