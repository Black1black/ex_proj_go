package handler

import "ex_proj_go/internal/entity"

type (
	Users interface {
		GetByID(id int64) (*entity.User, error)
	}

	Authorization interface {
		Login(userId int64, token string) error
		GetIdByEmail(email string, hashedPassword string) (int64, error)
	}
)
