package users

import "strconv"

/* USERS */

type UserPostSchema struct {
	Id       string `json:"id"`
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
	user.ID = uint(id)
	user.Name = u.Name
	user.Password = u.Password
	return &user, err
}

/* LOGIN */

type PostLoginSchema struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (p PostLoginSchema) parse() *User {
	user := User{}
	user.Name = p.Name
	user.Password = p.Password
	return &user
}
