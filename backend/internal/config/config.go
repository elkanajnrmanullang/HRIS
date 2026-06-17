package config

import (
	"log"
	"path/filepath"

	"github.com/spf13/viper"
)

// Config
type Config struct {
	App      AppConfig
	Postgres PostgresConfig
}

// AppConfig
type AppConfig struct {
	Env  string
	Port int
}

// PostgresConfig
type PostgresConfig struct {
	User     string
	Password string
	DB       string
	Host     string
	Port     int
}

// LoadConfig
func LoadConfig(path string) (*Config, error) {
	viper.SetConfigFile(filepath.Join(path, ".env"))
	viper.AutomaticEnv()

	viper.SetDefault("APP_ENV", "development")
	viper.SetDefault("APP_PORT", 8080)
	viper.SetDefault("POSTGRES_HOST", "localhost")
	viper.SetDefault("POSTGRES_PORT", 5432)

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Warning: failed to read .env file, relying on OS environment variables: %v\n", err)
	}

	config := &Config{
		App: AppConfig{
			Env:  viper.GetString("APP_ENV"),
			Port: viper.GetInt("APP_PORT"),
		},
		Postgres: PostgresConfig{
			User:     viper.GetString("POSTGRES_USER"),
			Password: viper.GetString("POSTGRES_PASSWORD"),
			DB:       viper.GetString("POSTGRES_DB"),
			Host:     viper.GetString("POSTGRES_HOST"),
			Port:     viper.GetInt("POSTGRES_PORT"),
		},
	}

	return config, nil
}