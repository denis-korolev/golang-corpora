package route

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"parser/app/api/lemma"
	"parser/app/api/ping"
	_ "parser/docs"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()
	// Test
	router.GET("/ping", ping.HelloAction)

	//Lemma
	router.GET("/list", lemma.ListAction)

	// Swagger
	swaggerHandler := ginSwagger.WrapHandler(swaggerFiles.Handler)
	router.GET("/swagger/*any", swaggerHandler)

	return router
}
