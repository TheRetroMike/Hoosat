package handshake

import (
	"github.com/Hoosat-Oy/hoosatd/infrastructure/logger"
	"github.com/Hoosat-Oy/hoosatd/util/panics"
)

var log = logger.RegisterSubSystem("PROT")
var spawn = panics.GoroutineWrapperFunc(log)
