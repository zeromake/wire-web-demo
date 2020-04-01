package home

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Controller struct{}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) InitRouter(router *mux.Router) {
	router.HandleFunc("/", c.Home)
}

func (c *Controller) Home(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprint(w, "OK")
}
