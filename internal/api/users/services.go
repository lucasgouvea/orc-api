package users

import (
	Database "orc-api/internal/database"
)

func CreateUser(user User) error {
	user.hashPassword()
	db := Database.GetDB()
	err := db.Create(&user).Error
	return err
}
