package storage

import (
	"context"
	"fmt"
	"goconf/core/sdk"
	"goconf/core/sdk/config"
)

func Setup() {

	cacheAdapter, err := config.CacheConfig.Setup()
	if err != nil {
		fmt.Printf("cache setup error, %s\n", err.Error())
	}
	sdk.Runtime.SetCacheAdapter(cacheAdapter)

	r := sdk.Runtime.GetCacheAdapter()
	err = r.Set("test","testSet",300)
	if err != nil {
		fmt.Println("set err",err)
	}
	s, err := r.Get("test")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("get",s)
	
	re := r.GetClient()
	ctx := context.Background()
	// re.HSet(ctx, "myhash", "name","abc","age",22,"sex","M")
	v,_ :=re.HGetAll(ctx,"myhash").Result()
	fmt.Println(v["age"])

}
