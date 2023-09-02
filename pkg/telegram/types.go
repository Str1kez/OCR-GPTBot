package telegram

import (
	"github.com/Str1kez/OCR-GPTBot/internal/config"
	"github.com/sashabaranov/go-openai"
	"gopkg.in/telebot.v3"
)

type Bot struct {
	bot               *telebot.Bot
	config            *config.BotConfig
	completionClient  completion
	recognitionClient recognition
	settingsStorage   storage
}

type Settings struct {
	Context          string  `json:"context"`
	Stream           bool    `json:"stream"`
	Temperature      float32 `json:"temperature"`
	FrequencyPenalty float32 `json:"frequency_penalty"`
}

type recognition interface {
	GetTextFromImage(image []byte) (string, error)
}

type completion interface {
	GetCompletionText(text string, settings Settings) (string, error)
	GetCompletionStream(text string, settings Settings) (*openai.ChatCompletionStream, error)
}

type storage interface {
	Get(userId int64, key string) ([]byte, error)
	GetAll(userId int64) (Settings, error)
	Set(userId int64, key string, value interface{}) error
	SetAll(userId int64, settings Settings) error
	Del(userId int64, key string) error
	DelAll(userId int64) error
}
