package service

import (
	"fmt"
	"goconf/app/admin/models"

	"goconf/app/admin/service/dto"
	cDto "goconf/common/dto"
	log "goconf/core/logger"
	"goconf/core/sdk/service"
)

type SysDept struct {
	service.Service
}

// SetDeptPage 设置dept页面数据
func (e *SysDept) SetDeptPage(c *dto.SysDeptGetPageReq) (m []models.SysDept, err error) {
	var list []models.SysDept
	err = e.getList(c, &list)
	for i := 0; i < len(list); i++ {
		if list[i].ParentId != 0 {
			continue
		}
		info := e.deptPageCall(&list, list[i])
		m = append(m, info)
	}
	return
}

// GetSysDeptList 获取组织数据
func (e *SysDept) getList(c *dto.SysDeptGetPageReq, list *[]models.SysDept) error {
	var err error
	var data models.SysDept

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
		).
		Find(list).Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

func (e *SysDept) deptPageCall(deptlist *[]models.SysDept, menu models.SysDept) models.SysDept {
	list := *deptlist
	min := make([]models.SysDept, 0)
	for j := 0; j < len(list); j++ {
		if menu.DeptId != list[j].ParentId {
			continue
		}
		mi := models.SysDept{}
		mi.DeptId = list[j].DeptId
		mi.ParentId = list[j].ParentId
		mi.DeptPath = list[j].DeptPath
		mi.DeptName = list[j].DeptName
		mi.Sort = list[j].Sort
		mi.Leader = list[j].Leader
		mi.Phone = list[j].Phone
		mi.Email = list[j].Email
		mi.Status = list[j].Status
		mi.CreatedAt = list[j].CreatedAt
		mi.Children = []models.SysDept{}
		ms := e.deptPageCall(deptlist, mi)
		min = append(min, ms)
	}
	menu.Children = min
	return menu
}

func (e *SysDept) SetDeptLabel() (m []dto.DeptLabel, err error) {
	list := make([]models.SysDept, 0)
	err = e.Orm.Find(&list).Error
	if err != nil {
		log.Error("find dept list error, %s", err.Error())
		return
	}
	// fmt.Println(list)
	m = make([]dto.DeptLabel, 0)
	var item dto.DeptLabel
	for i := range list {
		if list[i].ParentId != 0 {
			continue
		}
	
		item = dto.DeptLabel{}
		item.Id = list[i].DeptId
		item.Label = list[i].DeptName
		fmt.Println(item,list[i])

		deptInfo := deptLabelCall(&list, item)

		m = append(m, deptInfo)
		// fmt.Println(m)
	}
	return
}

func deptLabelCall(deptList *[]models.SysDept, dept dto.DeptLabel) dto.DeptLabel {
	// fmt.Println("----dept",dept)
	// list := *deptList
	// children := make([]dto.DeptLabel, 0)

	// for _, d := range list {
		
	// 	// fmt.Println(dept,d)
	// 	if dept.Id == d.ParentId {
	// 		fmt.Println("range dept",dept,d)
	// 		child := dto.DeptLabel{
	// 			Id:       d.DeptId,
	// 			Label:    d.DeptName,
	// 			Children: []dto.DeptLabel{deptLabelCall(deptList, dto.DeptLabel{Id: d.DeptId, Label: d.DeptName})},
	// 		}
			

	// 		// fmt.Println("child",dept)
			
	// 		children = append(children, child)
	// 	}
		
	// }

	// dept.Children = children
	// return dept

	list := *deptList
	var mi dto.DeptLabel
	min := make([]dto.DeptLabel, 0)
	for j := 0; j < len(list); j++ {
		
		if dept.Id != list[j].ParentId {
			continue
		}
		fmt.Println(dept, list[j])
		mi = dto.DeptLabel{Id: list[j].DeptId, Label: list[j].DeptName, Children: []dto.DeptLabel{}}
		fmt.Println(mi)
		ms := deptLabelCall(deptList, mi)
		fmt.Println("ms",ms)
		min = append(min, ms)
	}
	dept.Children = min
	return dept
}


func (e *SysDept) GetWithRoleId(roleId int) ([]int, error) {
	deptIds := make([]int, 0)
	deptList := make([]dto.DeptIdList, 0)
	if err := e.Orm.Table("sys_role_dept").
		Select("sys_role_dept.dept_id").
		Joins("LEFT JOIN sys_dept on sys_dept.dept_id=sys_role_dept.dept_id").
		Where("role_id = ?", roleId).
		Where(" sys_role_dept.dept_id not in(select sys_dept.parent_id from sys_role_dept LEFT JOIN sys_dept on sys_dept.dept_id=sys_role_dept.dept_id where role_id =? )", roleId).
		Find(&deptList).Error; err != nil {
		return nil, err
	}
	for i := 0; i < len(deptList); i++ {
		deptIds = append(deptIds, deptList[i].DeptId)
	}
	return deptIds, nil
}
