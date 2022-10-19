package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/j3yzz/encryptor/controllers"
)

func InitRoutes() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api/v1/")
	{
		auth := api.Group("/auth/")
		{
			auth.POST("/register", controllers.Register)
		}
	}

	return router
}
