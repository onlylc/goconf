package models

import (
	"goconf/core/sdk/config"
)

type SysJob struct {
	InvokeTarget   string
	CronExpression string
	Name           string
	Args           interface{}
	EntryId        int
}

func (e *SysJob) GetList(list *[]SysJob) (err error) {
	// 创建一个用于存储 SysJob 结构体的临时切片
	var data []SysJob

	// 遍历配置中的数据，并创建对应的 SysJob 结构体
	for _, v := range config.MyCronConfig.Data {
		job := &SysJob{
			InvokeTarget:   v.InvokeTarget,
			CronExpression: v.CronExpression,
			Name:           v.Name,
			Args:           v.Args,
		}
		// 将创建的 SysJob 结构体追加到临时切片中

		data = append(data, *job)
	}
	// 将临时切片赋值给 list 参数，返回给调用者
	*list = data
	return nil
}
