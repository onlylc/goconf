package service

import (
	"errors"
	"goconf/app/admin/models"
	"goconf/app/admin/service/dto"
	"goconf/common/actions"
	"goconf/core/sdk/service"

	"gorm.io/gorm"
)

type SysUser struct {
	service.Service
}

// Get 获取SysUser 对象

func (e *SysUser) Get(d *dto.SysUserById, p *actions.DataPermission, model *models.SysUser) error {
	var data models.SysUser

	err := e.Orm.Model(&data).Debug().
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权限查看")
		e.Log.Error("db error: %s", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err 
	}
	return nil 
}
