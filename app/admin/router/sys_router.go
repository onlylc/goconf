package router

import (
	"goconf/core/sdk/config"
	jwt "goconf/core/sdk/pkg/jwtauth"

	"goconf/app/admin/apis"
	"goconf/common/middleware/handler"
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
	// go ws.WebsocketManager.Start()
	// go ws.WebsocketManager.SendService()
	// go ws.WebsocketManager.SendAllService()

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

	v1 := r.Group("/api/v1")
	{

	}
	registerBaseRouter(v1, authMiddleware)
}

func registerBaseRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	// wss := r.Group("")
	// {
	// 	wss.GET("/ws/:id/:channel", ws.WebsocketManager.WsClient)
	// 	wss.GET("wslogout/:id/:channel", ws.WebsocketManager.UnWsClient)
	// }

	// v1auth := v1.Group("")
	// {
	// 	v1auth.GET("/loggingout", func(c *gin.Context) {
	// 		c.JSON(200, gin.H{
	// 			"message": "pong",
	// 		})
	// 	})
	// }
}
