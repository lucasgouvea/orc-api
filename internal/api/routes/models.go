package routes

import (
	Companies "orc-api/internal/api/companies"
	Drivers "orc-api/internal/api/drivers"
	"strconv"

	Shared "orc-api/internal/shared"
	"time"
)

type RoutePlan struct {
	Shared.Model
	StartDate   time.Time
	EndDate     time.Time
	Routes      []Route
	TotalRoutes int
}

type Route struct {
	Shared.Model
	CompanyID   uint
	DriverID    uint
	RoutePlanID uint
	Company     Companies.Company
	Driver      Drivers.Driver
	RouteType   routeType
	Date        time.Time
}

func (RoutePlan) TableName() string {
	return "route_plans"
}

func (Route) TableName() string {
	return "routes"
}

func (rp RoutePlan) Schema() RoutePlanSchema {
	return RoutePlanSchema{
		Id:        strconv.FormatUint(uint64(rp.ID), 10),
		EndDate:   rp.EndDate,
		StartDate: rp.StartDate,
	}
}
