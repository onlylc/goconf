package models

type SysJob struct {
	InvokeTarget   string
	CronExpression string
	Name           string
	Age            int
	Sex            int
	Args           string
}

func (e *SysJob) GetList(list *[]SysJob) (err error) {
	data := []SysJob{
		{InvokeTarget: "ExamplesOne", CronExpression: "*/6 * * * * *", Name: "google", Age: 22, Sex: 1, Args: "cront1"},
		{InvokeTarget: "ExamplesOne", CronExpression: "*/6 * * * * *", Name: "baidu", Age: 21, Sex: 1, Args: "cront2"},
		{InvokeTarget: "ExamplesOne", CronExpression: "*/6 * * * * *", Name: "qq", Age: 20, Sex: 1, Args: "cront3"},
	}
	*list = data
	return nil
}
