package router

import (
	"goconf/app/admin/apis"
	jwt "goconf/core/sdk/pkg/jwtauth"

	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSysUserRouter)
}

func registerSysUserRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.SysUser{}

	v1auth := v1.Group("").Use(authMiddleware.MiddlewareFunc())
	{
		v1auth.GET("/getinfo", api.GetInfo)
	}
}
