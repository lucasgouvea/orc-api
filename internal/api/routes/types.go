package routes

import "github.com/go-playground/validator/v10"

type routeType int

const (
	ENTRANCE routeType = iota + 1
	EXIT
	BOTH
)

func getRouteTypes() []string {
	return []string{"ENTRANCE", "EXIT", "BOTH"}
}

func isRouteType(rt string) bool {
	for _, _rt := range getRouteTypes() {
		if _rt == rt {
			return true
		}
	}
	return false
}

func NewRouteType(r string) routeType {
	switch r {
	case "ENTRANCE":
		return ENTRANCE
	case "EXIT":
		return EXIT
	case "BOTH":
		return BOTH
	}
	panic(InvalidRouteTypeErr)
}

func ValidateRouteType(fl validator.FieldLevel) bool {
	value := fl.Field().Interface()

	if value == (*string)(nil) {
		return true // in case its nil, skips validation
	}
	str := value.(string)
	if isRouteType(str) {
		return true
	}

	return false
}
