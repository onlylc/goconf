package service

import (
	"goconf/app/admin/models"
	"goconf/core/sdk/service"
)

type SysRole struct {
	service.Service
}

func (e *SysRole) GetRoleMenuId(roleId int) ([]int, error) {
	menuIds := make([]int, 0)
	model := models.SysRole{}
	model.RoleId = roleId 
	if err := e.Orm.Model(&model).Preload("SysMenu").First(&model).Error; err != nil {
		return nil, err 
	}
	l := *model.SysMenu
	for i := 0; i <len(l); i++ {
		menuIds = append(menuIds, l[i].MenuId)
	}
	return menuIds, nil 
}

func (e *SysRole) GetById(roleId int) ([]string, error) {
	permissions := make([]string, 0)
	model := models.SysRole{}
	model.RoleId = roleId 
	if err := e.Orm.Model(&model).Preload("SysMenu").First(&model).Error; err != nil {
		return nil, err 
	}
	l := *model.SysMenu
	for i := 0; i < len(l); i++ {
		if l[i].Permission != "" {
			permissions = append(permissions, l[i].Permission)
		}
	}
	return permissions, nil
}