package config

import (
	"fmt"
	"log"
	"os"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	PostgresConfig
	RedisConfig
	HTTPServerConfig
	DockertestConfig
	AppConfig
}

type AppConfig struct {
	GraceTime int `envconfig:"GRACE_TIME" default:"5"`
}

type PostgresConfig struct {
	PgUser     string `envconfig:"DATABASE_USER" default:"postgres"`
	PgPassword string `envconfig:"DATABASE_PASSWORD" default:"postgres"`
	PgHost     string `envconfig:"DATABASE_HOST" default:"localhost"`
	PgPort     string `envconfig:"DATABASE_PORT" default:"5432"`
	PgDatabase string `envconfig:"DATABASE_NAME" default:"postgres"`
	PgURL      string `envconfig:"DATABASE_URL"`
	SSLMode    string `envconfig:"SSL_MODE" default:"disable"`
}

type HTTPServerConfig struct {
	HTTPHost     string `envconfig:"HOST" default:"localhost"`
	HTTPPort     string `envconfig:"PORT" default:"3000"`
	WriteTimeout int    `envconfig:"WRITE_TIMEOUT" default:"10"`
	ReadTimeout  int    `envconfig:"READ_TIMEOUT" default:"10"`
}

type DockertestConfig struct {
	DockertestTimeout int `envconfig:"DOCKERTEST_TIMEOUT" default:"30"`
}

type RedisConfig struct {
	RedisHost string `envconfig:"REDIS_HOST" default:"localhost"`
	RedisPort string `envconfig:"REDIS_PORT" default:"6379"`
	RedisURL  string `envconfig:"REDIS_URL"`
	RateTTL   int    `envconfig:"RATE_TTL"`
}

func LoadConfig() (Config, error) {
	var config Config
	noPrefix := ""
	err := envconfig.Process(noPrefix, &config)
	if err != nil {
		return Config{}, fmt.Errorf("failed loading config: %w", err)
	}
	log.Printf("HTTP server address: %s%s\n", config.GetHTTPHost(), config.GetHTTPPort())

	return config, nil
}

func (cfg PostgresConfig) GetPgURL() string {
	if cfg.PgURL != "" {
		return cfg.PgURL
	}

	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.PgUser,
		cfg.PgPassword,
		cfg.PgHost,
		cfg.PgPort,
		cfg.PgDatabase,
		cfg.SSLMode,
	)
}

func (cfg HTTPServerConfig) GetHTTPPort() string {
	if envPort := os.Getenv("PORT"); envPort != "" {
		return fmt.Sprintf(":%s", envPort)
	}

	return cfg.HTTPPort
}

func (cfg HTTPServerConfig) GetHTTPHost() string {
	if envHost := os.Getenv("HOST"); envHost != "" {
		return envHost
	}

	return cfg.HTTPHost
}

func (cfg RedisConfig) GetRedisURL() string {
	if cfg.RedisURL != "" {
		return cfg.RedisURL
	}

	return fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort)
}
