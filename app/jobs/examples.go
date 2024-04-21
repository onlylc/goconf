package jobs

import (
	"fmt"
	"time"
)

func InitJob() {
	jobList = map[string]JobExec{
		"ExamplesOne": ExamplesOne{},
	}
}

type ExamplesOne struct {
}

func (t ExamplesOne) Exec(arg interface{}) error {
	str := time.Now().Format(timeFormat) + " [INFO] JobCore ExamplesOne exec success"
	switch arg.(type) {
	case string:
		if arg.(string) != "" {
			fmt.Println("string", arg.(string))
			fmt.Println(str, arg.(string))
		} else {
			fmt.Println("arg is nil")
			fmt.Println(str, "arg is nil")
		}
		break
	}
	return nil
}
