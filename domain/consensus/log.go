package consensus

import (
	"github.com/Hoosat-Oy/hoosatd/infrastructure/logger"
	"github.com/Hoosat-Oy/hoosatd/util/panics"
)

var log = logger.RegisterSubSystem("BDAG")
var spawn = panics.GoroutineWrapperFunc(log)
