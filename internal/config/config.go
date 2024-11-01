package config

import (
	"time"

	"github.com/spf13/viper"
)

type DatabaseConfig struct {
	URL string `mapstructure:"url"`
}

type ServerConfig struct {
	Port string `mapstructure:"port"`
}

type JWTConfig struct {
	SecretKey     string        `mapstructure:"secret_key"`
	TokenLifetime time.Duration `mapstructure:"token_lifetime"`
}

type LoggingConfig struct {
	Level  string `mapstructure:"level"`
	Format string `mapstructure:"format"`
}

type EmailConfig struct {
	Password string `mapstructure:"password"`
	Address  string `mapstructure:"address"`
	SmtpHost string `mapstructure:"smtp_host"`
	SmtpPort string `mapstructure:"smtp_port"`
}

type Config struct {
	Database DatabaseConfig `mapstructure:"database"`
	Server   ServerConfig   `mapstructure:"server"`
	JWT      JWTConfig      `mapstructure:"jwt"`
	Logging  LoggingConfig  `mapstructure:"logging"`
	Email    EmailConfig    `mapstructure:"email"`
}

// LoadConfig - функция для загрузки конфигурации из файла.
func LoadConfig(path string) (*Config, error) {
	viper.SetConfigName("config") // Имя файла конфигурации (без расширения).
	viper.SetConfigType("yaml")   // Формат файла.
	viper.AddConfigPath(path)     // Путь к директории, где находится файл конфигурации.

	// Читаем конфигурационный файл.
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config

	// Декодируем конфигурацию в структуру Cfg.
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
