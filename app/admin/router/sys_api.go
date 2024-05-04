package router

import (
	"goconf/common/middleware"
	jwt "goconf/core/sdk/pkg/jwtauth"
	"goconf/app/admin/apis"
	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSysApiRouter)
}

// registerSysApiRouter
func registerSysApiRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.SysApi{}
	r := v1.Group("/sys-api").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetPage)
		r.GET("/:id", api.Get)
		r.PUT("/:id", api.Update)
	}
}
