package drivers

import (
	Database "orc-api/internal/database"
	Errors "orc-api/internal/errors"
	Shared "orc-api/internal/shared"

	"gorm.io/gorm/clause"
)

func listDrivers(params Shared.Params) ([]Driver, error) {
	fields := []string{"id", "created_at", "name", "age", "license_a", "license_b", "license_c", "license_d", "license_e"}
	drivers := []Driver{}
	db := Database.GetDB()
	err := db.Limit(params.Limit).Offset(params.Offset).Select(fields).Find(&drivers).Error
	return drivers, err
}

func createDriver(schema DriverPostSchema) (*Driver, error) {
	driver := schema.parse()

	db := Database.GetDB()
	err := db.Clauses(clause.Returning{}).Create(&driver).Error
	return driver, err
}

func updateDriver(id int, schema DriverPatchSchema) error {
	driverMap := schema.parse()

	db := Database.GetDB()
	res := db.Model(&Driver{}).Clauses(clause.Returning{}).Where("id = ?", id).Updates(driverMap)
	if res.Error == nil && res.RowsAffected == 0 {
		return Errors.ResourceNotFoundErr
	}
	return res.Error
}

func deleteDriver(id int) error {
	db := Database.GetDB()
	res := db.Delete(&Driver{}, id)
	if res.Error == nil && res.RowsAffected == 0 {
		return Errors.ResourceNotFoundErr
	}
	return res.Error
}
