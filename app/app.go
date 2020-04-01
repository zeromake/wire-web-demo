package app

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zeromake/wire-web-demo/config"
)

// NewServer new server
func NewServer(cfg *config.Config, router *mux.Router) (*http.Server, error) {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Application.Port),
		Handler: router,
	}
	return server, nil
}
