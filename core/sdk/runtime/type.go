package runtime

import (
	"goconf/core/logger"
	"goconf/core/storage"
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
)

type Runtime interface {
	SetDb(key string, db *gorm.DB)
	GetDb() map[string]*gorm.DB
	GetDbByKey(key string) *gorm.DB

	SetEsDb(key string, esDb *elasticsearch.Client)
	GetEsDb() map[string]*elasticsearch.Client
	GetEsDbByKey(key string) *elasticsearch.Client

	// SetCacheAdapter cache
	SetCacheAdapter(storage.AdapterCache)
	GetCacheAdapter() storage.AdapterCache
	GetCachePrefix(string) storage.AdapterCache

	GetMemoryQueue(string) storage.AdapterQueue
	SetQueueAdapter(storage.AdapterQueue)
	GetQueueAdapter() storage.AdapterQueue
	GetQueuePrefix(string) storage.AdapterQueue

	// SetEngine 使用的路由
	SetEngine(engine http.Handler)
	GetEngine() http.Handler

	// SetLogger 使用go-admin定义的logger，参考来源go-micro
	SetLogger(logger logger.Logger)
	GetLogger() logger.Logger

	// SetMiddleware middleware
	SetMiddleware(string, interface{})
	GetMiddleware() map[string]interface{}
	GetMiddlewareKey(key string) interface{}

	SetCasbin(key string, enforcer *casbin.SyncedEnforcer)
	GetCasbin() map[string]*casbin.SyncedEnforcer
	GetCasbinKey(key string) *casbin.SyncedEnforcer

	SetCrontab(key string, crontab *cron.Cron)
	GetCrontab() map[string]*cron.Cron
	GetCrontabKey(key string) *cron.Cron
}
