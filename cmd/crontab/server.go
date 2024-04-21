package crontab

import (
	"goconf/common/database"
	"goconf/core/config/source/file"
	"goconf/core/logger"
	"goconf/core/sdk"
	"goconf/core/sdk/config"
	"os"
	"os/signal"
	"syscall"

	"github.com/robfig/cron/v3"
	"github.com/spf13/cobra"
)

var (
	configYml string
	apiCheck  bool
	StartCmd  = &cobra.Command{
		Use:          "cron",
		Short:        "Start cron server",
		Example:      "go-admin server -c config/settings.yml",
		SilenceUsage: true,
		PreRun: func(cmd *cobra.Command, args []string) {
			setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "config/settings.yml", "Start server with provided configuration file")
	StartCmd.PersistentFlags().BoolVarP(&apiCheck, "api", "a", false, "Start server with check api data")

	//注册路由 fixme 其他应用的路由，在本目录新建文件放在init方法
	// AppRouters = append(AppRouters, router.InitRouter)
}

func setup() {
	// 注入配置扩展项
	// config.ExtendConfig = &ext.ExtConfig
	//1. 读取配置
	config.Setup(
		file.NewSource(file.WithPath(configYml)),
		database.Setup,
		// es.Setup,
	)

}

func run() error {
	var log *logger.Helper
	log = logger.NewHelper(sdk.Runtime.GetLogger()).WithFields(map[string]interface{}{
	})

	c := cron.New(cron.WithSeconds())
	
	_, err := c.AddFunc("@every 1s", func() { log.Error("name:age") })
	if err != nil {

		return err
	}

	_, err1 := c.AddFunc("@every 1s", func() { logger.Info("value:test") })
	if err1 != nil {

		return err1
	}


	c.Start()

	exitSignal := make(chan os.Signal, 1)
	signal.Notify(exitSignal, syscall.SIGINT, syscall.SIGTERM)
	<-exitSignal

	c.Stop()

	return nil
}