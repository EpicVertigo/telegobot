package main

import (
	"log"
	"os"
	"runtime"

	tg "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"

	"ext"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Telega
	token := os.Getenv("TELEGRAM_TOKEN")
	bot, err := tg.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	// bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tg.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	ext.InitializeScheduler(bot)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		// Manual GC call (experimental, just for Raspberry 1)
		runtime.GC()

		if update.Message.Text == "tasks" {
			ext.SendIncompleteTasks(bot)
		} else {

			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tg.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}
}
