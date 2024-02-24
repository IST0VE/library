package main

import (
	"log"

	"github.com/IST0VE/library/config"
	"github.com/IST0VE/library/internal/app"
	"github.com/IST0VE/library/internal/repository"
)

func main() {
	cfg := config.LoadConfig() // Загрузите конфигурацию
	repo, err := repository.NewRepository(cfg)
	if err != nil {
		log.Fatalf("failed to initialize repository: %v", err)
	}

	app.StartBot(cfg.TelegramToken, repo)
}
