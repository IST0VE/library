package config

type Config struct {
	TelegramToken string `yaml:"telegram_token"`
	DBHost        string `yaml:"db_host"`
	DBUser        string `yaml:"db_user"`
	DBPassword    string `yaml:"db_password"`
	DBName        string `yaml:"db_name"`
}
