package main

import (
	"github.com/romxxxx/nexepad/infrastructure/logger"
	"github.com/romxxxx/nexepad/util/panics"
)

var (
	backendLog = logger.NewBackend()
	log        = backendLog.Logger("RPIC")
	spawn      = panics.GoroutineWrapperFunc(log)
)
