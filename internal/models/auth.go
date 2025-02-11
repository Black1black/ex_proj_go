package models

type AuthToken struct {
	UserID int64  `gorm:"column:user_id;primaryKey;type:bigint;not null"`
	User   User   `gorm:"foreignKey:UserID;references:ID"`
	Token  string `gorm:"column:token;primaryKey;not null"`
}

func (AuthToken) TableName() string {
	return "auth_token"
}
