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

type RoutePatchSchema struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (r RoutePlanPostSchema) parse() RoutePlan {
	plan := RoutePlan{}
	plan.EndDate = r.EndDate
	plan.StartDate = r.StartDate
	return plan
}

/* func (u RoutePatchSchema) parse(_id string) (*Route, error) {
	route := Route{}

	id, err := strconv.Atoi(_id)
	if err != nil {
		return nil, err
	}
	route.ID = uint(id)
	route.Name = u.Name
	route.Password = u.Password
	return &route, err
} */
