package routes

import (
	"auth-system/controllers"
	"auth-system/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	// Public
	r.POST("/register", func(c *gin.Context) { controllers.Register(c, db) })
	r.POST("/login", func(c *gin.Context) { controllers.Login(c, db) })

	// Protected
	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())
	auth.GET("/profile", func(c *gin.Context) { controllers.Porfile(c, db) })
}
