package models

import (
	"errors"
	"time"

	"github.com/ybru-tech/georm"
	"gorm.io/gorm"
)

type Users struct {
	ID             int64     `gorm:"primaryKey;autoIncrement;type:bigint;not null"`
	Phone          *int      `gorm:"unique"`
	Email          *string   `gorm:"unique"`
	HashedPassword string    `gorm:"column:hashed_password;not null"`
	Status         string    `gorm:"not null;default:new"`
	DateReg        time.Time `gorm:"column:date_reg;not null;default:CURRENT_TIMESTAMP"`
	Name           string    `gorm:"not null"`
	Photo          *string
	Text           *string
}

func (Users) TableName() string {
	return "users"
}

func (u *Users) BeforeCreate(tx *gorm.DB) (err error) {
	return u.validate()
}

func (u *Users) BeforeUpdate(tx *gorm.DB) (err error) {
	return u.validate()
}

func (u *Users) validate() error {
	if u.Phone == nil && u.Email == nil {
		return errors.New("Either phone or email must be provided")
	}
	validStatuses := map[string]bool{
		"new":     true,
		"active":  true,
		"archive": true,
		"banned":  true,
	}
	if !validStatuses[u.Status] {
		return errors.New("Invalid status")
	}
	return nil
}

type UsersLocation struct {
	UserID   int64       `gorm:"column:user_id;primaryKey;type:bigint;not null"`
	Users    Users       `gorm:"foreignKey:UserID;references:ID"`
	Location georm.Point `gorm:"type:geometry(POINT,4326)"`
}

func (UsersLocation) TableName() string {
	return "users_location"
}
