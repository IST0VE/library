package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	TelegramToken string `mapstructure:"telegram_token"`
	DBHost        string `mapstructure:"db_host"`
	DBUser        string `mapstructure:"db_user"`
	DBPassword    string `mapstructure:"db_password"`
	DBName        string `mapstructure:"db_name"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %s", err)
	}

	return config, err
}
