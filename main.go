package main

import (
	"log"
	"os"

	"github.com/bps-pasaman-barat/bot_tele_bps.git/handler"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	token := os.Getenv("BOT_TELEGRAM_TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60
	commands := tgbotapi.NewSetMyCommands(

		tgbotapi.BotCommand{
			Command:     "menu",
			Description: "Menampilkan menu utama",
		},

	)

	bot.Request(commands)
	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {

		if update.Message == nil {
			continue
		}

		msg := handler.HandleMessage(update)

		bot.Send(msg)
	}
}