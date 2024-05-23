package main

import (
	"jwt-gin/controllers"
	"jwt-gin/middlewares"
	"jwt-gin/models"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDataBase()

	if models.DB == nil {
		log.Fatal("Database connection is not established")
	}

	router := gin.Default()

	public := router.Group("/api")

	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	protected := router.Group(("/api/admin"))
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user", controllers.CurrentUser)

	router.Run(":8080")
}
