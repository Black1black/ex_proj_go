package repository

import (
	"gorm.io/gorm"
)

type Authorization interface {
}

type Users interface {
}

type Repository struct {
	Authorization
	Users
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Users:         NewUsersPostgres(db),
	}
}
