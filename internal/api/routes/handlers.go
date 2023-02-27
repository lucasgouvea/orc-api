package routes

import (
	"net/http"
	"strconv"

	Shared "orc-api/internal/shared"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func RegisterRoutes(router *gin.RouterGroup) {

	/* Route Plan */
	router.GET("/route_plans", GetRoutePlans)
	router.POST("/route_plans", PostRoutePlan)
	router.PATCH("/route_plans/:id", PatchRoutePlan)
	router.DELETE("/route_plans/:id", DeleteRoutePlan)

	/* Route */
	router.POST("/routes", PostRoute)
	router.PATCH("/routes/:id", PatchRoute)
	router.DELETE("/routes/:id", DeleteRoute)
}

/* Route Plans */

func GetRoutePlans(ctx *gin.Context) {
	var err error
	var params Shared.Params
	var schemas []RoutePlanSchema

	query := ctx.Request.URL.Query()

	if params, err = Shared.ParseQuery(query); err != nil {
		Shared.HandleErr(ctx, err)
		return
	}

	if schemas, err = listRoutePlans(params); err != nil {
		Shared.HandleErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, schemas)
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

	ctx.JSON(http.StatusOK, id)
}

func DeleteRoutePlan(ctx *gin.Context) {
	var id int
	var err error

	if id, err = strconv.Atoi(ctx.Param("id")); err != nil {
		Shared.HandleErr(ctx, err)
		return
	}

	if err := deleteRoutePlan(id); err != nil {
		Shared.HandleErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, id)
}

/* Route */

func PostRoute(ctx *gin.Context) {
	postSchema := RoutePostSchema{}
	if err := ctx.ShouldBindWith(&postSchema, binding.JSON); err != nil {
		Shared.HandleErr(ctx, err)
		return
	}

	r := postSchema.parse()
	if err := createRoute(r); err != nil {
		Shared.HandleErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, postSchema)
}

func PatchRoute(ctx *gin.Context) {
	var err error
	var id int

	schema := RoutePatchSchema{}
	if id, err = strconv.Atoi(ctx.Param("id")); err != nil {
		Shared.HandleErr(ctx, err)
		return
	}

	if err := ctx.ShouldBindWith(&schema, binding.JSON); err != nil {
		Shared.HandleErr(ctx, err)
		return
	}

	if err := updateRoute(int(id), schema.parse()); err != nil {
		Shared.HandleErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, id)
}

func DeleteRoute(ctx *gin.Context) {
	var id int
	var err error

	if id, err = strconv.Atoi(ctx.Param("id")); err != nil {
		Shared.HandleErr(ctx, err)
		return
	}

	if err := deleteRoute(id); err != nil {
		Shared.HandleErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, id)
}
