package users

import (
	Shared "orc-api/internal/shared"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type BaseSchema = any

type User struct {
	Shared.Model
	Name         string
	Password     string
	Token        string
	TokenExpires time.Time
	Blocked      bool
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

func (u User) Schema() UserSchema {
	return UserSchema{
		Id:        strconv.FormatUint(uint64(u.ID), 10),
		CreatedAt: u.CreatedAt.String(),
		UpdatedAt: u.CreatedAt.String(),
		Name:      u.Name,
		Blocked:   u.Blocked,
	}
}
