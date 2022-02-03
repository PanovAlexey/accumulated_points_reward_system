package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *httpHandler) login(c *gin.Context) {

	c.String(http.StatusOK, "OK")
}
