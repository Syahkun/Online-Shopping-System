package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/controllers/authcontroller"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := authcontroller.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}
