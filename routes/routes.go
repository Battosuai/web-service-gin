package routes

import (
	"example/web-service-gin/services"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
    router.POST("/users", services.CreateUser)
    // Add more routes for Read, Update, Delete
}