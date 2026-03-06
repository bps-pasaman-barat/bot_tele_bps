package handler

import (
	"github.com/bps-pasaman-barat/bot_tele_bps.git/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var userStates = make(map[int64]string)

func HandleMessage(update tgbotapi.Update) tgbotapi.MessageConfig {

	chatID := update.Message.Chat.ID
	text := update.Message.Text
	msg := tgbotapi.NewMessage(chatID, "")

	state, exists := userStates[chatID]
	if !exists {
		state = "main"
	}

	if text == "/start" || text == "/menu" || text == "00" || text=="0" {
		userStates[chatID] = "main"
		msg.Text = service.MainMenu()
		return msg
	}

	if text == "99" || text=="9" {
		userStates[chatID] = "main"
		msg.Text = service.MainMenu()
		return msg
	}
	if text == "199" || text=="19" {
		userStates[chatID] = "menu_1"
		msg.Text = service.MenuCariData()
		return msg
	}

	switch state {
	case "main":
		switch text {
		case "1":
			userStates[chatID] = "menu_1"
			msg.Text = service.MenuCariData()
		case "2":
			userStates[chatID] = "menu_2"
			msg.Text = service.BRSMenu()
		case "3":
			userStates[chatID] = "menu_3"
			msg.Text = service.PublikasiMenu()
		case "4":
			userStates[chatID] = "menu_4"
			msg.Text = service.InfoKegiatanMenu()
		case "5":
			userStates[chatID] = "menu_5"
			msg.Text = service.ProfilMenu()
		case "6":
			userStates[chatID] = "menu_6"
			msg.Text = service.MediaSosialMenu()
		case "7":
			userStates[chatID] = "menu_7"
			msg.Text = service.GaleriInfografisMenu()
		case "8":
			userStates[chatID] = "menu_8"
			msg.Text = service.RekomendasiKegiatanStatistikMenu()
		case "9":
			userStates[chatID] = "menu_9"
			msg.Text = service.MetadataStatistikMenu()
		case "10":
			userStates[chatID] = "menu_10"
			msg.Text = service.PengaduanMenu()
		case "11":
			userStates[chatID] = "menu_11"
			msg.Text = service.ChatAdminMenu()
		// case "12":
		// 	userStates[chatID] = "menu_12"
		// 	msg.Text = service.SaranMasukanMenu()
		default:
			msg.Text = service.MenuTidakTersedia()
		}

	case "menu_1":
		switch text {
		case "1":
			userStates[chatID] = "menu_1_1" // Submenu level 2: Statistik Demografi & Sosial
			msg.Text = service.MenuStatistikDemografiSosial()
		case "2" :
			userStates[chatID] = "menu_1_2"
			msg.Text = service.MenuStatistikEkonomi()
		case "3" :
			userStates[chatID] = "menu_1_3"
			msg.Text = service.MenuStatistikEkonomi()
		case "4" :
			userStates[chatID] = "menu_1_4"
			msg.Text = service.MenuStatistikEkonomi()
		
		default:
			msg.Text = service.MenuPencarianTidakTersedia()
		}

	default:
		msg.Text = service.MenuTidakTersedia()
	}

	return msg
}