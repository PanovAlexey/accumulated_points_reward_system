package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func JSON() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")

		if c.Request.Header.Get("Content-Type") != "application/json" {
			c.JSON(http.StatusBadRequest, "Invalid Content-Type")
			c.Abort()
			return
		}
	}
}
