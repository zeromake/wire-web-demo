package services

import (
	"github.com/google/wire"
	"github.com/zeromake/wire-web-demo/services/local"
)

var Set = wire.NewSet(
	local.Set,
)
