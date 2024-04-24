package router

import (
	common "goconf/common/middleware"
	"goconf/core/sdk"
	log "goconf/core/logger"
	"os"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	var r *gin.Engine
	h := sdk.Runtime.GetEngine()
	if h == nil {
		log.Fatal("not found engine...")
		os.Exit(-1)
	}
	switch h.(type) {
	case *gin.Engine:
		r = h.(*gin.Engine)
	default:
		log.Fatal("not support other engine")
		os.Exit(-1)
	}
	// the jwt middleware
	authMiddleware, err := common.AuthInit()
	if err != nil {
		log.Fatalf("JWT Init Error, %s", err.Error())
	}

	// 注册业务路由
	// TODO: 这里可存放业务路由，里边并无实际路由只有演示代码
	InitSysRouter(r, authMiddleware)
	InitExamplesRouter(r, authMiddleware)
}
