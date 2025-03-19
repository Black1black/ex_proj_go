package users

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

func (r *Repository) GetByID(id int64) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil // Возвращает nil если запись не найдена
		}
		return nil, err // Возвращает ошибку в случае других проблем
	}
	return &user, nil
}

func (r *Repository) CreateUser() (string, error) {

	return "", nil
}
