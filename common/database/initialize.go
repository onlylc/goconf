package database

import (
	"fmt"
	"goconf/common/global"
	log "goconf/core/logger"
	"goconf/core/sdk/config"
	"goconf/core/sdk/pkg"
	toolsDB "goconf/core/tools/database"
	. "goconf/core/tools/gorm/logger"
	"time"

	"goconf/core/sdk"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func Setup() {
	for k, v := range config.DatabasesConfig {
		setupSimpleDatabase(k, v)
	}

}

func setupSimpleDatabase(host string, c *config.Database) {
	if global.Driver == "" {
		global.Driver = c.Driver
	}
	
	log.Infof("%s => %s", host, pkg.Green(c.Source))
	registers := make([]toolsDB.ResolverConfigure, len(c.Registers))
	for i := range c.Registers {
		registers[i] = toolsDB.NewResolverConfigure(
			c.Registers[i].Sources,
			c.Registers[i].Replicas,
			c.Registers[i].Policy,
			c.Registers[i].Tables)
	}
	
	resolverConfig := toolsDB.NewConfigure(c.Source, c.MaxIdleConns, c.MaxOpenConns, c.ConnMaxIdleTime, c.ConnMaxLifeTime, registers)

	db, err := resolverConfig.Init(&gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: New(
			logger.Config{
				SlowThreshold: time.Second,
				Colorful:      true,
				LogLevel: logger.LogLevel(
					log.DefaultLogger.Options().Level.LevelForGorm()),
			},
		),
	}, opens[c.Driver])

	if err != nil {
		fmt.Println(pkg.Red(c.Driver+" connect error :"), err)
		log.Fatal(pkg.Red(c.Driver+" connect error :"), err)
	} else {
		fmt.Println(pkg.Green(c.Driver + " connect success !"))
		log.Info(pkg.Green(c.Driver + " connect success !"))
	}

	sdk.Runtime.SetDb(host, db)
}
