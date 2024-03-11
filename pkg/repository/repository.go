package repository

import (
	"url-shortener/pkg/storage"
)

type Repository struct {
	storage.Storage
}

func NewRepository(s storage.Storage) *Repository {
	return &Repository{Storage: s}
}
