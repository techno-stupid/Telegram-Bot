// File: handler/message_handler.go
package handler

import (
	"context"
	"fmt"
	"google.golang.org/api/youtube/v3"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"google.golang.org/api/option"
)

// MessageHandler handles incoming messages
type MessageHandler struct {
	YouTubeService *youtube.Service
}

// NewMessageHandler creates a new MessageHandler instance
func NewMessageHandler(apiKey string) (*MessageHandler, error) {
	ctx := context.Background()
	youtubeService, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
	}

	return &MessageHandler{
		YouTubeService: youtubeService,
	}, nil
}

// HandleMessage processes the incoming message
func (h *MessageHandler) HandleMessage(update tgbotapi.Update) (string, int) {
	// Your custom logic goes here
	receivedMessage := update.Message.Text

	// Check if the message is a command to search YouTube
	if receivedMessage == "/yt" {
		return "Please enter a search query after the command, e.g., /yt cats", update.Message.MessageID
	}

	// Check if the message starts with the YouTube command
	commandPrefixLength := len("/yt ")
	if len(receivedMessage) > commandPrefixLength && receivedMessage[:commandPrefixLength] == "/yt " {
		searchQuery := receivedMessage[commandPrefixLength:]
		videoLink, err := h.searchYouTube(searchQuery)
		if err != nil {
			log.Println("Error searching YouTube:", err)
			return "Error searching YouTube.", update.Message.MessageID
		}

		return videoLink, update.Message.MessageID
	}

	// Default response for non-matching messages
	return "I'm sorry, I didn't understand that. Type /yt followed by your search query to find a YouTube video.", update.Message.MessageID
}

// searchYouTube searches YouTube for videos and returns the link to the first video
func (h *MessageHandler) searchYouTube(query string) (string, error) {
	searchResponse, err := h.YouTubeService.Search.List([]string{"id", "snippet"}).Q(query).MaxResults(1).Type("video").Do()
	if err != nil {
		return "", err
	}

	if searchResponse == nil || len(searchResponse.Items) == 0 {
		return "No videos found for the given query.", nil
	}

	videoID := searchResponse.Items[0].Id.VideoId
	if videoID == "" {
		return "No videos found for the given query.", nil
	}

	videoLink := fmt.Sprintf("https://www.youtube.com/watch?v=%s", videoID)

	return videoLink, nil
}
