package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader(authorizationHeader)
		if apiKey == "" {
			newErrorResponse(c, http.StatusUnauthorized, "empty")
			return
		}
		c.Next()
	}
}
