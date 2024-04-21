package api

import (
	"fmt"
	"goconf/app/other/router"
)

func init() {
	fmt.Println("other api")
	AppRouters = append(AppRouters, router.InitRouter)
}
