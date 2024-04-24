package handler

import (
	"fmt"
	"goconf/app/admin/models"
	"goconf/core/sdk/pkg"
	jwt "goconf/core/sdk/pkg/jwtauth"
	"net/http"

	"goconf/core/sdk/api"
	"goconf/core/sdk/pkg/response"

	"github.com/gin-gonic/gin"
)

func PayloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(map[string]interface{}); ok {
		u, _ := v["user"].(SysUser)
		r, _ := v["role"].(SysRole)
		return jwt.MapClaims{
			jwt.IdentityKey:  u.UserId,
			jwt.RoleIdKey:    r.RoleId,
			jwt.RoleKey:      r.RoleKey,
			jwt.NiceKey:      u.Username,
			jwt.DataScopeKey: r.DataScope,
			jwt.RoleNameKey:  r.RoleName,
		}
	}
	return jwt.MapClaims{}
}

func IdentityHandler(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	return map[string]interface{}{
		"IdentityKey": claims["identity"],
		"UserName":    claims["nice"],
		"RoleKey":     claims["rolekey"],
		"UserId":      claims["identity"],
		"RoleIds":     claims["roleid"],
		"DataScope":   claims["datascope"],
	}
}

func Authenticator(c *gin.Context) (interface{}, error) {
	log := api.GetRequestLogger(c)
	db, err := pkg.GetOrm(c)
	if err != nil {
		response.Error(c, 500, err, "数据库连接获取失败")
		return nil, jwt.ErrFailedAuthentication
	}
	var loginVals Login
	var status = "2"
	var msg = "登录成功"
	var username = ""
	defer func() {
		fmt.Println(status, msg, username)
	}()
	if err = c.ShouldBind(&loginVals); err != nil {
		username = loginVals.Username
		msg = "数据解析失败"
		status = "1"
		return nil, jwt.ErrMissingLoginValues

	}
	
	sysUser, role, e := loginVals.GetUser(db)
	if e == nil {
		username = loginVals.Username
		return map[string]interface{}{"user": sysUser, "role": role}, nil
	} else {
		msg = "登录失败"
		status = "1"
		log.Warnf("%s login failed!", loginVals.Username)

	}
	return nil, jwt.ErrFailedAuthentication

}

func Authorizator(data interface{}, c *gin.Context) bool {

	if v, ok := data.(map[string]interface{}); ok {
		u, _ := v["user"].(models.SysUser)
		r, _ := v["role"].(models.SysRole)
		c.Set("role", r.RoleName)
		c.Set("roleIds", r.RoleId)
		c.Set("userId", u.UserId)
		c.Set("userName", u.Username)
		c.Set("dataScope", r.DataScope)
		return true
	}
	return false
}

func Unauthorized(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  message,
	})
}
