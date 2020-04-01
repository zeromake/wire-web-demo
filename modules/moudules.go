package modules

import (
	"github.com/google/wire"
	"github.com/zeromake/wire-web-demo/modules/minio"
)

var Set = wire.NewSet(
	minio.NewMinio,
)
