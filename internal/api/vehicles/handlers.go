package vehicles

import (
	"net/http"
	"strconv"

	Shared "orc-api/internal/shared"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/vehicles", GetVehicles)
	router.POST("/vehicles", PostVehicle)
	router.PATCH("/vehicles/:id", PatchVehicle)
	router.DELETE("/vehicles/:id", DeleteVehicle)
}

func GetVehicles(ctx *gin.Context) {
	var err error
	var params Shared.Params
	var schemas []VehicleSchema

	query := ctx.Request.URL.Query()

	if params, err = Shared.ParseQuery(query); err != nil {
		Shared.HandleErr(ctx, err)
		return
	}

	if schemas, err = listVehicles(params); err != nil {
		Shared.HandleErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, schemas)
}

func PostVehicle(ctx *gin.Context) {
	var vehicle *Vehicle
	var err error

	schema := VehiclePostSchema{}
	if err = ctx.ShouldBindWith(&schema, binding.JSON); err != nil {
		Shared.HandleErr(ctx, err)
		return
	}

	if vehicle, err = createVehicle(schema); err != nil {
		Shared.HandleErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, *vehicle)
}

func PatchVehicle(ctx *gin.Context) {
	var id int
	var err error

	schema := VehiclePatchSchema{}

	if id, err = strconv.Atoi(ctx.Param("id")); err != nil {
		Shared.HandleErr(ctx, err)
		return
	}

	if err := ctx.ShouldBindWith(&schema, binding.JSON); err != nil {
		Shared.HandleErr(ctx, err)
		return
	}

	if err := updateVehicle(id, schema); err != nil {
		Shared.HandleErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, id)
}

func DeleteVehicle(ctx *gin.Context) {
	var id int
	var err error

	if id, err = strconv.Atoi(ctx.Param("id")); err != nil {
		Shared.HandleErr(ctx, err)
		return
	}

	if err := deleteVehicle(id); err != nil {
		Shared.HandleErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, id)
}
