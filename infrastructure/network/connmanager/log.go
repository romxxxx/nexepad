package connmanager

import (
	"github.com/romxxxx/nexepad/infrastructure/logger"
	"github.com/romxxxx/nexepad/util/panics"
)

var log = logger.RegisterSubSystem("CMGR")
var spawn = panics.GoroutineWrapperFunc(log)
