package main

import (
	"fmt"
	"os"

	Drivers "orc-api/internal/api/drivers"
	Users "orc-api/internal/api/users"
	Database "orc-api/internal/database"
	Shared "orc-api/internal/shared"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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
		AllowOrigins:     []string{"http://orc-spa.s3-website-us-east-1.amazonaws.com"},
		AllowMethods:     []string{"GET", "PATCH", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin", "content-type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	v1Router := router.Group("/v1")

	Users.RegisterRoutes(v1Router)
	Drivers.RegisterRoutes(v1Router)

	return router.Run("0.0.0.0:8081")
}

func migrate() error {
	fmt.Println(" *** Running migrations ***")
	models := []any{&Users.User{}, &Drivers.Driver{}}
	return Database.Migrate(models)
}

func seed() {
	Users.Seed()
}
