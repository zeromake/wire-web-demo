module github.com/zeromake/wire-web-demo

go 1.14

require (
	github.com/google/wire v0.4.0
	github.com/gorilla/mux v1.7.4
	github.com/minio/minio-go/v6 v6.0.50
	github.com/pkg/errors v0.9.1
	github.com/spf13/viper v1.6.2
)

replace github.com/google/wire v0.4.0 => github.com/zeromake/wire v0.0.0-20200331155702-af47ad758569
