package types

import "io"

type FileProvider interface {
	PutObject(string, io.Reader, int64)  error
	ExistsObject(string) bool
}
