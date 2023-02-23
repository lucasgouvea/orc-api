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
	router.DELETE("/drivers/:id", DeleteDriver)
}

func GetDrivers(ctx *gin.Context) {
	var err error
	var params Shared.Params
	var schemas []DriverSchema

	query := ctx.Request.URL.Query()

	if params, err = Shared.ParseQuery(query); err != nil {
		Shared.HandleErr(ctx, err)
		return
	}

	if schemas, err = listDrivers(params); err != nil {
		Shared.HandleErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, schemas)
}

func PostDriver(ctx *gin.Context) {
	var err error

	schema := DriverPostSchema{}
	if err = ctx.ShouldBindWith(&schema, binding.JSON); err != nil {
		Shared.HandleErr(ctx, err)
		return
	}

	if _, err = createDriver(schema); err != nil {
		Shared.HandleErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, schema)
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

func DeleteDriver(ctx *gin.Context) {
	var id int
	var err error

	if id, err = strconv.Atoi(ctx.Param("id")); err != nil {
		Shared.HandleErr(ctx, err)
		return
	}

	if err := deleteDriver(id); err != nil {
		Shared.HandleErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, id)
}
