package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/sharin-sushi/0022loginwithJWT/t0022Go/internal/utility"
)

func AuthMiddleware(c *gin.Context) {
	tokenString, err := c.Cookie("token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		c.Abort()
		return
	}

	token, err := utility.ParseToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid token",
		})
		c.Abort()
		return
	}

	c.Next()
}
