package auth

type (
	Authorization interface {
		AddToken(userId int64, token string) error
		GetIdByEmail(email string, hashedPassword string) (int64, error)
	}
)
