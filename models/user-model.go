package model

type User struct {
	BaseModel
	Email    string `gorm:"unique"`
	Password string
}