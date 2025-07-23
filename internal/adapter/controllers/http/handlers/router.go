package handlers

import "github.com/gin-gonic/gin"

func InitRoutes(router *gin.Engine, userHandler *UserHandler) {

	auth := router.Group("/auth")
	{
		auth.POST("/register", userHandler.Register)
	}
}
