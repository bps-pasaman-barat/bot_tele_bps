package handler

import (
	"github.com/bps-pasaman-barat/bot_tele_bps.git/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var userStates = make(map[int64]string)

type Action struct {
	NextState string
	ReplyFunc func() string
}

var menuFlow = map[string]map[string]Action{
	"main": {
		"1":  {"menu_1", service.MenuCariData},
		"2":  {"menu_2", service.BRSMenu},
		"3":  {"menu_3", service.PublikasiMenu},
		"4":  {"menu_4", service.InfoKegiatanMenu},
		"5":  {"menu_5", service.ProfilMenu},
		"6":  {"menu_6", service.MediaSosialMenu},
		"7":  {"menu_7", service.GaleriInfografisMenu},
		"8":  {"menu_8", service.RekomendasiKegiatanStatistikMenu},
		"9":  {"menu_9", service.MetadataStatistikMenu},
		"10": {"menu_10", service.PengaduanMenu},
		"11": {"menu_11", service.ChatAdminMenu},
		"99": {"main", service.MainMenu},
	},
	"menu_1": {
		"1":  {"menu_1_1", service.MenuStatistikDemografiSosial},
		"2":  {"menu_1_2", service.MenuStatistikEkonomi},
		"3":  {"menu_1_3", service.MenuStatistikEkonomi},
		"4":  {"menu_1_4", service.MenuStatistikEkonomi},
		"99": {"main", service.MainMenu},
	},
	"menu_1_1": {
		"1":  {"menu_1_1_item", service.KependudukanMigrasi},
		"2":  {"menu_1_1_item", service.TenagaKerja},
		"3":  {"menu_1_1_item", service.Pendidikan},
		"4":  {"menu_1_1_item", service.Kesehatan},
		"5":  {"menu_1_1_item", service.KonsumsiPendapatan},
		"6":  {"menu_1_1_item", service.PerlindunganSosial},
		"7":  {"menu_1_1_item", service.PemukimanDanPerumahan},
		"8":  {"menu_1_1_item", service.HukumDanKriminalt},
		"9":  {"menu_1_1_item", service.Budaya},
		"10": {"menu_1_1_item", service.AktivitasKomunitas},
		"11": {"menu_1_1_item", service.PenggunaanWaktu},
		"99": {"menu_1", service.MenuCariData},
	},
	"menu_1_1_item": {
		"99": {"menu_1_1", service.MenuStatistikDemografiSosial},
	},
	"menu_1_2": {
		"1":  {"menu_1_2_item", service.StatistikMakroEkonomi},
		"2":  {"menu_1_2_item", service.NeracaEkonomi},
		"3":  {"menu_1_2_item", service.StatistikBisnis},
		"4":  {"menu_1_2_item", service.StatsitikSektoral},
		"5":  {"menu_1_2_item", service.KeuanganPemerintah},
		"6":  {"menu_1_2_item", service.PerdaganganInternasionalNeraca},
		"7":  {"menu_1_2_item", service.HargaHarga},
		"8":  {"menu_1_2_item", service.BiayaTenagaKerja},
		"9":  {"menu_1_2_item", service.IlmuPenegtahuanTeknologiINovasi},
		"10": {"menu_1_2_item", service.PertanianKehutana},
		"11": {"menu_1_2_item", service.Energi},
		"12": {"menu_1_2_item", service.PertambanganManufaktur},
		"13": {"menu_1_2_item", service.Transportasi},
		"14": {"menu_1_2_item", service.Pariwisata},
		"15": {"menu_1_2_item", service.Perbankan},
		"99": {"menu_1", service.MenuCariData},
	},
	"menu_1_2_item": {
		"99": {"menu_1_2", service.MenuStatistikEkonomi},
	},
	"menu_2":  {"99": {"main", service.MainMenu}},
	"menu_3":  {"99": {"main", service.MainMenu}},
	"menu_4":  {"99": {"main", service.MainMenu}},
	"menu_5":  {"99": {"main", service.MainMenu}},
	"menu_6":  {"99": {"main", service.MainMenu}},
	"menu_7":  {"99": {"main", service.MainMenu}},
	"menu_8":  {"99": {"main", service.MainMenu}},
	"menu_9":  {"99": {"main", service.MainMenu}},
	"menu_10": {"99": {"main", service.MainMenu}},
	"menu_11": {"99": {"main", service.MainMenu}},
}

func getFallbackMessage(state string) string {
	if state == "main" {
		return service.MenuTidakTersedia()
	}
	return service.MenuPencarianTidakTersedia()
}

func HandleMessage(update tgbotapi.Update) tgbotapi.MessageConfig {

	chatID := update.Message.Chat.ID
	text := update.Message.Text
	msg := tgbotapi.NewMessage(chatID, "")

	state, exists := userStates[chatID]
	if !exists {
		state = "main"
	}

	if text == "/start" || text == "/menu" || text == "00" || text == "0" {
		userStates[chatID] = "main"
		msg.Text = service.MainMenu()
		return msg
	}

	if stateFlow, stateExists := menuFlow[state]; stateExists {
		if action, validText := stateFlow[text]; validText {
			userStates[chatID] = action.NextState
			msg.Text = action.ReplyFunc()
			return msg
		}
	}

	msg.Text = getFallbackMessage(state)
	return msg
}
