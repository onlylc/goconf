package dto

import (
	"goconf/app/admin/models"
	"goconf/common/dto"

	common "goconf/common/models"
)

type SysUserById struct {
	dto.ObjectById
	common.ControlBy
}

func (s *SysUserById) GetId() interface{} {
	if len(s.Ids) > 0 {
		s.Ids = append(s.Ids, s.Id)
		return s.Ids
	}
	return s.Id
}

func (s *SysUserById) GenerateM() (common.ActiveRecord, error) {
	return &models.SysUser{}, nil
}
