package pkg

type (
	Mode string
)

const (
	ModeDev  Mode = "dev"
    ModeTest Mode = "test"
    ModeProd Mode = "prod"
	Mysql = "mysql"
	Sqlite = "sqlite"
)

func (e Mode) String() string { 
	return string(e)
}
