package remote

import (
	"io"

	"github.com/google/wire"
	"github.com/minio/minio-go/v6"
	"github.com/zeromake/wire-web-demo/types"
)

type Service struct {
	Client *minio.Client
	Bucket string
}

func NewService(client *minio.Client) *Service {
	return &Service{
		Client: client,
		Bucket: "temp",
	}
}

func (s *Service) PutObject(name string, r io.Reader, size int64) error {
	_, err := s.Client.PutObject(
		s.Bucket,
		name,
		r,
		size,
		minio.PutObjectOptions{},
	)
	return err
}

func (s *Service) ExistsObject(name string) bool {
	_, err := s.Client.StatObject(s.Bucket, name, minio.StatObjectOptions{})
	return err == nil
}

var Set = wire.NewSet(NewService, wire.Bind((*types.FileProvider)(nil), (*Service)(nil)))
