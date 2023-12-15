package main

import (
	"github.com/romxxxx/nexepad/infrastructure/logger"
	"github.com/romxxxx/nexepad/util/panics"
)

var (
	backendLog = logger.NewBackend()
	log        = backendLog.Logger("APLG")
	spawn      = panics.GoroutineWrapperFunc(log)
)
