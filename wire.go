//+build wireinject
//go:generate wire

package main

import (
	"net/http"

	"github.com/google/wire"
	"github.com/zeromake/wire-web-demo/app"
	"github.com/zeromake/wire-web-demo/config"
	"github.com/zeromake/wire-web-demo/controllers"
	"github.com/zeromake/wire-web-demo/services"
)

func NewServer() (*http.Server, error) {
	panic(
		wire.Build(
			controllers.Set,
			services.Set,
			//modules.Set,
			config.LoadConfig,
			app.NewServer,
		),
	)
}
