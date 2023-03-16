package telegram

import (
	chatgpt "github.com/Str1kez/chatGPT-bot/pkg/chatGPT"
	"gopkg.in/telebot.v3"
)

type Bot struct {
	bot               *telebot.Bot
	openAI            *chatgpt.ChatGPT
	recognitionClient recognition
}

type recognition interface {
	RecognitionFromBytes(photo []byte) (string, error)
}

func NewBot(settings telebot.Settings, chat *chatgpt.ChatGPT, recognitionClient recognition) (*Bot, error) {
	bot, err := telebot.NewBot(settings)
	if err != nil {
		return nil, err
	}
	return &Bot{bot: bot, openAI: chat, recognitionClient: recognitionClient}, nil
}

func (b *Bot) Start() {
	b.bot.Start()
}
