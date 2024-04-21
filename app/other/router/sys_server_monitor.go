package router

import (
	"goconf/app/other/apis"

	"github.com/gin-gonic/gin"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerSysServerMonitorRouter)
}

func registerSysServerMonitorRouter(v1 *gin.RouterGroup) {
	api := apis.ServerMonitor{}

	r := v1.Group("/server-monitor")
	{
		r.GET("", api.ServerInfo)
	}
}
