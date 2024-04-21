package config

import (
	"encoding/json"
	"fmt"
	"goconf/common/database"
	"goconf/core/config/source/file"
	log "goconf/core/logger"
	"goconf/core/sdk/config"

	"github.com/spf13/cobra"
)

var (
	configYml string
	StartCmd  = &cobra.Command{
		Use:     "config",
		Short:   "Get Application config info",
		Example: "go-admin config -c config/settings.yml",
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "config/settings.yml", "Start server with provided configuration file")
}

func run() {
	config.Setup(
		file.NewSource(file.WithPath(configYml)),
		database.Setup,
		// storage.Setup,
	)

	// log := logger.NewHelper(logger.DefaultLogger).
	// 	WithFields(map[string]interface{}{
	// 		"log": "test",
	// 	})

	application, errs := json.MarshalIndent(config.ApplicationConfig, "", "   ") //转换成JSON返回的是byte[]
	if errs != nil {
		fmt.Println(errs.Error())
	}
	log.Error("application:", string(application))

	// jwt, errs := json.MarshalIndent(config.JwtConfig, "", "   ") //转换成JSON返回的是byte[]
	// if errs != nil {
	// 	fmt.Println(errs.Error())
	// }
	// fmt.Println("jwt:", string(jwt))

	// // todo 需要兼容
	// database, errs := json.MarshalIndent(config.DatabasesConfig, "", "   ") //转换成JSON返回的是byte[]
	// if errs != nil {
	// 	fmt.Println(errs.Error())
	// }
	// fmt.Println("database:", string(database))

	// gen, errs := json.MarshalIndent(config.GenConfig, "", "   ") //转换成JSON返回的是byte[]
	// if errs != nil {
	// 	fmt.Println(errs.Error())
	// }
	// fmt.Println("gen:", string(gen))

	// loggerConfig, errs := json.MarshalIndent(config.LoggerConfig, "", "   ") //转换成JSON返回的是byte[]
	// if errs != nil {
	// 	fmt.Println(errs.Error())
	// }
	// fmt.Println("logger:", string(loggerConfig))
	// redisConfig, errs := json.MarshalIndent(config.CacheConfig, "", "   ") //转换成JSON返回的是byte[]
	// if errs != nil {
	// 	fmt.Println(errs.Error())
	// }
	// fmt.Println("cache:", string(redisConfig))
	log.Debug("DDDDDDDDDDDD")
	select {}

}
