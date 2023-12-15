package rpchandlers

import (
	"github.com/nexepanet/nexepad/infrastructure/logger"
	"github.com/nexepanet/nexepad/util/panics"
)

var log = logger.RegisterSubSystem("RPCS")
var spawn = panics.GoroutineWrapperFunc(log)
