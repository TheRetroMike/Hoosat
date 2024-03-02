package main

import (
	"github.com/Hoosat-Oy/hoosatd/infrastructure/logger"
	"github.com/Hoosat-Oy/hoosatd/util/panics"
)

var (
	backendLog = logger.NewBackend()
	log        = backendLog.Logger("APLG")
	spawn      = panics.GoroutineWrapperFunc(log)
)
