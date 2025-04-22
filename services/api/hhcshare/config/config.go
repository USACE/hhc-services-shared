package config

import (
	"time"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	AuthPublicKey         string
	ApplicationKey        string        `env:"APPLICATION_KEY"`
	ApplicationLog        bool          `env:"APPLICATION_LOG" envDefault:"false"`
	AuthEnvironment       string        `env:"AUTH_ENVIRONMENT,required"`
	AwsAccessKeyId        string        `env:"AWS_ACCESS_KEY_ID"`
	AwsSecretAccessKey    string        `env:"AWS_SECRET_ACCESS_KEY"`
	AwsDefaultRegion      string        `env:"AWS_DEFAULT_REGION"`
	AwsEndpointUrlS3      string        `env:"AWS_ENDPOINT_URL_S3"`
	AwsEndpointUrlSqs     string        `env:"AWS_ENDPOINT_URL_SQS"`
	DbUser                string        `env:"LDM_DB_USER,required"`
	DbPass                string        `env:"LDM_DB_PASS,required"`
	DbName                string        `env:"LDM_DB_NAME,required"`
	DbHost                string        `env:"LDM_DB_HOST,required"`
	DbSslMode             string        `env:"LDM_DB_SSL_MODE" envDefault:"require"`
	DbPoolMaxConns        int           `env:"LDM_DB_POOL_MAX_CONNS" envDefault:"10"`
	DbPoolMinConns        int           `env:"LDM_DB_POOL_MIN_CONNS" envDefault:"5"`
	DbPoolMaxConnIdleTime time.Duration `env:"LDM_DB_POOL_MAX_CONN_IDLE_TIME" envDefault:"30m"`
}

// ParseEnvVars parses environment variables and sets them to the Config struct
func (c *Config) ParseEnvVars() error {
	return env.ParseWithOptions(c, env.Options{
		OnSet: func(tag string, value interface{}, isDefault bool) {
			if tag == "AUTH_ENVIRONMENT" {
				if value == "DEV" {
					c.AuthPublicKey = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArYxyX6mFWXEEpi8GhEs8GbUjZwYLIJ7ixEIoIZN1f4C7LoNMxz5mrDZcojNi91xSXqtFLlXfYTc/sI4JLYUEzKE0fNUxY9jldzI36ZLvIMqGg7KqaFukI3WO1AVejkJ77Lox+V20nJoZTrO577uElfIsqlJc11HHojME4f/Q7OOYoTPE4yYOGP8WbLPg4CSiSNR+ZYA4JdDLMZxD+FduhHkE7QbPZGsZqXCnr1UDzgNUaXFbufsmGo1N2h9eQOTNu6aV9zI7DdMZkVCbApwEov+p2n8EMp3xAZ5tAviXNzP8z3oifsw8XQLFFCyUUEr8e3kCmLW97lV7ys5iWnNhMQIDAQAB"
				} else if value == "TEST" {
					c.AuthPublicKey = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArYxyX6mFWXEEpi8GhEs8GbUjZwYLIJ7ixEIoIZN1f4C7LoNMxz5mrDZcojNi91xSXqtFLlXfYTc/sI4JLYUEzKE0fNUxY9jldzI36ZLvIMqGg7KqaFukI3WO1AVejkJ77Lox+V20nJoZTrO577uElfIsqlJc11HHojME4f/Q7OOYoTPE4yYOGP8WbLPg4CSiSNR+ZYA4JdDLMZxD+FduhHkE7QbPZGsZqXCnr1UDzgNUaXFbufsmGo1N2h9eQOTNu6aV9zI7DdMZkVCbApwEov+p2n8EMp3xAZ5tAviXNzP8z3oifsw8XQLFFCyUUEr8e3kCmLW97lV7ys5iWnNhMQIDAQAB"
				} else if value == "PROD" {
					c.AuthPublicKey = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAgaLcKGp9KKeN+8REa4oHK41PQYpvIeP7XpXmPB70cV8uBBx8Er3SDrZ2TAz9UKZ2Z6m6QRreQjgk2FI+EQ2bHWToMRhnthIzbuHzI64GyBjCnGhu3sd0OFb9wTAvu6TcV7w+q7+WrVIF1vzHlpFo7qLewxJjEAKzJGx3EgDFhlRCPXG4BjP4Lsg/rBpV3ltZ74HtTlx3r7XeDKCIIgqAJOQueaQtwR7Snp2FFY3is/PHrWNKWLw3lRV0Lm4VtGHm4YOAqCwq6FfyHLjjohp2JXuzTVB+9s7cmbLq1dyDBCWkX02s4g3AZuJcycyrie+8TDvbCJ+ogHLcixwLDizXaQIDAQAB"
				} else {
					c.AuthPublicKey = "secret"
				}
			}
		},
	})
}
