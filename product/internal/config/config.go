package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

type Config struct {
	IsDebug       bool          `env:"IS_DEBUG" env-default:"false"`
	IsDevelopment bool          `env:"IS_DEV" env-default:"false"`
	ListenConfig  ListenConfig  `env:"LISTEN_CONFIG"`
	StorageConfig StorageConfig `env:"STORAGE_CONFIG"`
	AppConfig     AppConfig     `env:"APP_CONFIG"`
	//Kafka         struct {
	//	URL   string `env:"KAFKA_URL" env-default:"kafka:9092"`
	//	Topic string `env:"TOPIC" env-default:"proxy_checker"`
	//}
}

type Option func(*Config)

func WithDebug(isDebug bool) Option {
	return func(c *Config) {
		c.IsDebug = isDebug
	}
}

func WithDevelopment(isDevelopment bool) Option {
	return func(c *Config) {
		c.IsDevelopment = isDevelopment
	}
}

// Add similar functional options for other fields

func NewConfig(options ...Option) *Config {
	c := &Config{
		// Set default values
		IsDebug:       false,
		IsDevelopment: false,
		ListenConfig:  ListenConfig{ /*...*/ },
		StorageConfig: StorageConfig{ /*...*/ },
		AppConfig:     AppConfig{ /*...*/ },
	}

	for _, option := range options {
		option(c)
	}

	if err := cleanenv.ReadEnv(c); err != nil {
		helpText := "An error occurred during reading config"
		help, _ := cleanenv.GetDescription(c, &helpText)
		log.Println(help)
		log.Fatal(err)
	}

	return c
}
