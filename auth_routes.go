package routes

import (
	"github.com/gin-gonic/gin"
	"bookapi/controllers"
)

func RegisterAuthRoutes(router *gin.Engine) {
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)
}
