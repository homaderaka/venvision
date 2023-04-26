package config

type StorageConfig struct {
	Host     string `env:"HOST" env-default:"postgres"`
	Port     string `env:"PORT" env-default:"5432"`
	Database string `env:"DATABASE" env-default:"postgres"`
	Username string `env:"POSTGRES_USERNAME" env-default:"postgres"`
	Password string `env:"POSTGRES_PASSWORD" env-default:"postgres"`
}
