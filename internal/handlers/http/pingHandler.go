package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *httpHandler) HandlePing(c *gin.Context) {
	//@ToDo: add checking database availability
	c.String(http.StatusOK, "OK")
}
