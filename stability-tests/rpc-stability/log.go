package main

import (
	"github.com/nexepanet/nexepad/infrastructure/logger"
	"github.com/nexepanet/nexepad/util/panics"
)

var (
	backendLog = logger.NewBackend()
	log        = backendLog.Logger("JSTT")
	spawn      = panics.GoroutineWrapperFunc(log)
)
