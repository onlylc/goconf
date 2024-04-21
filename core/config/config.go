package config

import (
	"context"
	"goconf/core/config/reader"
	"goconf/core/config/source"
	"goconf/core/config/source/file"

	"goconf/core/config/loader"
)

type Config interface {
	// Values provide the reader.Values interface
	reader.Values
	// Init the config
	Init(opts ...Option) error
	// Options in the config
	Options() Options
	// Close Stop the config loader/watcher
	Close() error
	// Load config sources
	Load(source ...source.Source) error
	// Sync Force a source changeset sync
	Sync() error
	// Watch a value for changes
	Watch(path ...string) (Watcher, error)
}

type Watcher interface {
	Next() (reader.Value, error)
	Stop() error
}

// Entity 配置实体
type Entity interface {
	OnChange()
}

type Options struct {
	Loader loader.Loader
	Reader reader.Reader
	Source []source.Source

	// for alternative data
	Context context.Context

	Entity Entity
}

type Option func(o *Options)

var (
	// DefaultConfig Default Config Manager
	DefaultConfig Config
)

func NewConfig(opts ...Option) (Config, error) {
	return newConfig(opts...)
}

// Bytes Return config as raw json
func Bytes() []byte {
	return DefaultConfig.Bytes()
}

// Map Return config as a map
func Map() map[string]interface{} {
	return DefaultConfig.Map()
}

// Scan values to a go type
func Scan(v interface{}) error {
	return DefaultConfig.Scan(v)
}

// Sync Force a source changeset sync
func Sync() error {
	return DefaultConfig.Sync()
}

// Get a value from the config
func Get(path ...string) reader.Value {
	return DefaultConfig.Get(path...)
}

// Load config sources
func Load(source ...source.Source) error {
	return DefaultConfig.Load(source...)
}

// Watch a value for changes
func Watch(path ...string) (Watcher, error) {
	return DefaultConfig.Watch(path...)
}

// LoadFile is short hand for creating a file source and loading it
func LoadFile(path string) error {
	return Load(file.NewSource(
		file.WithPath(path),
	))
}
