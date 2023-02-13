package vehicles

import (
	Database "orc-api/internal/database"
	Errors "orc-api/internal/errors"
	Shared "orc-api/internal/shared"

	"gorm.io/gorm/clause"
)

func listVehicles(params Shared.Params) ([]Vehicle, error) {
	fields := []string{"id", "created_at", "model", "license_plate"}
	vehicles := []Vehicle{}
	db := Database.GetDB()
	err := db.Limit(params.Limit).Offset(params.Offset).Select(fields).Find(&vehicles).Error
	return vehicles, err
}

func createVehicle(schema VehiclePostSchema) (*Vehicle, error) {
	vehicle := schema.parse()

	db := Database.GetDB()
	err := db.Clauses(clause.Returning{}).Create(&vehicle).Error
	return vehicle, err
}

func updateVehicle(id int, schema VehiclePatchSchema) error {
	vehicleMap := schema.parse(id)

	db := Database.GetDB()
	res := db.Model(&Vehicle{}).Clauses(clause.Returning{}).Where("id = ?", id).Updates(vehicleMap)
	if res.Error == nil && res.RowsAffected == 0 {
		return Errors.ResourceNotFoundErr
	}
	return res.Error
}

func deleteVehicle(id int) error {
	db := Database.GetDB()
	res := db.Delete(&Vehicle{}, id)
	if res.Error == nil && res.RowsAffected == 0 {
		return Errors.ResourceNotFoundErr
	}
	return res.Error
}
