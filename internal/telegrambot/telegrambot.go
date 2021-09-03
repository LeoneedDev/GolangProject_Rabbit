package telegrambot

import (
	"log"
	"reflect"
	"nigga/configs"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

)


func TelegramApi() {
	conf := configs.New()
	bot, err := tgbotapi.NewBotAPI(conf.Bot.Token)
	if err != nil {
		log.Printf("Ошибка у бота: %s", err)
	}

	update := tgbotapi.NewUpdate(0)
	update.Timeout = 60
	updates, err := bot.GetUpdatesChan(update)
	if err != nil {
		log.Printf("Ошибка у бота: %s", err)
	}

	for update_item := range updates {
		if update_item.Message == nil {
			continue
		} else {
			if reflect.TypeOf(update_item.Message.Text).Kind() == reflect.String && update_item.Message.Text != "" {
				switch update_item.Message.Text {
				case "/start":
					msg := tgbotapi.NewMessage(update_item.Message.Chat.ID, "Спроси у меня где кролик\n/where_rabbit")
					bot.Send(msg)
				case "/where_rabbi":
					msg := tgbotapi.NewMessage(update_item.Message.Chat.ID, "В поле гуляет \nНравиться?\n/like")
					bot.Send(msg)
				case "/like":
					msg := tgbotapi.NewMessage(update_item.Message.Chat.ID, "Отлично, буду знать что тебе понравилось")
					bot.Send(msg)
					log.Printf("Номер чата и имя пользователя, которому понравилось: №%d %s", update_item.Message.Chat.ID, update_item.Message.Chat.UserName)
				default:
					msg := tgbotapi.NewMessage(update_item.Message.Chat.ID, "Неверный запрос")
					bot.Send(msg)
				}

			} else {
				msg := tgbotapi.NewMessage(update_item.Message.Chat.ID, "Напиши /start и узнаешь где негр")
				bot.Send(msg)
			}

		}

	}

}
