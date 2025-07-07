package main

import (
	"auth-system/config"
	"auth-system/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnectDatabase()
	r := gin.Default()

	routes.SetupRoutes(r, db)

	r.Run(":8080")
}
