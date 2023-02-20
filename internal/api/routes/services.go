package routes

import (
	Database "orc-api/internal/database"
	Errors "orc-api/internal/errors"
	Shared "orc-api/internal/shared"

	"gorm.io/gorm/clause"
)

func listRoutePlans(params Shared.Params) ([]RoutePlanSchema, error) {
	plans := []RoutePlan{}
	schemas := []RoutePlanSchema{}
	db := Database.GetDB()
	err := db.Limit(params.Limit).Offset(params.Offset).Select("id", "created_at", "end_date", "start_date").Find(&plans).Error
	for _, p := range plans {
		schemas = append(schemas, p.Schema())
	}
	return schemas, err
}

func createRoutePlan(schema RoutePlanPostSchema) (RoutePlanSchema, error) {
	plan := schema.parse()
	db := Database.GetDB()
	err := db.Clauses(clause.Returning{}).Create(&plan).Error
	return plan.Schema(), err
}

func updateRoutePlan(id int, schema RoutePlanPatchSchema) error {
	plan := schema.parse()
	db := Database.GetDB()
	res := db.Model(&RoutePlan{}).Clauses(clause.Returning{}).Where("id = ?", id).Updates(plan)
	if res.Error == nil && res.RowsAffected == 0 {
		return Errors.ResourceNotFoundErr
	}
	return res.Error
}

func deleteRoutePlan(id int) error {
	db := Database.GetDB()
	res := db.Delete(&RoutePlan{}, id)
	if res.Error == nil && res.RowsAffected == 0 {
		return Errors.ResourceNotFoundErr
	}
	return res.Error
}
