package apis

import (
	"fmt"
	"goconf/app/admin/models"
	"goconf/app/admin/service"
	"goconf/core/sdk"
	"goconf/core/sdk/api"
	"goconf/core/sdk/pkg"

	"github.com/gin-gonic/gin"
)

type SysDept struct {
	api.Api
}

func GetPage() {

	db := sdk.Runtime.GetDbByKey("*")
	if db == nil {
		fmt.Println("Database connection is nil")
		return
	}

	var data models.SysDept

	if err := db.Where("dept_id = ?", 1).Find(&data).Error; err != nil {
		fmt.Println("查询数据时发生错误:", err)
		return
	}

	fmt.Println(data.DeptName)

}

func (e SysDept) GetDeptTreeRoleSelect(c *gin.Context) {
	s := service.SysDept{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	id, err := pkg.StringToInt(c.Param("roleId"))
	result, err := s.SetDeptLabel()
	if err != nil {
		e.Error(500, err, err.Error())
		return
	}
	menuIds := make([]int, 0)
	if id != 0 {
		menuIds, err = s.GetWithRoleId(id)
		if err != nil {
			e.Error(500, err, err.Error())
			return
		}
	}
	e.OK(gin.H{
		"depts":       result,
		"checkedKeys": menuIds,
	}, "")
}
