package repository

import "github.com/jmoiron/sqlx"


type Posts interface{}

type Users interface{}

type Repository struct {
	Posts
	Users
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
