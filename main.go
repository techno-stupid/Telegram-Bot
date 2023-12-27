package main

import (
	"TechnoStupidBot/app"
	"TechnoStupidBot/handler"
	"TechnoStupidBot/telegram"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/spf13/viper"
	"log"
)

func init() {
	app.Init()
	fmt.Println(viper.GetString("TELEGRAM_API_TOKEN"))
}

func main() {
	botHandler, err := telegram.NewBotHandler()
	if err != nil {
		panic(err)
	}

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30

	updates, err := botHandler.StartPolling(updateConfig)
	if err != nil {
		panic(err)
	}

	messageHandler, err := handler.NewMessageHandler(viper.GetString("YOUTUBE_API_KEY"))
	if err != nil {
		log.Fatal(err)
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		text, replyToMessageID := messageHandler.HandleMessage(update)

		if err := botHandler.SendMessage(update.Message.Chat.ID, text, replyToMessageID); err != nil {
			panic(err)
		}
	}
}
