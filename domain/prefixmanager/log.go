package prefixmanager

import (
	"github.com/nexepanet/nexepad/infrastructure/logger"
	"github.com/nexepanet/nexepad/util/panics"
)

var log = logger.RegisterSubSystem("PRFX")
var spawn = panics.GoroutineWrapperFunc(log)
