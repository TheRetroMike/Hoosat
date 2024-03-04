package main

import (
	"github.com/Hoosat-Oy/htnd/infrastructure/logger"
	"github.com/Hoosat-Oy/htnd/util/panics"
)

var (
	backendLog = logger.NewBackend()
	log        = backendLog.Logger("JSTT")
	spawn      = panics.GoroutineWrapperFunc(log)
)
