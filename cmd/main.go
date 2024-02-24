package main

import (
	"log"

	"github.com/yourusername/my-telegram-bot/config"
	"github.com/yourusername/my-telegram-bot/internal/app"
	"github.com/yourusername/my-telegram-bot/internal/repository"
)

func main() {
	cfg := config.LoadConfig() // Загрузите конфигурацию
	repo, err := repository.NewRepository(cfg)
	if err != nil {
		log.Fatalf("failed to initialize repository: %v", err)
	}

	app.StartBot(cfg.TelegramToken, repo)
}
