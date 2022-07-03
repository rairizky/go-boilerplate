package main

import (
	"boiler-go/config"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// load env
	godotenv.Load()

	// setup database
	db := config.SetupDB()
	fmt.Println("Connection to Database Successfull")

	// setup routes
	appPort := fmt.Sprintf(":%s", os.Getenv("APP_PORT"))
	router := config.SetupRoutes(db)
	router.Run(appPort)
}
