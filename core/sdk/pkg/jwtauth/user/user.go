package user

import (
	"fmt"
	jwt "goconf/core/sdk/pkg/jwtauth"

	"goconf/core/sdk/pkg"

	"github.com/gin-gonic/gin"
)

func ExtractClaims(c *gin.Context) jwt.MapClaims {
	claims, exists := c.Get(jwt.JwtPayloadKey)
	if !exists {
		return make(jwt.MapClaims)
	}

	return claims.(jwt.MapClaims)
}

// GetUserId 获取一个int的userId
func GetUserId(c *gin.Context) int {
	data := ExtractClaims(c)
	identity, err := data.Identity()
	if err != nil {
		fmt.Println(pkg.GetCurrentTimeStr() + " [WARING] " + c.Request.Method + " " + c.Request.URL.Path + " GetUserId 缺少 identity error: " + err.Error())
		return 0
	}

	return int(identity)
}

func GetUserIdStr(c *gin.Context) string {
	data := ExtractClaims(c)

	return data.String("identity")
}
