// PROPRIETARY AND CONFIDENTIAL
//
// Unauthorized copying of this file via any medium is strictly prohibited.
//
// Copyright (c) 2020 Snowplow Analytics Ltd. All rights reserved.

package core

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
)

// PubSubConfig configures the destination for records consumed
type PubSubConfig struct {
	ProjectID         string
	TopicName         string
	ServiceAccountB64 string
}

// KinesisConfig configures the destination for records consumed
type KinesisConfig struct {
	StreamName string
	Region     string
	RoleARN    string
}

// SentryConfig configures the Sentry error tracker
type SentryConfig struct {
	Dsn   string
	Tags  string
	Debug bool
}

// Config for holding all configuration details
type Config struct {
	Target   string
	LogLevel string
	Sentry   SentryConfig
	Kinesis  KinesisConfig
	PubSub   PubSubConfig
}

// NewConfig resolves the config from the environment
func NewConfig() *Config {
	var defaultConfig = &Config{
		PubSub: PubSubConfig{
			ProjectID:         "",
			TopicName:         "",
			ServiceAccountB64: "",
		},
		Kinesis: KinesisConfig{
			StreamName: "",
			Region:     "",
			RoleARN:    "",
		},
		Sentry: SentryConfig{
			Dsn:   "",
			Tags:  "{}",
			Debug: false,
		},
		LogLevel: "info",
		Target:   "stdout",
	}

	// Override values from environment
	return configFromEnv(defaultConfig)
}

// configFromEnv loads the config struct from environment variables
func configFromEnv(c *Config) *Config {
	return &Config{
		PubSub: PubSubConfig{
			ProjectID:         getEnvOrElse("PUBSUB_PROJECT_ID", c.PubSub.ProjectID),
			TopicName:         getEnvOrElse("PUBSUB_TOPIC_NAME", c.PubSub.TopicName),
			ServiceAccountB64: getEnvOrElse("PUBSUB_SERVICE_ACCOUNT_B64", c.PubSub.ServiceAccountB64),
		},
		Kinesis: KinesisConfig{
			StreamName: getEnvOrElse("KINESIS_STREAM_NAME", c.Kinesis.StreamName),
			Region:     getEnvOrElse("KINESIS_REGION", c.Kinesis.Region),
			RoleARN:    getEnvOrElse("KINESIS_ROLE_ARN", c.Kinesis.RoleARN),
		},
		Sentry: SentryConfig{
			Dsn:   getEnvOrElse("SENTRY_DSN", c.Sentry.Dsn),
			Tags:  getEnvOrElse("SENTRY_TAGS", c.Sentry.Tags),
			Debug: getEnvBoolOrElse("SENTRY_DEBUG", c.Sentry.Debug),
		},
		LogLevel: getEnvOrElse("LOG_LEVEL", c.LogLevel),
		Target:   getEnvOrElse("TARGET", c.Target),
	}
}

// GetTarget builds and returns the target that is configured
func (c *Config) GetTarget() (Target, error) {
	if c.Target == "stdout" {
		return NewStdoutTarget()
	} else if c.Target == "kinesis" {
		return NewKinesisTarget(c.Kinesis.Region, c.Kinesis.StreamName, c.Kinesis.RoleARN)
	} else if c.Target == "pubsub" {
		return NewPubSubTarget(c.PubSub.ProjectID, c.PubSub.TopicName, c.PubSub.ServiceAccountB64)
	} else {
		return nil, fmt.Errorf("Invalid target found; expected one of 'stdout, kinesis' and got '%s'", c.Target)
	}
}

// --- HELPERS

// getEnvOrElse returns an environment variable value or a default
func getEnvOrElse(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

// getEnvBoolOrElse returns an environment variable value and casts it to a boolean or passes a default
func getEnvBoolOrElse(key string, defaultVal bool) bool {
	if value, exists := os.LookupEnv(key); exists {
		mValue, err := strconv.ParseBool(value)
		if err != nil {
			log.Error(fmt.Sprintf("Error converting string to bool for key %s: %s; using default '%v'", key, err.Error(), defaultVal))
			return defaultVal
		}
		return mValue
	}
	return defaultVal
}
