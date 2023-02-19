package routes

import (
	"net/http"

	Shared "orc-api/internal/shared"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func RegisterRoutes(router *gin.RouterGroup) {
	//router.GET("/routes", GetRoutes)
	router.POST("/route_plans", PostRoutePlan)
	//router.PATCH("/routes/:id", PatchRoute)
}

/*
	 func GetRoutes(ctx *gin.Context) {
		var err error
		var params Shared.Params
		var routes []Route

		query := ctx.Request.URL.Query()

		if params, err = Shared.ParseQuery(query); err != nil {
			Shared.HandleErr(ctx, err)
			return
		}

		if routes, err = listRoutes(params); err != nil {
			Shared.HandleErr(ctx, err)
			return
		}

		ctx.JSON(http.StatusOK, routes)
	}
*/
func PostRoutePlan(ctx *gin.Context) {

	schema := RoutePlanPostSchema{}
	if err := ctx.ShouldBindWith(&schema, binding.JSON); err != nil {
		Shared.HandleErr(ctx, err)
		return
	}

	plan, err := createRoutePlan(schema)
	if err != nil {
		Shared.HandleErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, plan)
}

/* func PatchRoute(ctx *gin.Context) {
	var route *Route
	var err error

	schema := RoutePatchSchema{}
	id := ctx.Param("id")

	if err := ctx.ShouldBindWith(&schema, binding.JSON); err != nil {
		Shared.HandleErr(ctx, err)
		return
	}

	if route, err = schema.parse(id); err != nil {
		Shared.HandleErr(ctx, err)
		return
	}

	if err := updateRoute(route); err != nil {
		Shared.HandleErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, route)
}
*/
