// Package servers that run an HTTP server to provide endpoints for interacting with the service
package servers

import (
	"github.com/PanovAlexey/accumulated_points_reward_system/config"
	"github.com/PanovAlexey/accumulated_points_reward_system/internal/infrastructure/logging"
	"net/http"
	_ "net/http/pprof"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(config config.Config, handler http.Handler, logger logging.LoggerInterface) error {
	if config.Application.IsDebug {
		go func() {
			err := http.ListenAndServe(config.Server.DebugAddress, nil)
			if err != nil {
				logger.Error("error occurred while running http documentation server: %s", err.Error())
			}
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
