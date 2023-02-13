package companies

import (
	"net/http"
	"strconv"

	Shared "orc-api/internal/shared"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/companies", GetCompanies)
	router.POST("/companies", PostCompany)
	router.PATCH("/companies/:id", PatchCompany)
	router.DELETE("/companies/:id", DeleteCompany)
}

func GetCompanies(ctx *gin.Context) {
	var err error
	var params Shared.Params
	var schemas []CompanySchema

	query := ctx.Request.URL.Query()

	if params, err = Shared.ParseQuery(query); err != nil {
		Shared.HandleErr(ctx, err)
		return
	}

	if schemas, err = listCompanies(params); err != nil {
		Shared.HandleErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, schemas)
}

func PostCompany(ctx *gin.Context) {
	var err error

	schema := CompanySchema{}
	postSchema := CompanyPostSchema{}

	if err = ctx.ShouldBindWith(&postSchema, binding.JSON); err != nil {
		Shared.HandleErr(ctx, err)
		return
	}

	if schema, err = createCompany(postSchema); err != nil {
		Shared.HandleErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, schema)
}

func PatchCompany(ctx *gin.Context) {
	var id int
	var err error

	schema := CompanyPatchSchema{}

	if id, err = strconv.Atoi(ctx.Param("id")); err != nil {
		Shared.HandleErr(ctx, err)
		return
	}

	if err := ctx.ShouldBindWith(&schema, binding.JSON); err != nil {
		Shared.HandleErr(ctx, err)
		return
	}

	if err := updateCompany(id, schema); err != nil {
		Shared.HandleErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, id)
}

func DeleteCompany(ctx *gin.Context) {
	var id int
	var err error

	if id, err = strconv.Atoi(ctx.Param("id")); err != nil {
		Shared.HandleErr(ctx, err)
		return
	}

	if err := deleteCompany(id); err != nil {
		Shared.HandleErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, id)
}
