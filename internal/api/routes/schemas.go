package routes

import (
	"time"
)

type RoutePlanSchema struct {
	EndDate   time.Time `json:"end_date"`
	StartDate time.Time `json:"start_date"`
}

type RoutePlanPostSchema struct {
	EndDate   time.Time `json:"end_date" binding:"required"`
	StartDate time.Time `json:"start_date" binding:"required"`
}

type RoutePlanPatchSchema struct {
	EndDate   *string `json:"end_date"`
	StartDate *string `json:"start_date"`
}

func (r RoutePlanPostSchema) parse() RoutePlan {
	plan := RoutePlan{}
	plan.EndDate = r.EndDate
	plan.StartDate = r.StartDate
	return plan
}

func (r RoutePlanPatchSchema) parse() map[string]any {
	var m map[string]any = make(map[string]any)

	if r.EndDate != nil {
		m["end_date"] = *r.EndDate
	}
	if r.StartDate != nil {
		m["start_date"] = *r.StartDate
	}

	return m
}
