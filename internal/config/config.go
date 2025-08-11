package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config holds the application configuration
type Config struct {
	DatabaseURL string `mapstructure:"DATABASE_URL"`
}

// LoadConfig loads the configuration from environment variables or a .env file
func LoadConfig() (*Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	// Read the configuration file
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	// Set default values
	viper.SetDefault("DATABASE_URL", "postgres://user:password@localhost:5432/mydb")

	// Unmarshal the configuration into the Config struct
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("unable to decode into struct: %w", err)
	}

	return &config, nil
}
