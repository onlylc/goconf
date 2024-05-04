package config

type Args struct {
	Hot   int
	Size  int
	Gt    int
	Lt    int
	Index string
	Plat  string
}

type CronJob struct {
	Name           string
	CronExpression string
	InvokeTarget   string
	Args           Args
}

type MyCron struct {
	Data []CronJob
}



var MyCronConfig = new(MyCron)
