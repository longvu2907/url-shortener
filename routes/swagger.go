package routes

import (
	_ "url-shortener/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func AddSwaggerRoutes(r *gin.Engine) {
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
