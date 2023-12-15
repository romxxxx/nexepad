package consensus

import (
	"github.com/romxxxx/nexepad/infrastructure/logger"
	"github.com/romxxxx/nexepad/util/panics"
)

var log = logger.RegisterSubSystem("BDAG")
var spawn = panics.GoroutineWrapperFunc(log)
