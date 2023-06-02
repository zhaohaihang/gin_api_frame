package model

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName       string `gorm:"unique"`
	PasswordDigest string
	Biography      string `gorm:"size:1000"`
	Address        string
	Email          string
	Phone          string
	Status         string
	Avatar         string `gorm:"size:1000"`
	LastLogin      int64
	Location       Point  `gorm:"type:point"`
	Extra          string `gorm:"size:1000"`
}

const (
	PassWordCost = 12       
	ACTIVE       = "active" 
)

func (user *User) BeforeSave(db *gorm.DB) error {
	user.LastLogin = time.Now().UnixMilli()
	return nil
}

func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}

func (user *User) AvatarURL() string {
	signedGetURL := user.Avatar
	return signedGetURL
}
