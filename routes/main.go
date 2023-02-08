package routes

import "github.com/gin-gonic/gin"

func Config(r *gin.Engine) {
	AddSwaggerRoutes(r)
	addAuthRoutes(r)
	addShortUrlRoutes(r)
}
