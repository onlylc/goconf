package router

import (
	"github.com/gin-gonic/gin"
	jwt "goconf/core/sdk/pkg/jwtauth"
	)
var (
	routerNoCheckRole = make([]func(*gin.RouterGroup),0)
	routerCheckRole = make([]func(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) ,0)
)
