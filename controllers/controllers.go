package controllers

import (
	"github.com/google/wire"
	"github.com/gorilla/mux"
	"github.com/zeromake/wire-web-demo/controllers/home"
	"github.com/zeromake/wire-web-demo/controllers/upload"
	"github.com/zeromake/wire-web-demo/types"
)

func InitRouter(handles []types.Controller) *mux.Router {
	router := mux.NewRouter()
	for _, h := range handles {
		h.InitRouter(router)
	}
	return router
}

var Set = wire.NewSet(
	InitRouter,
	wire.Slice(
		[]types.Controller(nil),
		home.NewController,
		upload.NewController,
	),
)
