package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func JSON() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Header.Get("Content-Type") != "application/json" {
			http.Error(c.Writer, "Invalid Content-Type", http.StatusBadRequest)
			return
		}

		c.Writer.Header().Set("Content-Type", "application/json")
	}
}
