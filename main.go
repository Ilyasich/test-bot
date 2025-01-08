package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

var bot *tgbotapi.BotAPI

func main() {
	http.HandleFunc("/", TelegramWebhookHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Стандартный порт для локального запуска.
	}

	log.Printf("Starting server on port %s...", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func init() {
	// Загружаем переменные окружения
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Инициализируем Telegram-бота
	botToken := os.Getenv("TG_TOKEN")
	if botToken == "" {
		log.Fatal("TG_TOKEN is missing in environment variables")
	}

	bot, err = tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatal(err)
	}
}

// Обработчик HTTP-запросов для Telegram-бота
func TelegramWebhookHandler(w http.ResponseWriter, r *http.Request) {
	var update tgbotapi.Update
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	if update.Message != nil && update.Message.Text == "/start" {
		chatID := update.Message.Chat.ID
		if err := SendStartMessage(bot, chatID); err != nil {
			log.Printf("Error sending start message: %v", err)
		}
	}

	fmt.Fprint(w, "OK")
}

func SendStartMessage(bot *tgbotapi.BotAPI, chatID int64) error {
	img, err := os.Open("./image.png")
	if err != nil {
		return err
	}
	defer img.Close()

	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		[]tgbotapi.InlineKeyboardButton{
			tgbotapi.NewInlineKeyboardButtonURL("✨ Join the official channel", "https://t.me/geo_dao"),
		},
	)

	msg := tgbotapi.NewPhoto(chatID, tgbotapi.FileReader{Name: "image.png", Reader: img})
	msg.Caption = "🌍 Global Ecosystem of Opportunities (GEO) - AI & WEB3 educational-gaming platform for your healthy life: connect, join with friends and family, explore, learn, play, improve, monitor and save your Health every day, complete quests & farm coins."
	//msg := tgbotapi.NewMessage(chatID, "Welcome to the bot!")
	msg.ParseMode = tgbotapi.ModeHTML
	msg.ReplyMarkup = keyboard

	if _, err := bot.Send(msg); err != nil {
		return err
	}
	return nil
}

// package main

// import (
// 	"log"
// 	"os"

// 	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
// 	"github.com/joho/godotenv"
// )

// func main() {
// 	// Загружаем переменные окружения из .env
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatalf("Error loading .env file: %v", err)
// 	}

// 	// Получаем токен из переменной окружения
// 	botToken := os.Getenv("TG_TOKEN")
// 	if botToken == "" {
// 		log.Fatal("TELEGRAM_BOT_TOKEN is missing in environment variables")
// 	}

// 	// Настройка Telegram-бота
// 	bot, err := tgbotapi.NewBotAPI(botToken)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	bot.Debug = true
// 	log.Printf("Authorized on account %s", bot.Self.UserName)

// 	// Создаем канал для получения обновлений
// 	u := tgbotapi.NewUpdate(0)
// 	u.Timeout = 30
// 	updates := bot.GetUpdatesChan(u)
// 	defer bot.StopReceivingUpdates()

// 	// Основной цикл обработки сообщений
// 	for update := range updates {
// 		if update.Message != nil {
// 			chatID := update.Message.Chat.ID
// 			if update.Message.Text == "/start" {
// 				if err := SendStartMessage(bot, chatID); err != nil {
// 					log.Printf("Error sending start message: %v", err)
// 				}
// 			}
// 		}
// 	}
// }

// // Функция отправки стартового сообщения
// func SendStartMessage(bot *tgbotapi.BotAPI, chatID int64) error {
// 	img, err := os.Open("./image.png")
// 	if err != nil {
// 		return err
// 	}
// 	defer img.Close()

// 	keyboard := tgbotapi.NewInlineKeyboardMarkup(
// 		[]tgbotapi.InlineKeyboardButton{
// 			tgbotapi.NewInlineKeyboardButtonURL("✨ Join the official channel", "https://t.me/geo_dao"),
// 		},
// 	)

// 	msg := tgbotapi.NewPhoto(chatID, tgbotapi.FileReader{Name: "image.png", Reader: img})
// 	msg.Caption = "🌍 Global Ecosystem of Opportunities (GEO) - AI & WEB3 educational-gaming platform for your healthy life: connect, join with friends and family, explore, learn, play, improve, monitor and save your Health every day, complete quests & farm coins."
// 	msg.ParseMode = tgbotapi.ModeHTML
// 	msg.ReplyMarkup = keyboard

// 	if _, err := bot.Send(msg); err != nil {
// 		return err
// 	}
// 	return nil
// }
