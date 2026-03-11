package main

import "github.com/bps-pasaman-barat/bot_tele_bps.git/question"

// import (
// 	"log"
// 	"os"
// 	"time"

// 	"github.com/bps-pasaman-barat/bot_tele_bps.git/handler"
// 	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
// 	"github.com/joho/godotenv"
// )

// var userTimers = map[int64]*time.Timer{}
// var sessionActive = map[int64]bool{}

func main() {

	question.Question()

	// godotenv.Load()

	// token := os.Getenv("BOT_TELEGRAM_TOKEN")

	// bot, err := tgbotapi.NewBotAPI(token)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// updateConfig := tgbotapi.NewUpdate(0)
	// updateConfig.Timeout = 60
	// commands := tgbotapi.NewSetMyCommands(

	// 	tgbotapi.BotCommand{
	// 		Command:     "menu",
	// 		Description: "Menampilkan menu utama",
	// 	},
	// )

	// bot.Request(commands)

	// updates := bot.GetUpdatesChan(updateConfig)

	// for update := range updates {

	// 	if update.Message == nil {
	// 		continue
	// 	}
	// 	chatID := update.Message.Chat.ID

	// 	if !sessionActive[chatID] {
	// 		if update.Message.Text != "/menu" {
	// 			msg := tgbotapi.NewMessage(chatID, "Silakan ketik /menu untuk memulai layanan.")
	// 			bot.Send(msg)
	// 			continue
	// 		}
	// 		sessionActive[chatID] = true
	// 	}

	// 	msg := handler.HandleMessage(update)
	// 	bot.Send(msg)

	// 	if t, ok := userTimers[chatID]; ok {
	// 		t.Stop()
	// 	}

	// 	userTimers[chatID] = time.AfterFunc(5*time.Minute, func() {

	// 		thanks := tgbotapi.NewMessage(chatID,
	// 			`Terima kasih telah menggunakan BOT LAYANAN BPS PASAMAN BARAT! 😊

	// 			Untuk mendapatkan informasi dan update terbaru dari kami, jangan lupa follow Instagram BPS Pasaman Barat:
	// 			https://www.instagram.com/bps_pasbar`)

	// 		bot.Send(thanks)

	// 		delete(sessionActive, chatID)
	// 		delete(userTimers, chatID)
	// 	})

	// }
}
