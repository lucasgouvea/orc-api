package routes

import (
	Companies "orc-api/internal/api/companies"
	Drivers "orc-api/internal/api/drivers"

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
	Company   Companies.Company
	Driver    Drivers.Driver
	RouteType routeType
	Date      time.Time
}
