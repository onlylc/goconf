package api

import (
	"fmt"
	"goconf/app/jobs/router"
)

func init() {
	fmt.Println("job api")
	AppRouters = append(AppRouters, router.InitRouter)
}
