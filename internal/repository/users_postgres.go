package repository

import (
	"gorm.io/gorm"
)

type UsersPostgres struct {
	db *gorm.DB
}

func NewUsersPostgres(db *gorm.DB) *UsersPostgres {
	return &UsersPostgres{db: db}
}

func (r *UsersPostgres) CreateUser() (string, error) {

	return "", nil
}
