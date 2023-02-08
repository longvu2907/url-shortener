package routes

import (
	"url-shortener/controllers"
	"url-shortener/middlewares"

	"github.com/gin-gonic/gin"
)

func addShortUrlRoutes(r *gin.Engine) {
	r.GET("/:shortUrl", controllers.HandleShortUrlRedirect)

	group := r.Group("url-shortener")

	group.Use(middlewares.AuthMiddleware)
	group.POST("/create-short-url", controllers.CreateShortUrl)
	group.GET("/url-list", controllers.ListShortUrl)
}
