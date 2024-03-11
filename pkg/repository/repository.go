package repository

type Repository struct {
	Storage
}

func NewRepository(s Storage) *Repository {
	return &Repository{Storage: s}
}
