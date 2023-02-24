package users

import (
	"strconv"
	"time"
)

/* USERS */

type UserSchema struct {
	Id        string `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Name      string `json:"name"`
	Blocked   bool   `json:"blocked"`
}

type UserPostSchema struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserPatchSchema struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (u UserPostSchema) parse() *User {
	user := User{}
	user.Name = u.Name
	user.Password = u.Password
	return &user
}

func (u UserPatchSchema) parse(_id string) (*User, error) {
	user := User{}

	id, err := strconv.Atoi(_id)
	if err != nil {
		return nil, err
	}
	user.ID = uint64(id)
	user.Name = u.Name
	user.Password = u.Password
	return &user, err
}

/* LOGIN */

type PostLoginSchema struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type AuthSchema struct {
	Name        string    `json:"name"`
	Token       string    `json:"token"`
	Expires     string    `json:"expires"`
	ExpiresTime time.Time `json:"-"`
}

func (p PostLoginSchema) parse() *User {
	user := User{}
	user.Name = p.Name
	user.Password = p.Password
	return &user
}
