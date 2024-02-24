package app

import (
	"log"
	"strings"

	"github.com/IST0VE/library/internal/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func StartBot(token string, repo *repository.Repository) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		chatID := update.Message.Chat.ID
		currentUserState := GetUserState(chatID)

		// Обработка не командных сообщений
		if currentUserState != "" && update.Message.Text != "" {
			switch currentUserState {
			case "awaitingBookDetails":
				// Разбор введённых данных и добавление книги в базу
				details := strings.Split(update.Message.Text, ";")
				if len(details) != 3 {
					msg := tgbotapi.NewMessage(chatID, "Пожалуйста, введите данные в формате: Название; Автор; Год.")
					bot.Send(msg)
					return
				}
				title, author, year := details[0], details[1], details[2]
				// Предположим, что у вас есть функция для добавления книги
				// AddBook(title, author, year)
				SetUserState(chatID, "") // Сброс состояния пользователя
				msg := tgbotapi.NewMessage(chatID, "Книга добавлена!")
				bot.Send(msg)

			case "awaitingSearchQuery":
				// Выполнение поиска по базе и отправка результатов
				query := update.Message.Text
				// Предположим, что у вас есть функция для поиска книг
				// SearchBooks(query)
				SetUserState(chatID, "") // Сброс состояния пользователя
				msg := tgbotapi.NewMessage(chatID, "Результаты поиска: ...")
				bot.Send(msg)
			}
			continue
		}

		// Обработка команд
		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case "добавить":
				SetUserState(chatID, "awaitingBookDetails")
				msg := tgbotapi.NewMessage(chatID, "Введите данные книги в формате: Название; Автор; Год")
				bot.Send(msg)

			case "поиск":
				SetUserState(chatID, "awaitingSearchQuery")
				msg := tgbotapi.NewMessage(chatID, "Введите запрос для поиска (название или автор)")
				bot.Send(msg)
			}
		}
	}

	// Обработка команд здесь
}

var userStates map[int64]string = make(map[int64]string)

// SetUserState устанавливает состояние пользователя
func SetUserState(chatID int64, state string) {
	userStates[chatID] = state
}

// GetUserState возвращает текущее состояние пользователя
func GetUserState(chatID int64) string {
	return userStates[chatID]
}

// Допустим, у вас есть функция для изменения состояния пользователя
// SetUserState(chatID int64, state string)
// GetUserState(chatID int64) string

// Здесь логика для обработки ответов после команды "добавить" и "поиск"
// Например, проверка состояния пользователя и соответствующая обработка
