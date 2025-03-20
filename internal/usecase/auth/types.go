package auth

type (
	Authorization interface {
		AddToken(userId int64, token string) error
		GetIDByEmail(email string, hashedPassword string) (int64, error)
		DeleteToken(id int64, token string) error
	}
)
