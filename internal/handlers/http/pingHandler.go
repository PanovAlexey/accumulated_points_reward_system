package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// HandlePing endpoint for checking the health of the service
func (h *httpHandler) HandlePing(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}
