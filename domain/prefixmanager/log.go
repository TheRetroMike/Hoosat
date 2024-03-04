package prefixmanager

import (
	"github.com/Hoosat-Oy/htnd/infrastructure/logger"
	"github.com/Hoosat-Oy/htnd/util/panics"
)

var log = logger.RegisterSubSystem("PRFX")
var spawn = panics.GoroutineWrapperFunc(log)
