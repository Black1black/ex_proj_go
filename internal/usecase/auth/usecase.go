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

func (s *Usecase) GetIDByEmail(email string, hashedPassword string) (int64, error) {
	id, err := s.authRepo.GetIDByEmail(email, hashedPassword)
	if err == nil {
		return 0, err
	}
	return id, nil
}

func (s *Usecase) DeleteRefreshToken(id int64, token string) error {

	err := s.authRepo.DeleteToken(id, token)
	if err == nil {
		return err
	}
	return nil
}
