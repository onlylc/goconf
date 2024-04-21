package router

import (
	"goconf/core/tools/transfer"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerMonitorRouter)
}

func registerMonitorRouter(v1 *gin.RouterGroup) {
	v1.GET("/metrics", transfer.Handler(promhttp.Handler()))

	v1.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
}
