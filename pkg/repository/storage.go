package repository

const (
	ErrURLExists   = "Error URL exists"
	ErrURLNotFound = "Error URL not found"
)

type Storage interface {
	SaveURL(string, string) error
	GetURL(string) (string, error)
	DeleteURL(string) error
}
