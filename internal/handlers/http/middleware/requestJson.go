package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RequestJSON() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Header.Get("Content-Type") != "application/json" {
			c.JSON(http.StatusBadRequest, "Invalid Content-Type")
			c.Abort()
			return
		}
	}
}
