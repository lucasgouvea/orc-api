package routes

import (
	Database "orc-api/internal/database"
	Errors "orc-api/internal/errors"
	Shared "orc-api/internal/shared"

	"gorm.io/gorm/clause"
)

func listRoutePlans(params Shared.Params) ([]RoutePlan, error) {
	plans := []RoutePlan{}
	db := Database.GetDB()
	err := db.Limit(params.Limit).Offset(params.Offset).Select("id", "created_at", "end_date", "start_date").Find(&plans).Error
	return plans, err
}

func createRoutePlan(schema RoutePlanPostSchema) (RoutePlanSchema, error) {
	plan := schema.parse()
	db := Database.GetDB()
	err := db.Clauses(clause.Returning{}).Create(&plan).Error
	return plan.Schema(), err
}

func updateRoutePlan(id int, schema RoutePlanPatchSchema) error {
	m := schema.parse()
	db := Database.GetDB()
	res := db.Clauses(clause.Returning{}).Where("id = ?", id).Updates(m)
	if res.RowsAffected == 0 {
		return Errors.ResourceNotFoundErr
	}
	return res.Error
}
