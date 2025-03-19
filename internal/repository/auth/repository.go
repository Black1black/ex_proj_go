package auth

import (
	"ex_proj_go/internal/models"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) AddToken(userId int64, token string) error {
	authToken := models.AuthToken{
		UserID: int64(userId),
		Token:  token,
	}
	if err := r.db.Create(&authToken).Error; err != nil {
		return err
	}

	return nil
}
