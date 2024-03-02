package prefixmanager

import (
	"github.com/Hoosat-Oy/hoosatd/infrastructure/logger"
	"github.com/Hoosat-Oy/hoosatd/util/panics"
)

var log = logger.RegisterSubSystem("PRFX")
var spawn = panics.GoroutineWrapperFunc(log)
