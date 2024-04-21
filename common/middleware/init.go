package middleware

import (
	"goconf/common/actions"
	"goconf/core/sdk"
	jwt "goconf/core/sdk/pkg/jwtauth"

	"github.com/gin-gonic/gin"
)

const (
	JwtTokenCheck   string = "JwtToken"
	RoleCheck       string = "AuthCheckRole"
	PermissionCheck string = "PermissionAction"
)

func InitMiddleware(r *gin.Engine) {

	// 数据库链接
	r.Use(WithContextDb)
	// 日志处理
	r.Use(LoggerToFile())
	// 自定义错误处理
	r.Use(CustomError)
	r.Use(NoCache)

	// 跨域处理
	r.Use(Options)
	// Secure is a middleware function that appends security
	r.Use(Secure)

	sdk.Runtime.SetMiddleware(JwtTokenCheck, (*jwt.GinJWTMiddleware).MiddlewareFunc)
	sdk.Runtime.SetMiddleware(RoleCheck, AuthCheckRole())
	sdk.Runtime.SetMiddleware(PermissionCheck, actions.PermissionAction())

}
