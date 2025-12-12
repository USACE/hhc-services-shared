package config

import (
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/caarlos0/env/v11"
)

// Config holds application configuration variables
type Config struct {
	AwsConfig              aws.Config
	UsePathStyle           bool          `env:"USE_PATH_STYLE" envDefault:"false"`
	ApplicationKey         string        `env:"APPLICATION_KEY"`
	AuthEnvironment        string        `env:"AUTH_ENVIRONMENT"`
	AuthPublicKey          string        `env:"AUTH_PUBLIC_KEY"`
	Dbuser                 string        `env:"PGUSER,required"`
	Dbpass                 string        `env:"PGPASSWORD,required"`
	Dbname                 string        `env:"PGDATABASE,required"`
	Dbhost                 string        `env:"PGHOST,required"`
	Dbsslmode              string        `env:"PGSSLMODE" envDefault:"require"`
	PgxPoolMaxconns        int           `env:"PGX_POOL_MAXCONNS" envDefault:"10"`
	PgxPoolMinconns        int           `env:"PGX_POOL_MINCONNS" envDefault:"5"`
	PgxPoolMaxconnIdletime time.Duration `env:"PGX_POOL_MAXCONN_IDLETIME" envDefault:"30m"`
	S3Bucket               string        `env:"S3_BUCKET,required"`
	S3DefaultIndex         string        `env:"S3_DEFAULT_INDEX" envDefault:"index.html"`
	S3PrefixStatic         string        `env:"S3_PREFIX_STATIC" envDefault:"/"`
	ResourceAccessRoles    []string      `env:"RESOURCE_ACCESS_ROLES" envDefault:"public"`
	ContextKey             string        `env:"CONTEXT_KEY" envDefault:"user"`
	ApiLog                 bool          `env:"API_LOG" envDefault:"true"`
}

// ParseEnvVars parses environment variables and sets them to the Config struct
func (c *Config) ParseEnvVars() error {
	return env.ParseWithOptions(c, env.Options{
		OnSet: func(tag string, value any, isDefault bool) {
			if tag == "AUTH_ENVIRONMENT" {
				if strings.ToLower(value.(string)) == "dev" {
					c.AuthPublicKey = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArYxyX6mFWXEEpi8GhEs8GbUjZwYLIJ7ixEIoIZN1f4C7LoNMxz5mrDZcojNi91xSXqtFLlXfYTc/sI4JLYUEzKE0fNUxY9jldzI36ZLvIMqGg7KqaFukI3WO1AVejkJ77Lox+V20nJoZTrO577uElfIsqlJc11HHojME4f/Q7OOYoTPE4yYOGP8WbLPg4CSiSNR+ZYA4JdDLMZxD+FduhHkE7QbPZGsZqXCnr1UDzgNUaXFbufsmGo1N2h9eQOTNu6aV9zI7DdMZkVCbApwEov+p2n8EMp3xAZ5tAviXNzP8z3oifsw8XQLFFCyUUEr8e3kCmLW97lV7ys5iWnNhMQIDAQAB"
				} else if strings.ToLower(value.(string)) == "test" {
					c.AuthPublicKey = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArYxyX6mFWXEEpi8GhEs8GbUjZwYLIJ7ixEIoIZN1f4C7LoNMxz5mrDZcojNi91xSXqtFLlXfYTc/sI4JLYUEzKE0fNUxY9jldzI36ZLvIMqGg7KqaFukI3WO1AVejkJ77Lox+V20nJoZTrO577uElfIsqlJc11HHojME4f/Q7OOYoTPE4yYOGP8WbLPg4CSiSNR+ZYA4JdDLMZxD+FduhHkE7QbPZGsZqXCnr1UDzgNUaXFbufsmGo1N2h9eQOTNu6aV9zI7DdMZkVCbApwEov+p2n8EMp3xAZ5tAviXNzP8z3oifsw8XQLFFCyUUEr8e3kCmLW97lV7ys5iWnNhMQIDAQAB"
				} else if strings.ToLower(value.(string)) == "prod" {
					c.AuthPublicKey = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAgaLcKGp9KKeN+8REa4oHK41PQYpvIeP7XpXmPB70cV8uBBx8Er3SDrZ2TAz9UKZ2Z6m6QRreQjgk2FI+EQ2bHWToMRhnthIzbuHzI64GyBjCnGhu3sd0OFb9wTAvu6TcV7w+q7+WrVIF1vzHlpFo7qLewxJjEAKzJGx3EgDFhlRCPXG4BjP4Lsg/rBpV3ltZ74HtTlx3r7XeDKCIIgqAJOQueaQtwR7Snp2FFY3is/PHrWNKWLw3lRV0Lm4VtGHm4YOAqCwq6FfyHLjjohp2JXuzTVB+9s7cmbLq1dyDBCWkX02s4g3AZuJcycyrie+8TDvbCJ+ogHLcixwLDizXaQIDAQAB"
				} else {
					c.AuthPublicKey = "secret"
				}
			}
		},
	})
}
