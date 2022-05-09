package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *httpHandler) getOrders(c *gin.Context) {

	c.JSON(http.StatusOK, map[string]interface{}{
		"temp info": "endpoint in progress",
	})
}
