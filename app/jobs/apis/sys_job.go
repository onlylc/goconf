package apis

import (
	"goconf/app/jobs/service"
	"goconf/common/dto"
	"goconf/core/sdk"
	"goconf/core/sdk/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SysJob struct {
	api.Api
}

func (e SysJob) RemoveJobForService(c *gin.Context) {
	v := dto.GeneralDelDto{}
	s := service.SysJob{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&v).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		return
	}
	s.Cron = sdk.Runtime.GetCrontabKey(c.Request.Host)
	err = s.RemoveJob(&v)
	if err != nil {
		e.Logger.Errorf("RemoveJob error, %s", err.Error())
		e.Error(500, err, "")
		return
	}
	e.OK(nil, s.Msg)
}

func (e SysJob) StartJobForService(c *gin.Context) {
	e.MakeContext(c)
	log := e.GetLogger()
	db, err := e.GetOrm()
	if err != nil {
		log.Error(err)
		return
	}
	var v dto.GeneralGetDto
	err = c.BindUri(&v)
	if err != nil {
		log.Warnf("参数验证错误, error: %s", err)
		e.Error(http.StatusUnprocessableEntity, err, "参数验证失败")
		return
	}
	s := service.SysJob{}
	s.Orm = db
	s.Log = log
	s.Cron = sdk.Runtime.GetCrontabKey(c.Request.Host)
	err = s.StartJob(&v)
	if err != nil {
		log.Errorf("GetCrontabKey error, %s", err.Error())
		e.Error(500, err, err.Error())
		return
	}
	e.OK(nil, s.Msg)
}
