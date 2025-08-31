package config

import (
	"github.com/spf13/viper"
	"github.com/terminator791/clean-architecture-GO/internal/infrastructure/database"
)

// Config holds all configuration for the application
type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database database.Config `mapstructure:"database"`
}

// ServerConfig holds server configuration
type ServerConfig struct {
	Port int `mapstructure:"port"`
}

// Load loads configuration from environment variables and config files
func Load() (*Config, error) {
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")

	// Set defaults
	viper.SetDefault("server.port", 8080)
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", 5432)
	viper.SetDefault("database.user", "postgres")
	viper.SetDefault("database.password", "password")
	viper.SetDefault("database.dbname", "products")
	viper.SetDefault("database.sslmode", "disable")

	// Enable environment variable binding
	viper.AutomaticEnv()
	viper.SetEnvPrefix("APP")

	// Bind environment variables
	viper.BindEnv("server.port", "APP_SERVER_PORT")
	viper.BindEnv("database.host", "APP_DB_HOST")
	viper.BindEnv("database.port", "APP_DB_PORT")
	viper.BindEnv("database.user", "APP_DB_USER")
	viper.BindEnv("database.password", "APP_DB_PASSWORD")
	viper.BindEnv("database.dbname", "APP_DB_NAME")
	viper.BindEnv("database.sslmode", "APP_DB_SSLMODE")

	// Read config file if it exists
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, err
		}
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}