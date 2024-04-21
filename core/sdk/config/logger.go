package config

import "goconf/core/sdk/pkg/logger"

type Logger struct {
	Type      string
	Path      string
	Level     string
	Stdout    string
	EnabledDB bool
	Cap       uint
}

func (e Logger) Setup() {
	logger.SetupLogger(
		logger.WithType(e.Type),
		logger.WithPath(e.Path),
		logger.WithLevel(e.Level),
		logger.WithStdout(e.Stdout),
		logger.WithCap(e.Cap),
	)
}

var LoggerConfig = new(Logger)
