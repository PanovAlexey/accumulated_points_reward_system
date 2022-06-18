package servers

import (
	"github.com/PanovAlexey/accumulated_points_reward_system/config"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(config config.Config, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           config.Server.Address,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}
