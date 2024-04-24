package router

import (
	jwt "goconf/core/sdk/pkg/jwtauth"
	"goconf/app/admin/apis"
	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSysMenuRouter)
}

func registerSysMenuRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.SysMenu{}

	r1 := v1.Group("").Use(authMiddleware.MiddlewareFunc())
	{
		r1.GET("/menurole", api.GetMenuRole)
	}
}
