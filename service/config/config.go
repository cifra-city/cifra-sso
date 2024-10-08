package config

import (
	"time"

	"github.com/spf13/viper"
)

// Config - структура для хранения всех настроек.
type Config struct {
	DatabaseURL   string        `mapstructure:"url"`
	ServerPort    string        `mapstructure:"port"`
	JWTSecretKey  string        `mapstructure:"secret_key"`
	TokenLifetime time.Duration `mapstructure:"token_lifetime"`
	LoggingLevel  string        `mapstructure:"level"`
	LoggingFormat string        `mapstructure:"format"`
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

	// Декодируем конфигурацию в структуру Config.
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
