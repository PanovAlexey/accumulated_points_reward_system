package http

import "github.com/gin-gonic/gin"

type httpHandler struct{}

func NewHandler() *httpHandler {
	return &httpHandler{}
}

func (h *httpHandler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/ping", h.HandlePing)

	return router
}
