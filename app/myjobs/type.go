package myjobs

import "github.com/robfig/cron/v3"

type Job interface {
	Run()
	addJob(*cron.Cron) (error)
}

type JobExec interface {
	Exec(arg interface{}) error
}

func CallExec(e JobExec, arg interface{}) error {
	return e.Exec(arg)
}
