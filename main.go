package main

import (
	"fmt"
	"os"

	Company "orc-api/internal/api/companies"
	Drivers "orc-api/internal/api/drivers"
	Routes "orc-api/internal/api/routes"
	Users "orc-api/internal/api/users"
	Vehicles "orc-api/internal/api/vehicles"

	Database "orc-api/internal/database"
	Shared "orc-api/internal/shared"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

func main() {
	args := os.Args
	fmt.Printf("Execution args: %v\n", args)

	if !Shared.IsProdEnv() {
		fmt.Println("Loading .env")
		if err := godotenv.Load(); err != nil {
			panic(err)
		}
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("required_company_type", Company.ValidateCompanyType)
		v.RegisterValidation("optional_company_type", Company.ValidateOptionalCompanyType)
	}

	Database.Start(Shared.GetEnvVars().DB_HOST, Shared.GetEnvVars().DB_USER, Shared.GetEnvVars().DB_PASSWORD, Shared.GetEnvVars().DB_NAME)
	if err := migrate(); err != nil {
		panic(err)
	}

	if len(args) > 1 {
		if args[1] == "seed" {
			seed()
		}
	} else {
		if err := startAPI(); err != nil {
			panic(err)
		}
	}
}

func startAPI() error {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://orc.lucasgouvea.com"},
		AllowMethods:     []string{"GET", "PATCH", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin", "content-type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	router.Use(Users.ValidateJWTHandler)

	v1Router := router.Group("/v1")

	Users.RegisterRoutes(v1Router)
	Drivers.RegisterRoutes(v1Router)
	Vehicles.RegisterRoutes(v1Router)
	Company.RegisterRoutes(v1Router)
	Routes.RegisterRoutes(v1Router)

	return router.Run("0.0.0.0:8081")
}

func migrate() error {
	fmt.Println(" *** Running migrations ***")
	models := []any{&Users.User{}, &Drivers.Driver{}, &Vehicles.Vehicle{}, &Company.Company{}, &Routes.RoutePlan{}}
	return Database.Migrate(models)
}

func seed() {
	Users.Seed()
}
