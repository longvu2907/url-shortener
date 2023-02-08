package middlewares

import (
	"net/http"
	"url-shortener/utils/token"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	err := token.TokenValid(c)

	if err != nil {
		c.String(http.StatusUnauthorized, "Unauthorized")
		c.Abort()
		return
	}

	c.Next()
}
