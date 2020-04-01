package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type MinioConfig struct {
	Enable bool   // 是否启用 HTTP
	Host   string // HTTP host
	Port   int    // HTTP 端口
	Access string // Access
	Secret string // Secret
	Secure bool   // Secure
	Bucket string
}

type Application struct {
	Port int
}

type Config struct {
	Minio MinioConfig `toml:"minio"`
	Application  Application `toml:"application"`
}

func LoadConfig() (*Config, error) {
	var (
		cfg = &Config{}
	)
	viper.AddConfigPath("./configs/")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	err = viper.Unmarshal(cfg)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return cfg, nil
}
