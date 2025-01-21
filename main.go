package main

import (
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {

	botToken := os.Getenv("TOKEN")
	if botToken == "" {
		log.Fatalf("BOT TOKEN not set")
	}

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatalf("Ошибка при создании бота: %v", err)
	}

	bot.Debug = true
	fmt.Printf("Бот %s успешно запущен\n", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			chatID := update.Message.Chat.ID

			if update.Message.Text == "/start" {
				msg := tgbotapi.NewMessage(chatID, "Какие данные о погоде вы хотите узнать?\n1. Описание, температура, влажность")
				tgbotapi.NewInlineKeyboardButtonData("1. Описание, температура, влажность", "1. Описание, температура, влажность")
				bot.Send(msg)
				continue
			}
			if update.Message.Text == "1" {
				msg := tgbotapi.NewMessage(chatID, "Введите город:")
				bot.Send(msg)
				continue
			}
			cityName := update.Message.Text
			weatherInfo := getWeather(cityName) // Запрос погоды для города
			msg := tgbotapi.NewMessage(chatID, weatherInfo)
			bot.Send(msg)
			continue
		}
	}
}
