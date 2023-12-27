package handler

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

// MessageHandler handles incoming messages
type MessageHandler struct{}

// HandleMessage processes the incoming message
func (h *MessageHandler) HandleMessage(update tgbotapi.Update) (string, int) {
	// Add your custom logic for processing messages here
	return generateResponse(update.Message.Text), update.Message.MessageID
}

func generateResponse(text string) string {
	return text
}
