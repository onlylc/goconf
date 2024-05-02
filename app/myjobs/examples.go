package myjobs

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
	str := time.Now().Format(timeFormat) + " [INFO] jobCore ExamplesOne exec succes"
	switch arg := arg.(type) {
	case string:
		if arg != "" {
			fmt.Println("string", arg)
			fmt.Println(str, arg)
		} else {
			fmt.Println("arg is nil")
			fmt.Println(str, "arg is nil")
		}
	}
	return nil
}
