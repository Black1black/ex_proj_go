package users

import (
	"ex_proj_go/internal/entity"
)

type Usecase struct {
	usersRepo Users
}

func NewUsecase(usersRepo Users) *Usecase {
	return &Usecase{
		usersRepo: usersRepo,
	}
}

func (u *Usecase) GetByID(id int64) (*entity.User, error) {
	modelUser, err := u.usersRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return &entity.User{
		ID:      modelUser.ID,
		Phone:   modelUser.Phone,
		Email:   modelUser.Email,
		Status:  modelUser.Status,
		DateReg: modelUser.DateReg,
		Name:    modelUser.Name,
		Photo:   modelUser.Photo,
		Text:    modelUser.Text,
	}, nil
}
