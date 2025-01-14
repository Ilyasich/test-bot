package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	// Загружаем переменные окружения из .env
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Получаем токен из переменной окружения
	botToken := os.Getenv("TG_TOKEN")
	if botToken == "" {
		log.Fatal("TG_TOKEN is missing in environment variables")
	}

	// Инициализация Telegram-бота
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatalf("Failed to initialize bot: %v", err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	// Настройка получения обновлений
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30
	updates := bot.GetUpdatesChan(u)

	// Основной цикл обработки сообщений
	for update := range updates {
		// Проверяем, что это сообщение от пользователя
		if update.Message != nil {
			log.Printf("Received message from %s: %s", update.Message.From.UserName, update.Message.Text)

			chatID := update.Message.Chat.ID
			if update.Message.Text == "/start" {
				if err := SendStartMessage(bot, chatID); err != nil {
					log.Printf("Error sending start message: %v", err)
				} else {
					log.Printf("Start message sent to chat ID: %d", chatID)
				}
			}
		}
	}
}

// Функция отправки стартового сообщения
func SendStartMessage(bot *tgbotapi.BotAPI, chatID int64) error {
	img, err := os.Open("./image.png")
	if err != nil {
		log.Printf("Error opening image: %v", err)
		return err
	}
	defer img.Close()

	// Создаем клавиатуру с кнопкой
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		[]tgbotapi.InlineKeyboardButton{
			tgbotapi.NewInlineKeyboardButtonURL("✨ Join the official channel", "https://t.me/geo_dao"),
		},
	)

	// Создаем сообщение с изображением
	msg := tgbotapi.NewPhoto(chatID, tgbotapi.FileReader{Name: "image.png", Reader: img})
	msg.Caption = "🌍 Global Ecosystem of Opportunities (GEO) - AI & WEB3 educational-gaming platform for your healthy life: connect, join with friends and family, explore, learn, play, improve, monitor and save your Health every day, complete quests & farm coins."
	msg.ParseMode = tgbotapi.ModeHTML
	msg.ReplyMarkup = keyboard

	// Отправляем сообщение
	if _, err := bot.Send(msg); err != nil {
		log.Printf("Error sending photo message: %v", err)
		return err
	}

	return nil
}
