package route

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"parser/app/api/ping"
	_ "parser/docs"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()
	router.GET("/ping", ping.PingAction)

	// Swagger
	swaggerHandler := ginSwagger.WrapHandler(swaggerFiles.Handler)
	router.GET("/swagger/*any", swaggerHandler)

	return router
}
