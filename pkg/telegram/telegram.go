package telegram

import (
	"github.com/Str1kez/chatGPT-bot/internal/config"
	"gopkg.in/telebot.v3"
)

type Bot struct {
	bot               *telebot.Bot
	config            *config.BotConfig
	completionClient  completion
	recognitionClient recognition
}

type recognition interface {
	RecognitionFromBytes(photo []byte) (string, error)
}

type completion interface {
	PerformCompletion(text string) (string, error)
}

func NewBot(settings telebot.Settings, config *config.BotConfig, chat completion, recognitionClient recognition) (*Bot, error) {
	bot, err := telebot.NewBot(settings)
	if err != nil {
		return nil, err
	}
	return &Bot{bot: bot, config: config, completionClient: chat, recognitionClient: recognitionClient}, nil
}

func (b *Bot) Start() {
	b.bot.Start()
}
