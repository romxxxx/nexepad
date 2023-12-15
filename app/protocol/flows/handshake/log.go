package handshake

import (
	"github.com/romxxxx/nexepad/infrastructure/logger"
	"github.com/romxxxx/nexepad/util/panics"
)

var log = logger.RegisterSubSystem("PROT")
var spawn = panics.GoroutineWrapperFunc(log)
