package service

type Service struct {
	usersRepo Users
	authRepo  Authorization
}

func NewService(usersRepo Users, authRepo Authorization) *Service {
	return &Service{
		usersRepo: usersRepo,
		authRepo:  authRepo,
	}
}
