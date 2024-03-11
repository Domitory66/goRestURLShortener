package repository

import authorization "url-shortener"

type Authorization interface {
	CreateUser(authorization.User) (int, error)
	GetUser(username string) (authorization.User, error)
}
