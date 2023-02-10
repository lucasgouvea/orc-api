package drivers

import (
	"net/http"
	"strconv"

	Shared "orc-api/internal/shared"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/drivers", GetDrivers)
	router.POST("/drivers", PostDriver)
	router.PATCH("/drivers/:id", PatchDriver)
}

func GetDrivers(ctx *gin.Context) {
	var err error
	var params Shared.Params
	var drivers []Driver

	query := ctx.Request.URL.Query()

	if params, err = Shared.ParseQuery(query); err != nil {
		Shared.HandleErr(ctx, err)
		return
	}

	if drivers, err = listDrivers(params); err != nil {
		Shared.HandleErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, drivers)
}

func PostDriver(ctx *gin.Context) {
	schema := DriverPostSchema{}
	if err := ctx.ShouldBindWith(&schema, binding.JSON); err != nil {
		Shared.HandleErr(ctx, err)
		return
	}

	driver := schema.parse()

	if err := createDriver(driver); err != nil {
		Shared.HandleErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, driver)
}

func PatchDriver(ctx *gin.Context) {
	var id int
	var err error

	schema := DriverPatchSchema{}

	if id, err = strconv.Atoi(ctx.Param("id")); err != nil {
		Shared.HandleErr(ctx, err)
		return
	}

	if err := ctx.ShouldBindWith(&schema, binding.JSON); err != nil {
		Shared.HandleErr(ctx, err)
		return
	}

	if err := updateDriver(id, schema); err != nil {
		Shared.HandleErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, id)
}
