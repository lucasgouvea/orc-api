package companies

import (
	Database "orc-api/internal/database"
	Errors "orc-api/internal/errors"
	Shared "orc-api/internal/shared"

	"gorm.io/gorm/clause"
)

func listCompanies(params Shared.Params) ([]Company, error) {
	fields := []string{"id", "created_at", "name", "type"}
	companies := []Company{}
	db := Database.GetDB()
	err := db.Limit(params.Limit).Offset(params.Offset).Select(fields).Find(&companies).Error
	return companies, err
}

func createCompany(schema CompanyPostSchema) (*Company, error) {
	company := schema.parse()

	db := Database.GetDB()
	err := db.Clauses(clause.Returning{}).Create(&company).Error
	return company, err
}

func updateCompany(id int, schema CompanyPatchSchema) error {
	vehicleMap := schema.parse(id)

	db := Database.GetDB()
	res := db.Model(&Company{}).Clauses(clause.Returning{}).Where("id = ?", id).Updates(vehicleMap)
	if res.Error == nil && res.RowsAffected == 0 {
		return Errors.ResourceNotFoundErr
	}
	return res.Error
}

func deleteCompany(id int) error {
	db := Database.GetDB()
	res := db.Delete(&Company{}, id)
	if res.Error == nil && res.RowsAffected == 0 {
		return Errors.ResourceNotFoundErr
	}
	return res.Error
}
