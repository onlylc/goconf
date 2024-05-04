package dto

import (
	"goconf/app/admin/models"
	"goconf/common/dto"
	common "goconf/common/models"
)

type SysApiGetPageReq struct {
	dto.Pagination `search:"-"`
	Title          string `form:"title"  search:"type:contains;column:title;table:sys_api" comment:"标题"`
	Path           string `form:"path"  search:"type:contains;column:path;table:sys_api" comment:"地址"`
	Action         string `form:"action"  search:"type:exact;column:action;table:sys_api" comment:"请求方式"`
	ParentId       string `form:"parentId"  search:"type:exact;column:parent_id;table:sys_api" comment:"按钮id"`
	Type           string `form:"type" search:"-" comment:"类型"`
	SysApiOrder
}

type SysApiOrder struct {
	TitleOrder     string `search:"type:order;column:title;table:sys_api" form:"titleOrder"`
	PathOrder      string `search:"type:order;column:path;table:sys_api" form:"pathOrder"`
	CreatedAtOrder string `search:"type:order;column:created_at;table:sys_api" form:"createdAtOrder"`
}

func (m *SysApiGetPageReq) GetNeedSearch() interface{} {
	return *m
}

// SysApiGetReq 功能获取请求参数
type SysApiGetReq struct {
	Id int `uri:"id"`
}

func (s *SysApiGetReq) GetId() interface{} {
	return s.Id
}

// SysApiUpdateReq 功能更新请求参数
type SysApiUpdateReq struct {
	Id     int    `uri:"id" comment:"编码"` // 编码
	Handle string `json:"handle" comment:"handle"`
	Title  string `json:"title" comment:"标题"`
	Path   string `json:"path" comment:"地址"`
	Type   string `json:"type" comment:""`
	Action string `json:"action" comment:"类型"`
	common.ControlBy
}

func (s *SysApiUpdateReq) Generate(model *models.SysApi) {
	if s.Id != 0 {
		model.Id = s.Id
	}
	model.Handle = s.Handle
	model.Title = s.Title
	model.Path = s.Path
	model.Type = s.Type
	model.Action = s.Action
}

func (s *SysApiUpdateReq) GetId() interface{} {
	return s.Id
}
