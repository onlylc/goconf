package myjobs

import (
	"fmt"
	"goconf/core/sdk/config"
)


func InitJob() {
	jobList = map[string]JobExec{
		"ExamplesOne": ExamplesOne{},
		"mysql": Mysql{},

	}
}

type ExamplesOne struct {
}

func (t ExamplesOne) Exec(args interface{}) error {
	argStruct, ok := args.(config.Args)	
	if !ok {
		return fmt.Errorf("invalid argument type: %T", argStruct)
	}
	if argStruct.Index == "" {
		fmt.Println("nil",argStruct)
	}
	fmt.Println(argStruct)
	return nil
}

type Mysql struct {
}

func (t Mysql) Exec(args interface{}) error {
	argStruct, ok := args.(config.Args)	
	if !ok {
		return fmt.Errorf("invalid argument type: %T", argStruct)
	}

	fmt.Println("Mysql", argStruct)
	return nil
}