package model

import (
	"example/helpers"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Password string
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Password = helpers.HashPassword(u.Password)
	return
}
