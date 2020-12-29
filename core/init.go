// PROPRIETARY AND CONFIDENTIAL
//
// Unauthorized copying of this file via any medium is strictly prohibited.
//
// Copyright (c) 2020 Snowplow Analytics Ltd. All rights reserved.

package core

import (
	"encoding/json"
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/makasim/sentryhook"
	log "github.com/sirupsen/logrus"
	"time"
)

// Init contains the core initialization code for each implementation
func Init() (*Config, error) {
	logLevels := map[string]log.Level{
		"debug":   log.DebugLevel,
		"info":    log.InfoLevel,
		"warning": log.WarnLevel,
		"error":   log.ErrorLevel,
		"fatal":   log.FatalLevel,
		"panic":   log.PanicLevel,
	}

	cfg, err := NewConfig()
	if err != nil {
		return nil, fmt.Errorf("Failed to build Config: %s", err.Error())
	}

	// Configure Sentry
	if cfg.Sentry.Dsn != "" {
		err := sentry.Init(sentry.ClientOptions{
			Dsn:   cfg.Sentry.Dsn,
			Debug: cfg.Sentry.Debug,
		})
		if err != nil {
			return nil, fmt.Errorf("Failed to intialize Sentry: %s", err.Error())
		}
		defer sentry.Flush(2 * time.Second)

		sentryTagsMap := map[string]string{}
		err = json.Unmarshal([]byte(cfg.Sentry.Tags), &sentryTagsMap)
		if err != nil {
			return nil, fmt.Errorf("Failed to unmarshall SENTRY_TAGS to map: %s", err.Error())
		}
		sentry.ConfigureScope(func(scope *sentry.Scope) {
			for key, value := range sentryTagsMap {
				scope.SetTag(key, value)
			}
		})

		log.AddHook(sentryhook.New([]log.Level{log.PanicLevel, log.FatalLevel, log.ErrorLevel}))
	}

	// Configure logging level
	if level, ok := logLevels[cfg.LogLevel]; ok {
		log.SetLevel(level)
	} else {
		return nil, fmt.Errorf("Supported log levels are 'debug, info, warning, error, fatal, panic' - provided: %s", cfg.LogLevel)
	}

	// TODO: Should this be configurable?
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})

	log.Debugf("Config: %+v", cfg)
	return cfg, nil
}
