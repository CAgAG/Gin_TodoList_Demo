package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// 数据库模型

type User struct {
	gorm.Model
	UserName string `gorm:"unique"`
	Password string
}

const PasswordCost = 12

func (u *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PasswordCost)
	if err != nil {
		return err
	}
	u.Password = string(bytes)
	return nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return false
	}
	return true
}
