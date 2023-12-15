package grpcclient

import (
	"github.com/romxxxx/nexepad/infrastructure/logger"
	"github.com/romxxxx/nexepad/util/panics"
)

var log = logger.RegisterSubSystem("RPCC")
var spawn = panics.GoroutineWrapperFunc(log)
