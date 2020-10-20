package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/soganhei.com.br/atlas"
)

func authTokenJwt(jwt atlas.JwtServices) gin.HandlerFunc {
	return func(c *gin.Context) {

		token := strings.Split(c.Request.Header["Authorization"][0], " ")[1]

		tokenData, err := jwt.ParseAndVerifyToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		c.Set("token", tokenData)
		c.Next()
	}
}
