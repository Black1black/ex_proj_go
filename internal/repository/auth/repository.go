package auth

import (
	"errors"
	"ex_proj_go/internal/models"

	"fmt"

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

func (r *Repository) GetIDByEmail(email string, hashedPassword string) (int64, error) {
	var user *models.User
	// Поиск пользователя по адресу электронной почты и хешированному паролю
	result := r.db.Where("email = ? AND hashed_password = ?", email, hashedPassword).First(&user)

	// Проверка на ошибки при выполнении запроса
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return 0, errors.New("пользователь не найден")
		}
		return 0, result.Error
	}

	return user.ID, nil
}

func (r *Repository) DeleteToken(userID int64, token string) error {
	result := r.db.Where("user_id = ? AND token = ?", userID, token).Delete(&models.AuthToken{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no records found to delete")
	}

	return nil
}
