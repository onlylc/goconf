package router

import (
	"github.com/gin-gonic/gin"
	"goconf/app/other/apis"
)
func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerFileRouter)
}

func registerFileRouter(v1 *gin.RouterGroup) {
	var api = apis.File{}
	r := v1.Group("")
	{
		r.POST("/public/uploadFile", api.UploadFile)
	}
}
