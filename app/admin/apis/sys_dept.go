package apis

import (
	"fmt"
	"goconf/app/admin/models"
	"goconf/core/sdk"
	"goconf/core/sdk/api"
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
