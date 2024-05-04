package service

import (
	"goconf/app/admin/models"
	"goconf/app/admin/service/dto"
	"goconf/common/actions"
	cDto "goconf/common/dto"
	"goconf/common/global"
	"goconf/core/sdk/service"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type SysApi struct {
	service.Service
}

// GetPage 获取SysApi 列表
func (e *SysApi) GetPage(c *dto.SysApiGetPageReq, p *actions.DataPermission, list *[]models.SysApi, count *int64) error {
	var err error
	var data models.SysApi
	orm := e.Orm.Debug().Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		)
	if c.Type != "" {
		qType := c.Type
		if qType == "暂无" {
			qType = ""
		}
		if global.Driver == "postgres" {
			orm = orm.Where("type = ?", qType)
		} else {
			orm = orm.Where("`type` = ?", qType)
		}
	}
	err = orm.Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("Service GetSysApiPage error:%s", err)
		return err
	}
	return nil
}

// Get 获取SysApi对象 with id
func (e *SysApi) Get(d *dto.SysApiGetReq, p *actions.DataPermission, model *models.SysApi) *SysApi {
	var data models.SysApi
	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权限查看")
		e.Log.Errorf("Service GetSysApi error:%s", err)
		_ = e.AddError(err)
		return e
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		_ = e.AddError(err)
		return e
	}
	return e
}

// Update 修改SysApi对象
func (e *SysApi) Update(c *dto.SysApiUpdateReq, p *actions.DataPermission) error {
    var model = models.SysApi{}
	db := e.Orm.Debug().First(&model, c.GetId())
	if db.RowsAffected == 0 {
		return errors.New("无权限更新该数据")
	}
	c.Generate(&model)
	db = e.Orm.Save(&model)
	if err := db.Error; err != nil {
		return err
	}
	return nil 
}
