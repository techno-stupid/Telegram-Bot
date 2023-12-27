package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/spf13/viper"
)

// BotHandler represents the Telegram bot handler
type BotHandler struct {
	Bot *tgbotapi.BotAPI
}

// NewBotHandler creates a new BotHandler instance
func NewBotHandler() (*BotHandler, error) {
	token := viper.GetString("TELEGRAM_API_TOKEN")
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	bot.Debug = true

	return &BotHandler{Bot: bot}, nil
}

// StartPolling starts polling for updates
func (b *BotHandler) StartPolling(updateConfig tgbotapi.UpdateConfig) (<-chan tgbotapi.Update, error) {
	return b.Bot.GetUpdatesChan(updateConfig), nil
}

// SendMessage sends a message to the specified chat
func (b *BotHandler) SendMessage(chatID int64, text string, replyToMessageID int) error {
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ReplyToMessageID = replyToMessageID

	_, err := b.Bot.Send(msg)
	return err
}
