package main

import (
	// "login-signup-api/controllers"
	"login-signup-api/config"
	"login-signup-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.New()
	config.Connect()
	// public := router.Group("/api")

	// public.POST("/register", controllers.Register)
	routes.UserRoute(router)
	router.Run(":8080")

}