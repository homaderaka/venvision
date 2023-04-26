package config

type AppConfig struct {
	LogLevel string `env:"LOG_LEVEL" env-default:"trace"`
}
