package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *httpHandler) register(c *gin.Context) {

	c.String(http.StatusOK, "OK")
}
