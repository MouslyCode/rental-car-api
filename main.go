package main

import (
	"log"

	"github.com/MouslyCode/rental-car-api/database"
	"github.com/MouslyCode/rental-car-api/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.Connect()
	route := gin.Default()
	routes.User(route)

	route.Run(":8080")
}
