package types

import "github.com/gorilla/mux"

type Controller interface {
	InitRouter(router *mux.Router)
}
