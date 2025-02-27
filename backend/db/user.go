package db

import (
	"errors"
	"html"
	"log"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"size:255;not null; unique"`
	Username string `gorm:"size:255;not null; unique"`
	Password string `gorm:"size:255;not null"`
}

func GetUserById(uid uint) (User, error) {
	var user User

	db, err := ConnectDB()
	if err != nil {
		log.Println(err)
		return User{}, err
	}

	if err := db.Preload("Groceries").Where("id=?", uid).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return User{}, errors.New("song not found")
		}

	}

	return user, nil
}

func (u *User) HashedPassword() error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Failed to hash password: %v", err)
		return err
	}

	if err := VerifyPassword(u.Password, string(hashed)); err != nil {
		log.Printf("Failed to verify password: %v", err)
		return err
	}

	u.Password = string(hashed)
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	return nil
}

func VerifyPassword(password, hashPass string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(password))
}
