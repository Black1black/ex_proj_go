package users

import (
	"ex_proj_go/internal/models"
)

type (
	Users interface {
		GetByID(id int64) (*models.User, error)
	}
)
