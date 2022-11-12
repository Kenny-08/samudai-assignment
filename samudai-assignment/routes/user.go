package routes

import (
	"login-signup-api/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine){
	router.GET("/", controllers.GetUsers)
	router.POST("/", controllers.Register)
	router.DELETE("/:id", controllers.DeleteUser)
	router.PUT("/:id", controllers.UpdateUser)
}