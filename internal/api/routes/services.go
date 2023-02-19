package routes

import (
	Database "orc-api/internal/database"
	Errors "orc-api/internal/errors"
	Shared "orc-api/internal/shared"

	"gorm.io/gorm/clause"
)

func listRoutes(params Shared.Params) ([]Route, error) {
	routes := []Route{}
	db := Database.GetDB()
	err := db.Limit(params.Limit).Offset(params.Offset).Select("id", "created_at", "name").Find(&routes).Error
	return routes, err
}

func createRoutePlan(schema RoutePlanPostSchema) (RoutePlanSchema, error) {
	plan := schema.parse()
	db := Database.GetDB()
	err := db.Clauses(clause.Returning{}).Create(&plan).Error
	return plan.Schema(), err
}

func updateRoute(route *Route) error {
	db := Database.GetDB()
	res := db.Clauses(clause.Returning{}).Where("id = ?", route.ID).Updates(route)
	if res.RowsAffected == 0 {
		return Errors.ResourceNotFoundErr
	}
	return res.Error
}
