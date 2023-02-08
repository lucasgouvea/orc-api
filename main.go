package main

import (
	"fmt"
	"os"

	Users "orc-api/internal/api/users"
	Database "orc-api/internal/database"
	Shared "orc-api/internal/shared"

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

	Database.Start()

	if len(args) > 1 {
		if args[1] == "migrations:up" {
			migrationsUp()
		}
		if args[1] == "migrations:down" {
			migrationsDown()
		}
	} else {
		startAPI()
	}

}

func startAPI() {
	router := gin.Default()

	v1Router := router.Group("/v1")

	Users.RegisterRoutes(v1Router)

	router.Run("localhost:8080")
}

func migrationsUp() {
	fmt.Println(" *** Migrations Up ***")
	models := []any{&Users.User{}}
	Database.Migrate(models)
	Users.Seed()
}

func migrationsDown() {
	fmt.Println(" *** Migrations Down ***")
	tables := []string{"users"}
	Database.Drop(tables)
}
