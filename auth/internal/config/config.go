package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type ServiceConfig struct {
	ServerPort string `mapstructure:"server_port"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
}

func (dc DatabaseConfig) ToDSN() string {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dc.Host, dc.Port, dc.User, dc.Password, dc.Name,
	)
	return dsn
}

type AppConfig struct {
	Service  ServiceConfig  `mapstructure:"service"`
	Database DatabaseConfig `mapstructure:"database"`
}

func LoadConfig(path string) (AppConfig, error) {
	var config AppConfig

	viper.SetConfigFile(path)

	if err := viper.ReadInConfig(); err != nil {
		return config, fmt.Errorf("error reading config file: %w", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, fmt.Errorf("unable to unmarshal config: %w", err)
	}

	return config, nil
}
