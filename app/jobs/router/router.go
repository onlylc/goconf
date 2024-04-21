package router

import (
	jwt "goconf/core/sdk/pkg/jwtauth"

	"github.com/gin-gonic/gin"
)

var (
	routerNoCheckRole = make([]func(*gin.RouterGroup), 0)
	routerCheckRole   = make([]func(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware), 0)
)

func initRouter(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) *gin.Engine {
	noCheckRoleRouter(r)

	checkRoleRouter(r, authMiddleware)

	return r
}

func noCheckRoleRouter(r *gin.Engine) {
	v1 := r.Group("/api/v1")

	for _, f := range routerNoCheckRole {
		f(v1)
	}
}

func checkRoleRouter(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) {
	v1 := r.Group("/api/v1")

	for _, f := range routerCheckRole {
		f(v1, authMiddleware)
	}
}
