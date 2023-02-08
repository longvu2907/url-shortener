package routes

import (
	"url-shortener/controllers"

	"github.com/gin-gonic/gin"
)

func addAuthRoutes(r *gin.Engine) {
	group := r.Group("auth")

	group.POST("register", controllers.Register)
	group.POST("login", controllers.Login)
}
