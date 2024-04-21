package api

import (
	"context"
	"errors"
	"fmt"
	"goconf/app/jobs"
	"goconf/app/admin/router"
	"goconf/common/database"
	common "goconf/common/middleware"
	"goconf/common/middleware/handler"
	"goconf/core/config/source/file"
	"goconf/core/sdk"
	"goconf/core/sdk/config"
	"goconf/core/sdk/pkg"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"goconf/core/sdk/api"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var (
	configYml string
	apiCheck  bool
	StartCmd  = &cobra.Command{
		Use:          "server",
		Short:        "Start API server",
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

var AppRouters = make([]func(), 0)

func init() {
	StartCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "config/settings.yml", "Start server with provided configuration file")
	StartCmd.PersistentFlags().BoolVarP(&apiCheck, "api", "a", false, "Start server with check api data")
	fmt.Println("server init")
	AppRouters = append(AppRouters, router.InitRouter)
}

func setup() {
	config.Setup(
		file.NewSource(file.WithPath(configYml)),
		database.Setup,
		// storage.Setup,
	)

	usageStr := `starting api server...`
	log.Println(usageStr)
}

func run() error {
	if config.ApplicationConfig.Mode == pkg.ModeProd.String() {
		gin.SetMode(gin.ReleaseMode)
	}
	initRouter()

	for _, f := range AppRouters {
		f()
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", config.ApplicationConfig.Host, config.ApplicationConfig.Port),
		Handler: sdk.Runtime.GetEngine(),
	}

	go func() {
		jobs.InitJob()
		jobs.Setup(sdk.Runtime.GetDb())
	}()

	go func() {
		if config.SslConfig.Enable {
			if err := srv.ListenAndServeTLS(config.SslConfig.Pem, config.SslConfig.KeyStr); err != nil && errors.Is(err, http.ErrServerClosed) {
				log.Fatal("listen: ", err)
			}

		} else {
			if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
				log.Fatal("listen: ", err)
			}
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	fmt.Printf("%s Shutdown Server ... \r\n", pkg.GetCurrentTimeStr())
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")

	return nil
}

func initRouter() {
	var r *gin.Engine
	h := sdk.Runtime.GetEngine()
	if h == nil {
		h = gin.New()
		sdk.Runtime.SetEngine(h)
	}
	switch h.(type) {
	case *gin.Engine:
		r = h.(*gin.Engine)
	default:
		log.Fatal("not support other engine")
	}
	if config.SslConfig.Enable {
		r.Use(handler.TlsHandler())
	}
	r.Use(common.RequestId(pkg.TrafficKey)).
		Use(api.SetRequestLogger)

	common.InitMiddleware(r)
}
