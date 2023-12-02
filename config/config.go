package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type (
	Config struct {
		Server   ServerConfig
		Database DatabaseConfig
	}

	ServerConfig struct {
		Env          string
		Port         int
		ReadTimeout  time.Duration
		WriteTimeout time.Duration
	}

	DatabaseConfig struct {
		Host     string
		Port     int
		DBName   string
		User     string
		Password string
		TimeZone string
	}
)

// LoadConfig 'config-[env].yaml' 파일로부터 설정 정보를 읽는다.
func LoadConfig(env string) (Config, error) {
	viper.SetConfigName(fmt.Sprintf("config/config-%s", env))
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
		return Config{}, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
		return Config{}, err
	}

	return config, nil
}
