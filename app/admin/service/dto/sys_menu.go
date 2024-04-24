package dto

import "goconf/common/dto"

type SysMenuGetPageReq struct {
	dto.Pagination `search:"-"`
	Tile           string `form:"title" search:"type:contains;column:title;table:sys_menu" comment:"菜单名单"`
	Visible        int    `form:"visible" search:"type:exact;column:visible;table:sys_menu" comment:"显示状态"`
}

func (m *SysMenuGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type MenuLabel struct {
	Id       int         `json:"id,omitempty" gorm:"-"`
	Label    string      `json:"label,omitempty" gorm:"-"`
	Children []MenuLabel `json:"children,omitempty" gorm:"-"`
}

type SelectRole struct {
	RoleId int `uri:"roleId"`
}
