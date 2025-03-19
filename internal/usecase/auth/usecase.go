package auth

type Usecase struct {
	authRepo Authorization
}

func NewUsecase(authRepo Authorization) *Usecase {
	return &Usecase{
		authRepo: authRepo,
	}
}

func (s *Usecase) Login(userId int64, token string) error {
	err := s.authRepo.AddToken(userId, token)
	if err == nil {
		return err
	}
	return nil
}

func (s *Usecase) GetIdByEmail(email string, hashedPassword string) (int64, error) {
	id, err := s.authRepo.GetIdByEmail(email, hashedPassword)
	if err == nil {
		return 0, err
	}
	return id, nil
}
