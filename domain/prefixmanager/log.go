package prefixmanager

import (
	"github.com/romxxxx/nexepad/infrastructure/logger"
	"github.com/romxxxx/nexepad/util/panics"
)

var log = logger.RegisterSubSystem("PRFX")
var spawn = panics.GoroutineWrapperFunc(log)
