package service

import (
	"ex_proj_go/internal/repository"
)

type Authorization interface {
}

type Users interface {
}

type Service struct {
	Authorization
	Users
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Users:         NewUsersService(repos.Users),
	}
}
