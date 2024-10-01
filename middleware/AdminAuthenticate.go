package middleware

import (
	"net/http"

	"github.com/ahmed-afzal1/restaurant/helpers"
	"github.com/gin-gonic/gin"
)

func AdminAuthenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := helpers.GetTokenFromHeader(c)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		claims, err := helpers.ValidateToken(tokenString)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Set("id", claims.ID)

		c.Next()
	}
}
