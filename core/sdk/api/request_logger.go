package api

import (
	"goconf/core/logger"
	"goconf/core/sdk"
	"strings"

	"goconf/core/sdk/pkg"

	"github.com/gin-gonic/gin"
)

type loggerKey struct{}

func GetRequestLogger(c *gin.Context) *logger.Helper {

	var log *logger.Helper
	l, ok := c.Get(pkg.LoggerKey)
	if ok {
		ok = false
		log, ok = l.(*logger.Helper)
		if ok {
			return log
		}
	}
	//如果没有在上下文中放入logger
	requestId := pkg.GenerateMsgIDFromContext(c)
	log = logger.NewHelper(sdk.Runtime.GetLogger()).WithFields(map[string]interface{}{
		strings.ToLower(pkg.TrafficKey): requestId,
	})
	return log
}

func SetRequestLogger(c *gin.Context) {
	requestId := pkg.GenerateMsgIDFromContext(c)
	log := logger.NewHelper(sdk.Runtime.GetLogger()).WithFields(map[string]interface{}{
		strings.ToLower(pkg.TrafficKey): requestId,
	})
	c.Set(pkg.LoggerKey, log)
}
