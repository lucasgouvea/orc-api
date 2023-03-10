package users

import (
	"net/http"

	Shared "orc-api/internal/shared"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/users", GetUsers)
	router.POST("/users", PostUser)
	router.PATCH("/users/:id", PatchUser)

	router.POST("/login", PostLogin)
}

func GetUsers(ctx *gin.Context) {
	var err error
	var params Shared.Params
	var schemas []UserSchema

	query := ctx.Request.URL.Query()

	if params, err = Shared.ParseQuery(query); err != nil {
		Shared.HandleErr(ctx, err)
		return
	}

	if schemas, err = listUsers(params); err != nil {
		Shared.HandleErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, schemas)
}

func PostUser(ctx *gin.Context) {
	schema := UserPostSchema{}
	if err := ctx.ShouldBindWith(&schema, binding.JSON); err != nil {
		Shared.HandleErr(ctx, err)
		return
	}

	user := schema.parse()

	if err := createUser(user); err != nil {
		Shared.HandleErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, schema)
}

func PatchUser(ctx *gin.Context) {
	var user *User
	var schema UserSchema
	var err error

	patchSchema := UserPatchSchema{}
	id := ctx.Param("id")

	if err := ctx.ShouldBindWith(&schema, binding.JSON); err != nil {
		Shared.HandleErr(ctx, err)
		return
	}

	if user, err = patchSchema.parse(id); err != nil {
		Shared.HandleErr(ctx, err)
		return
	}

	if err := updateUser(user); err != nil {
		Shared.HandleErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, user.Schema())
}

func PostLogin(ctx *gin.Context) {
	var auth *AuthSchema
	var err error

	schema := PostLoginSchema{}

	if err = ctx.ShouldBindWith(&schema, binding.JSON); err != nil {
		Shared.HandleErr(ctx, err)
		return
	}

	if auth, err = login(schema); err != nil {
		Shared.HandleErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, auth)
}
