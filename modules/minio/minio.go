package minio

import (
	"fmt"

	"github.com/minio/minio-go/v6"
	"github.com/zeromake/wire-web-demo/config"
)

func NewMinio(cfg *config.Config) (*minio.Client, error) {
	client, err := minio.New(
		fmt.Sprintf("%s:%d", cfg.Minio.Host, cfg.Minio.Port),
		cfg.Minio.Access,
		cfg.Minio.Secret,
		cfg.Minio.Secure,
	)
	if err != nil {
		return nil, err
	}
	ok, err := client.BucketExists(cfg.Minio.Bucket)
	if err != nil {
		return nil, err
	}
	if !ok {
		err = client.MakeBucket(cfg.Minio.Bucket, "")
		if err != nil {
			return nil, err
		}
	}
	return client, nil
}
