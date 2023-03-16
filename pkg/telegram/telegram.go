package telegram

import (
	"time"

	chatgpt "github.com/Str1kez/chatGPT-bot/pkg/chatGPT"
	"gopkg.in/telebot.v3"
)

type Bot struct {
	bot    *telebot.Bot
	openAI *chatgpt.ChatGPT
}

func NewBot(settings telebot.Settings, chat *chatgpt.ChatGPT) (*Bot, error) {
	bot, err := telebot.NewBot(settings)
	if err != nil {
		return nil, err
	}
	return &Bot{bot: bot, openAI: chat}, nil
}

func (b *Bot) Start() {
	b.bot.Start()
}

func GenerateSettings(token string) telebot.Settings {
	return telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: time.Minute},
	}
}
