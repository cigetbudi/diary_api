package model

import (
	"diary_api/database"
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null;" json:"-"`
	Entries  []Entry
}

func (u *User) Save() (*User, error) {
	err := database.Database.Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) BeforeSave(*gorm.DB) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(passwordHash)
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	return nil
}

func (user *User) ValidatePassword(pass string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass))
}

func FindUserByUsername(un string) (User, error) {
	var user User
	err := database.Database.Where("username= ? ", un).Find(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func FindUserById(id uint) (User, error) {
	var user User
	err := database.Database.Preload("Entries").Where("ID = ?", id).Find(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}
