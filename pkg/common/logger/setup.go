package logger

import (
	"github.com/ianfedev/civicspot-backend/pkg/common/config"
)

// SetupEnvironmentLogger creates a logger from environment variables
func SetupEnvironmentLogger() {

	env := config.MustGet(config.Env)
	lvl := config.MustGet(config.LogLevel)
	enc := config.MustGet(config.LogLevelFormat)
	cal := config.MustGet(config.LogLevelCaller)
	stack := config.MustGet(config.LogLevelStack)

	lCfg := Config{
		Env:           config.EnvType(env),
		Level:         lvl,
		Encoding:      enc,
		DisableCaller: config.ParseBool(cal),
		DisableStack:  config.ParseBool(stack),
	}

	Init(lCfg)

}
