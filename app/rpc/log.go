package rpc

import (
	"github.com/romxxxx/nexepad/infrastructure/logger"
	"github.com/romxxxx/nexepad/util/panics"
)

var log = logger.RegisterSubSystem("RPCS")
var spawn = panics.GoroutineWrapperFunc(log)
