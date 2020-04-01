package local

import (
	"io"
	"os"
	"path"

	"github.com/google/wire"
	"github.com/zeromake/wire-web-demo/types"
)

const (
	DIR_MARK  os.FileMode = 0755
	FILE_MAEK os.FileMode = 0644
	FILE_FLAG             = os.O_CREATE | os.O_WRONLY | os.O_TRUNC
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) PutObject(name string, r io.Reader, _ int64) (err error) {
	dir := path.Dir(name)
	if !s.ExistsObject(dir) {
		err = os.MkdirAll(dir, DIR_MARK)
		if err != nil {
			return err
		}
	}
	dst, err := os.OpenFile(name, FILE_FLAG, FILE_MAEK)
	if err != nil {
		return err
	}
	defer func() {
		_ = dst.Close()
	}()
	_, err = io.Copy(dst, r)
	if err != nil {
		return err
	}
	return
}

func (s *Service) ExistsObject(name string) bool {
	return PathExists(name)
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

var Set = wire.NewSet(NewService, wire.Bind((*types.FileProvider)(nil), (*Service)(nil)))
