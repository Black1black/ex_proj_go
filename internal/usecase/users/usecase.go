package users

import (
	"ex_proj_go/internal/entity"
	"ex_proj_go/internal/models"
	"ex_proj_go/internal/usecase"
	"fmt"
)

type Usecase struct {
	usersRepo Users
	daoRepo   usecase.DAO
}

func NewUsecase(usersRepo Users, daoRepo usecase.DAO) *Usecase {
	return &Usecase{
		usersRepo: usersRepo,
		daoRepo:   daoRepo,
	}
}

func (u *Usecase) GetByID(id int64) (*entity.User, error) {
	user := models.User{}
	model, err := u.daoRepo.FindOneOrNone(user, map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}
	modelUser, ok := model.(models.User)
	if !ok {
		return nil, fmt.Errorf("Failed convert to model")
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
