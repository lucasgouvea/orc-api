package users

import (
	Shared "orc-api/internal/shared"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type BaseSchema = any

type User struct {
	Shared.Model
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
	Password  string    `json:"-"`
	Token     string    `json:"-"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) hashPassword() {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	if err != nil {
		panic(err)
	}
	u.Password = string(bytes)
}

func (u *User) checkPasswordHash(hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(u.Password))
	return err == nil
}
