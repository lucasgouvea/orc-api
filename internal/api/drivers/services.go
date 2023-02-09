package drivers

import (
	Database "orc-api/internal/database"
	Errors "orc-api/internal/errors"
	Shared "orc-api/internal/shared"

	"gorm.io/gorm/clause"
)

func listDrivers(params Shared.Params) ([]Driver, error) {
	drivers := []Driver{}
	db := Database.GetDB()
	err := db.Limit(params.Limit).Offset(params.Offset).Select("id", "created_at", "name").Find(&drivers).Error
	return drivers, err
}

func createDriver(driver *Driver) error {
	db := Database.GetDB()
	err := db.Clauses(clause.Returning{}).Create(&driver).Error
	return err
}

func updateDriver(driver *Driver) error {
	db := Database.GetDB()
	res := db.Clauses(clause.Returning{}).Where("id = ?", driver.ID).Updates(driver)
	if res.RowsAffected == 0 {
		return Errors.ResourceNotFoundErr
	}
	return res.Error
}
