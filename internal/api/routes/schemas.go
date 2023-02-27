package routes

import (
	Companies "orc-api/internal/api/companies"
	Drivers "orc-api/internal/api/drivers"
	"time"
)

/* Route Plan */

type RoutePlanSchema struct {
	Id        string    `json:"id"`
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

/* Route */

type RoutePostSchema struct {
	CompanyID   uint      `json:"company_id" binding:"required"`
	DriverID    uint      `json:"driver_id" binding:"required"`
	RoutePlanID uint      `json:"route_plan_id" binding:"required"`
	RouteType   string    `json:"route_type" binding:"required,route_type"`
	Date        time.Time `json:"date" binding:"required"`
}

type RoutePatchSchema struct {
	CompanyID   *uint      `json:"company_id"`
	DriverID    *uint      `json:"driver_id"`
	RoutePlanID *uint      `json:"route_plan_id"`
	RouteType   *string    `json:"route_type" binding:"route_type"`
	Date        *time.Time `json:"date"`
}

type RouteSchema struct {
	Id        string `json:"id"`
	Company   Companies.CompanySchema
	Driver    Drivers.DriverSchema
	RouteType routeType `json:"route_type"`
	Date      time.Time `json:"date"`
}

func (r RoutePostSchema) parse() Route {
	return Route{
		CompanyID:   r.CompanyID,
		DriverID:    r.DriverID,
		RoutePlanID: r.RoutePlanID,
		RouteType:   NewRouteType(r.RouteType),
		Date:        r.Date,
	}
}

func (r RoutePatchSchema) parse() map[string]any {
	var m map[string]any = make(map[string]any)

	if r.CompanyID != nil {
		m["company_id"] = *r.CompanyID
	}

	if r.Date != nil {
		m["date"] = *r.Date
	}

	if r.DriverID != nil {
		m["driver_id"] = *r.DriverID
	}

	if r.RoutePlanID != nil {
		m["route_plan_id"] = *r.RoutePlanID
	}

	if r.RouteType != nil {
		m["route_type"] = NewRouteType(*r.RouteType)
	}

	return m
}
