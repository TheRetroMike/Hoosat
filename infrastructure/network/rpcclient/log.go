package rpcclient

import (
	"github.com/Hoosat-Oy/htnd/infrastructure/logger"
	"github.com/Hoosat-Oy/htnd/util/panics"
)

var log = logger.RegisterSubSystem("RPCC")
var spawn = panics.GoroutineWrapperFunc(log)
