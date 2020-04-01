// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/zeromake/wire-web-demo/app"
	"github.com/zeromake/wire-web-demo/config"
	"github.com/zeromake/wire-web-demo/controllers"
	"github.com/zeromake/wire-web-demo/controllers/home"
	"github.com/zeromake/wire-web-demo/controllers/upload"
	"github.com/zeromake/wire-web-demo/services/local"
	"github.com/zeromake/wire-web-demo/types"
	"net/http"
)

// Injectors from wire.go:

func NewServer() (*http.Server, error) {
	configConfig, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}
	controller := home.NewController()
	service := local.NewService()
	uploadController := upload.NewController(service)
	v := []types.Controller{
		controller,
		uploadController,
	}
	router := controllers.InitRouter(v)
	server, err := app.NewServer(configConfig, router)
	if err != nil {
		return nil, err
	}
	return server, nil
}
