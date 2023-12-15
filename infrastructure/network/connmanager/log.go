package connmanager

import (
	"github.com/nexepanet/nexepad/infrastructure/logger"
	"github.com/nexepanet/nexepad/util/panics"
)

var log = logger.RegisterSubSystem("CMGR")
var spawn = panics.GoroutineWrapperFunc(log)
