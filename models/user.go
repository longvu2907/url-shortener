package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"size:255;not null;unique"`
	Password string `json:"password" gorm:"size:255;not null"`
}

// Create Hooks
func (u *User) BeforeCreate(tx *gorm.DB) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	return nil
}

func (u *User) CreateUser() (*User, error) {
	err := DB.Create(&u).Error

	if err != nil {
		return nil, err
	}

	return u, nil

}

func GetUserByUsername(username string) (*User, error) {
	user := &User{}

	err := DB.Model(User{}).Where("username = ?", username).Take(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetUserByID(id uint) (*User, error) {
	user := &User{}

	err := DB.Model(User{}).Where("id = ?", id).Take(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}
