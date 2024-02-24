package main

import (
	"log"

	"github.com/ist0ve/library/config"
	"github.com/ist0ve/library/internal/app"
	"github.com/ist0ve/library/internal/repository"
)

func main() {
	cfg := config.LoadConfig() // Загрузите конфигурацию
	repo, err := repository.NewRepository(cfg)
	if err != nil {
		log.Fatalf("failed to initialize repository: %v", err)
	}

	app.StartBot(cfg.TelegramToken, repo)
}
