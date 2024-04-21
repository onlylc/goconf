package service

// import (
// 	"goconf/app/admin/models"

// 	"goconf/common/dto"
// )

// type SysDept struct {
// 	server.Service
// }

// // SetDeptPage 设置dept页面数据
// func (e *SysDept) SetDeptPage(c *dto.SysDeptGetPageReq) (m []models.SysDept, err error) {
// 	var list []models.SysDept
// 	err = e.getList(c, &list)
// 	for i := 0; i < len(list); i++ {
// 		if list[i].ParentId != 0 {
// 			continue
// 		}
// 		info := e.deptPageCall(&list, list[i])
// 		m = append(m, info)
// 	}
// 	return
// }

// // GetSysDeptList 获取组织数据
// func (e *SysDept) getList(c *dto.SysDeptGetPageReq, list *[]models.SysDept) error {
// 	var err error
// 	var data models.SysDept

// 	err = e.Orm.Model(&data).
// 		Scopes(
// 			cDto.MakeCondition(c.GetNeedSearch()),
// 		).
// 		Find(list).Error
// 	if err != nil {
// 		e.Log.Errorf("db error:%s", err)
// 		return err
// 	}
// 	return nil
// }

// func (e *SysDept) deptPageCall(deptlist *[]models.SysDept, menu models.SysDept) models.SysDept {
// 	list := *deptlist
// 	min := make([]models.SysDept, 0)
// 	for j := 0; j < len(list); j++ {
// 		if menu.DeptId != list[j].ParentId {
// 			continue
// 		}
// 		mi := models.SysDept{}
// 		mi.DeptId = list[j].DeptId
// 		mi.ParentId = list[j].ParentId
// 		mi.DeptPath = list[j].DeptPath
// 		mi.DeptName = list[j].DeptName
// 		mi.Sort = list[j].Sort
// 		mi.Leader = list[j].Leader
// 		mi.Phone = list[j].Phone
// 		mi.Email = list[j].Email
// 		mi.Status = list[j].Status
// 		mi.CreatedAt = list[j].CreatedAt
// 		mi.Children = []models.SysDept{}
// 		ms := e.deptPageCall(deptlist, mi)
// 		min = append(min, ms)
// 	}
// 	menu.Children = min
// 	return menu
// }
