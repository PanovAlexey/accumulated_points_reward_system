// Package servers that run an HTTP server to provide endpoints for interacting with the service
package servers

import (
	"github.com/PanovAlexey/accumulated_points_reward_system/config"
	"net/http"
	_ "net/http/pprof"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(config config.Config, handler http.Handler) error {
	if config.Application.IsDebug == true {
		go func() {
			http.ListenAndServe(config.Server.DebugAddress, nil)
		}()
	}

	s.httpServer = &http.Server{
		Addr:           config.Server.Address,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}
