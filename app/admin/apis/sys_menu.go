package apis

import (
	"fmt"
	"goconf/app/admin/service"
	"goconf/app/admin/service/dto"
	"goconf/core/sdk/api"
	"goconf/core/sdk/pkg/jwtauth/user"
	"github.com/gin-gonic/gin"
)

type SysMenu struct {
	api.Api
}

func (e SysMenu) GetMenuTreeSelect(c *gin.Context) {
	m := service.SysMenu{}
	r := service.SysRole{}
	req := dto.SelectRole{}
	// fmt.Println("req.RoleID",req.RoleId)
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&m.Service).
		MakeService(&r.Service).
		Bind(&req, nil).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	// fmt.Println("---------------------------",req.RoleId,m)
	result, err := m.SetLabel()
	// fmt.Println("---------------------------",req.RoleId,m)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}

	menuIds := make([]int, 0)
	if req.RoleId != 0 {
		menuIds, err = r.GetRoleMenuId(req.RoleId)
		fmt.Println("menuIds",menuIds)
		if err != nil {
			e.Error(500, err, "")
			return
		}
	}

	e.OK(gin.H{
		"menus":       result,
		"checkedKeys": menuIds,
	}, "获取成功")
}

// GetMenuRole 根据登录角色名称获取菜单列表数据（左菜单使用）
// @Summary 根据登录角色名称获取菜单列表数据（左菜单使用）
// @Description 获取JSON
// @Tags 菜单
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/menurole [get]
// @Security Bearer
func (e SysMenu) GetMenuRole(c *gin.Context) {
	s := new(service.SysMenu)
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	result, err := s.SetMenuRole(user.GetRoleName(c))

	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}

	e.OK(result, "")
}