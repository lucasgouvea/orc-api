package routes

import (
	"net/http"
	"strconv"

	Shared "orc-api/internal/shared"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/route_plans", GetRoutePlans)
	router.POST("/route_plans", PostRoutePlan)
	router.PATCH("/route_plans/:id", PatchRoutePlan)
}

func GetRoutePlans(ctx *gin.Context) {
	var err error
	var params Shared.Params
	var plans []RoutePlan

	query := ctx.Request.URL.Query()

	if params, err = Shared.ParseQuery(query); err != nil {
		Shared.HandleErr(ctx, err)
		return
	}

	if plans, err = listRoutePlans(params); err != nil {
		Shared.HandleErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, plans)
}

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

func PatchRoutePlan(ctx *gin.Context) {
	var route *Route
	var err error
	var id int

	schema := RoutePlanPatchSchema{}
	if id, err = strconv.Atoi(ctx.Param("id")); err != nil {
		Shared.HandleErr(ctx, err)
		return
	}

	if err := ctx.ShouldBindWith(&schema, binding.JSON); err != nil {
		Shared.HandleErr(ctx, err)
		return
	}

	if err := updateRoutePlan(int(id), schema); err != nil {
		Shared.HandleErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, route)
}
