package drivers

import "strconv"

type DriverPostSchema struct {
	Id       string `json:"id"`
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type DriverPatchSchema struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (u DriverPostSchema) parse() *Driver {
	user := Driver{}
	user.Name = u.Name
	return &user
}

func (u DriverPatchSchema) parse(_id string) (*Driver, error) {
	user := Driver{}

	id, err := strconv.Atoi(_id)
	if err != nil {
		return nil, err
	}
	user.ID = id
	user.Name = u.Name
	return &user, err
}
