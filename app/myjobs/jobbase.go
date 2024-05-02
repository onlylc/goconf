package myjobs

import (
	"fmt"
	"goconf/app/myjobs/models"
	"goconf/core/sdk"
	"goconf/core/sdk/pkg/cronjob"
	"time"

	"github.com/robfig/cron/v3"
)

var timeFormat = "2006-01-02 15:04:05"

var jobList map[string]JobExec

func GetList(list interface{}) {

}

type JobCore struct {
	InvokeTarget   string
	Name           string
	CronExpression string
	Args           string
}

type ExecJob struct {
	JobCore
}

func (e *ExecJob) Run() {
	startTime := time.Now()
	var obj = jobList[e.InvokeTarget]
	if obj == nil {
		fmt.Println("[Job] ExecJob Run job nil")
		return
	}
	err := CallExec(obj, e.Args)
	if err != nil {
		fmt.Println(time.Now().Format(timeFormat), " [ERROR] mission failed! ", err)
	}
	endTime := time.Now()

	latencyTime := endTime.Sub(startTime)

	fmt.Printf("[Job] JobCore %s exec success , spend : %v\n",e.Name, latencyTime)
	// return

}

func Setup() {
	fmt.Println(time.Now().Format(timeFormat), " [INFO] JobCore Starting...")

	sdk.Runtime.SetCrontab("*", cronjob.NewWithSeconds())
	setup("*")
}

func setup(key string) {
	crontab := sdk.Runtime.GetCrontabKey(key)
	fmt.Println("crontab")
	sysJob := models.SysJob{}
	jobList := make([]models.SysJob, 0)
	err := sysJob.GetList(&jobList)
	fmt.Println(jobList)
	if err != nil {
		fmt.Println(time.Now().Format(timeFormat), " [ERROR] JobCore init error", err)
	}
	if len(jobList) == 0 {
		fmt.Println(time.Now().Format(timeFormat), " [INFO] JobCore total:0")
	}
	for i := 0; i < len(jobList); i++ {
		j := &ExecJob{}
		j.InvokeTarget = jobList[i].InvokeTarget
		j.CronExpression = jobList[i].CronExpression
		j.Args = jobList[i].Args
		j.Name = jobList[i].Name
		_ = AddJob(crontab, j)
	}
	crontab.Start()
	fmt.Println(time.Now().Format(timeFormat), " [INFO] JobCore start success.")
	// 关闭任务
	defer crontab.Stop()
	select {}
}

func AddJob(c *cron.Cron, job Job) error {
	if job == nil {
		fmt.Println("unknow")
		return nil
	}
	return job.addJob(c)
}

func (e *ExecJob) addJob(c *cron.Cron) error {
	_, err := c.AddJob(e.CronExpression, e)

	if err != nil {
		fmt.Println(time.Now().Format(timeFormat), " [ERROR] JobCore AddJob error", err)
		return err
	}
	return nil
}
