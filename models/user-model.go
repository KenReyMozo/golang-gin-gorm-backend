package model

type User struct {
	BaseModel
	Email    string `gorm:"unique"`
	Password string
}

type RawUser struct {
	Email    string
	Password string
}
