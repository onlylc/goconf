package service

import (
	"fmt"
	"goconf/app/admin/models"
	"goconf/app/admin/service/dto"
	cModels "goconf/common/models"
	"goconf/core/sdk/service"
	cDto "goconf/common/dto"
	"sort"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type SysMenu struct {
	service.Service
}

// menuCall 构建菜单树
func menuCall(menuList *[]models.SysMenu, menu models.SysMenu) models.SysMenu {
	list := *menuList

	min := make([]models.SysMenu, 0)
	for j := 0; j < len(list); j++ {

		if menu.MenuId != list[j].ParentId {
			continue
		}
		mi := models.SysMenu{}
		mi.MenuId = list[j].MenuId
		mi.MenuName = list[j].MenuName
		mi.Title = list[j].Title
		mi.Icon = list[j].Icon
		mi.Path = list[j].Path
		mi.MenuType = list[j].MenuType
		mi.Action = list[j].Action
		mi.Permission = list[j].Permission
		mi.ParentId = list[j].ParentId
		mi.NoCache = list[j].NoCache
		mi.Breadcrumb = list[j].Breadcrumb
		mi.Component = list[j].Component
		mi.Sort = list[j].Sort
		mi.Visible = list[j].Visible
		mi.CreatedAt = list[j].CreatedAt
		mi.SysApi = list[j].SysApi
		mi.Children = []models.SysMenu{}

		if mi.MenuType != cModels.Button {
			ms := menuCall(menuList, mi)
			min = append(min, ms)
		} else {
			min = append(min, mi)
		}
	}
	menu.Children = min
	return menu
}

// SetMenuRole 获取左侧菜单树使用
func (e *SysMenu) SetMenuRole(roleName string) (m []models.SysMenu, err error) {
	menus, err := e.getByRoleName(roleName)
	m = make([]models.SysMenu, 0)
	for i := 0; i < len(menus); i++ {
		if menus[i].ParentId != 0 {
			continue
		}
		menusInfo := menuCall(&menus, menus[i])
		m = append(m, menusInfo)
	}
	return
}

func menuDistinct(menuList []models.SysMenu) (result []models.SysMenu) {
	distinctMap := make(map[int]struct{}, len(menuList))
	for _, menu := range menuList {
		if _, ok := distinctMap[menu.MenuId]; !ok {
			distinctMap[menu.MenuId] = struct{}{}
			result = append(result, menu)
		}
	}
	return result
}

func recursiveSetMenu(orm *gorm.DB, mIds []int, menus *[]models.SysMenu) error {
	if len(mIds) == 0 || menus == nil {
		return nil
	}
	var subMenus []models.SysMenu
	err := orm.Where(fmt.Sprintf(" menu_type in ('%s', '%s', '%s') and menu_id in ?",
		cModels.Directory, cModels.Menu, cModels.Button), mIds).Order("sort").Find(&subMenus).Error
	if err != nil {
		return err
	}

	subIds := make([]int, 0)
	for _, menu := range subMenus {
		if menu.ParentId != 0 {
			subIds = append(subIds, menu.ParentId)
		}
		if menu.MenuType != cModels.Button {
			*menus = append(*menus, menu)
		}
	}
	return recursiveSetMenu(orm, subIds, menus)
}

func (e *SysMenu) getByRoleName(roleName string) ([]models.SysMenu, error) {
	var role models.SysRole
	var err error
	data := make([]models.SysMenu, 0)

	if roleName == "admin" {
		err = e.Orm.Where(" menu_type in ('M','C') and deleted_at is null").
			Order("sort").
			Find(&data).
			Error
		err = errors.WithStack(err)
	} else {
		role.RoleKey = roleName
		err = e.Orm.Model(&role).Where("role_key = ? ", roleName).Preload("SysMenu").First(&role).Error

		if role.SysMenu != nil {
			mIds := make([]int, 0)
			for _, menu := range *role.SysMenu {
				mIds = append(mIds, menu.MenuId)
			}
			if err := recursiveSetMenu(e.Orm, mIds, &data); err != nil {
				return nil, err
			}
			data = menuDistinct(data)
		}

	}

	sort.Sort(models.SysMenuSlice(data))
	return data, err

}

// GetList 获取菜单数据
func (e *SysMenu) GetList(c *dto.SysMenuGetPageReq, list *[]models.SysMenu) error {
	var err error
	var data models.SysMenu

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

// SetLabel 修改角色中 设置菜单基础数据
func (e *SysMenu) SetLabel() (m []dto.MenuLabel, err error) {
	var list []models.SysMenu
	err = e.GetList(&dto.SysMenuGetPageReq{}, &list)

	m = make([]dto.MenuLabel, 0)
	for i := 0; i < len(list); i++ {
		if list[i].ParentId != 0 {
			continue
		}
		e := dto.MenuLabel{}
		e.Id = list[i].MenuId
		e.Label = list[i].Title
		deptsInfo := menuLabelCall(&list, e)

		m = append(m, deptsInfo)
	}
	return
}

// menuLabelCall 递归构造组织数据
func menuLabelCall(eList *[]models.SysMenu, dept dto.MenuLabel) dto.MenuLabel {
	list := *eList

	min := make([]dto.MenuLabel, 0)
	for j := 0; j < len(list); j++ {

		if dept.Id != list[j].ParentId {
			continue
		}
		mi := dto.MenuLabel{}
		mi.Id = list[j].MenuId
		mi.Label = list[j].Title
		mi.Children = []dto.MenuLabel{}
		if list[j].MenuType != "F" {
			ms := menuLabelCall(eList, mi)
			min = append(min, ms)
		} else {
			min = append(min, mi)
		}
	}
	if len(min) > 0 {
		dept.Children = min
	} else {
		dept.Children = nil
	}
	return dept
}

