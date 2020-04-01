package upload

import (
	"fmt"
	"net/http"
	"path"

	"github.com/gorilla/mux"
	"github.com/zeromake/wire-web-demo/types"
)

type Controller struct {
	fileProvider types.FileProvider
}

func NewController(fileProvider types.FileProvider) *Controller {
	return &Controller{
		fileProvider: fileProvider,
	}
}

func (c *Controller) InitRouter(router *mux.Router) {
	router.HandleFunc("/upload", c.Upload)
}

func (c *Controller) Upload(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(int64(100 * 1024 * 1024))
	if err != nil {
		_, _ = fmt.Fprintf(w, err.Error())
		return
	}
	f := r.MultipartForm.File["file"][0]
	in, err := f.Open()
	if err != nil {
		_, _ = fmt.Fprintf(w, err.Error())
		return
	}
	defer func() {
		_ = in.Close()
	}()
	out := path.Join("./temp", f.Filename)
	if !c.fileProvider.ExistsObject(out) {
		err = c.fileProvider.PutObject(out, in, f.Size)
		if err != nil {
			_, _ = fmt.Fprintf(w, err.Error())
			return
		}
	} else {
		_, _ = fmt.Fprintf(w, "该文件已存在")
		return
	}
	_, _ = fmt.Fprintf(w, "OK")
}
