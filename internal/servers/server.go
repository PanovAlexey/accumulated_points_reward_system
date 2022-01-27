package servers

import (
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(config serviceConfigInterface, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           config.GetServerAddress() + ":" + config.GetServerPort(),
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}
