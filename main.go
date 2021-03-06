package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
)

var HELP = "/tomorrow - Прогноз погоды на завтра\n" +
	"/today - Прогноз погоды на сегодня\n" +
	"/week - Прогноз погоды на неделю\n" +
	"/change_city - Смена города"

var NOT_UNDERSTAND = "Не понел... Введите /help"

func main() {

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TG_TOKEN"))
	if err != nil {
		panic(err)
	}
	//bot.Debug = true

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		location := update.Message.Location
		if location != nil {
			msg.Text, _ = get_weather_by_coords(location.Longitude, location.Latitude)
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		} else {
			switch update.Message.Command() {
			case "start":
				msg.Text = "Привет. Я пока не готов, но надеюсь скоро меня сделают, а пока посмотри меню команд /help"
				btn := tgbotapi.KeyboardButton{
					RequestLocation: true,
					Text:            "Поделитесь вашим местоположением",
				}
				msg.ReplyMarkup = tgbotapi.NewReplyKeyboard([]tgbotapi.KeyboardButton{btn})
			case "today":
				msg.Text = "Погода сегодня"
			case "tomorrow":
				msg.Text = "Прогноз на завтра"
			case "week":
				msg.Text = "Прогноз на неделю"
			case "change_city":
				msg.Text = "Выберете город"
			case "help":
				msg.Text = HELP
			default:
				msg.Text = NOT_UNDERSTAND
			}
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
