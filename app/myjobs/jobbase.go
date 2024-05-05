package myjobs

import (
	"fmt"
	"goconf/app/myjobs/models"
	"goconf/core/sdk"
	"goconf/core/sdk/pkg"
	"goconf/core/sdk/pkg/cronjob"
	"os"
	"time"

	"github.com/robfig/cron/v3"
)

var timeFormat = "2006-01-02 15:04:05"

var lastModTime time.Time

var jobList map[string]JobExec

func GetList(list interface{}) {

}

type JobCore struct {
	InvokeTarget   string
	Name           string
	EntryId        int
	CronExpression string
	Args           interface{}
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

	fmt.Printf("[Job] JobCore %s exec success , spend : %v\n", e.Name, latencyTime)
	// return

}

func Setup() {
	fmt.Println(time.Now().Format(timeFormat), " [INFO] JobCore Starting...")

	key := pkg.GetCurrentTimeStr()
	sdk.Runtime.SetCrontab(key, cronjob.NewWithSeconds())
	
	setup(key)
}

func setup(key string) {
	lastModTime = time.Now()
	crontab := sdk.Runtime.GetCrontabKey(key)
	go watchConfig(crontab)
	sysJob := models.SysJob{}
	jobList := make([]models.SysJob, 0)
	err := sysJob.GetList(&jobList)
	if err != nil {
		fmt.Println(time.Now().Format(timeFormat), " [ERROR] JobCore init error", err)
	}
	if len(jobList) == 0 {
		fmt.Println(time.Now().Format(timeFormat), " [INFO] JobCore total:0")
	}
	for i := 0; i < len(jobList); i++ {
		// fmt.Println(jobList[i])
		j := &ExecJob{}
		j.InvokeTarget = jobList[i].InvokeTarget
		j.CronExpression = jobList[i].CronExpression
		j.Name = jobList[i].Name
		j.Args = jobList[i].Args
		_, _ = AddJob(crontab, j)
	}

	crontab.Start()
	fmt.Println(time.Now().Format(timeFormat), " [INFO] JobCore start success.")
	// 关闭任务
	
	
	select {}
}

func AddJob(c *cron.Cron, job Job) (int, error) {
	if job == nil {
		fmt.Println("unknow")
		return 0, nil
	}
	return job.addJob(c)
}

func (e *ExecJob) addJob(c *cron.Cron) (int, error) {
	id, err := c.AddJob(e.CronExpression, e)
	if err != nil {
		fmt.Println(time.Now().Format(timeFormat), " [ERROR] JobCore AddJob error", err)
		return 0, err
	}
	EntryId := int(id)
	return EntryId, nil
}

func watchConfig(c *cron.Cron) {

	for {
		time.Sleep(10 * time.Second) 
		if hasChanged() {
			fmt.Println("cron config file han changed, Reloading...")
			reloadJobs(c)
		}
	}
}

// hasChanged 检查配置文件是否发生了变化
func hasChanged() bool {
	// 获取当前配置文件的最新修改时间
	fileInfo, err := os.Stat("config/settings.yml")
	if err != nil {
		fmt.Println("Failed to get file info:", err)
		return false
	}
	modTime := fileInfo.ModTime()

	// 如果最新修改时间晚于上次加载配置文件的时间，则说明配置文件发生了变化
	return modTime.After(lastModTime)
}

func reloadJobs(c *cron.Cron) {
	// 清空现有的作业
	c.Stop()
	key := pkg.GetCurrentTimeStr()
	sdk.Runtime.SetCrontab(key, cronjob.NewWithSeconds())
	setup(key)
	
}