package router

import (
	"goconf/core/sdk/config"
	jwt "goconf/core/sdk/pkg/jwtauth"

	"goconf/app/admin/apis"
	"goconf/common/middleware"
	"goconf/common/middleware/handler"
	"goconf/core/sdk/pkg/ws"
	"mime"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"

	swaggerfiles "github.com/swaggo/files"
)

func InitSysRouter(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) *gin.RouterGroup {
	g := r.Group("")
	sysBaseRouter(g)
	sysStaticFileRouter(g)
	if config.ApplicationConfig.Mode != "prod" {
		sysSwaggerRouter(g)
	}
	sysCheckRoleRouterInit(g, authMiddleware)
	return g
}

func sysBaseRouter(r *gin.RouterGroup) {
	go ws.WebsocketManager.Start()
	go ws.WebsocketManager.SendService()
	go ws.WebsocketManager.SendAllService()

	if config.ApplicationConfig.Mode != "prod" {
		r.GET("/", apis.GoAdmin)
	}
	r.GET("/info", handler.Ping)
}

func sysStaticFileRouter(r *gin.RouterGroup) {
	err := mime.AddExtensionType(".js", "application/javascript")
	if err != nil {
		return
	}
	r.Static("/static", "./static")
	if config.ApplicationConfig.Mode != "prod" {
		r.Static("/form-generator", "./static/form-generator")
	}
}

func sysSwaggerRouter(r *gin.RouterGroup) {
	r.GET("/swagger/admin/*any", ginSwagger.WrapHandler(swaggerfiles.NewHandler(), ginSwagger.InstanceName("admin")))
}

func sysCheckRoleRouterInit(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {

	wss := r.Group("").Use(authMiddleware.MiddlewareFunc())
	{
		wss.GET("/ws/:id/:channel", ws.WebsocketManager.WsClient)
		wss.GET("/wslogout/:id/:channel", ws.WebsocketManager.UnWsClient)
	}
	v1 := r.Group("/api/v1")
	{
		v1.POST("/login", authMiddleware.LoginHandler)
		v1.GET("/refresh_token", authMiddleware.RefreshHandler)
	}
	registerBaseRouter(v1, authMiddleware)
}

func registerBaseRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.SysMenu{}
	api2 := apis.SysDept{}
	v1auth := v1.Group("").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		v1auth.GET("/roleMenuTreeselect/:roleId", api.GetMenuTreeSelect)
		v1auth.GET("/roleDeptTreeselect/:roleId",api2.GetDeptTreeRoleSelect)
		// v1auth.POST("/logout", handler.LogOut)
	}
}
