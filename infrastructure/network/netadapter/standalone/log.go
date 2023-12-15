package standalone

import (
	"github.com/romxxxx/nexepad/infrastructure/logger"
	"github.com/romxxxx/nexepad/util/panics"
)

var log = logger.RegisterSubSystem("NTAR")
var spawn = panics.GoroutineWrapperFunc(log)
