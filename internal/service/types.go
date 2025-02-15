package service

type (
	Authorization interface {
		CreateToken() (string, error)
	}

	Users interface {
	}
)
