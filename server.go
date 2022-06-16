package initServerApp

import (
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (server Server) Run(port string, handler http.Handler) error {
	server.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return server.httpServer.ListenAndServe()
}
