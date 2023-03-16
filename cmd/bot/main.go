package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Str1kez/chatGPT-bot/internal/config"
	chatgpt "github.com/Str1kez/chatGPT-bot/pkg/chatGPT"
	"github.com/Str1kez/chatGPT-bot/pkg/ocr"
	"github.com/Str1kez/chatGPT-bot/pkg/telegram"
	"gopkg.in/telebot.v3"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	chat := chatgpt.NewChatGPT(cfg.OpenAIToken)
	recognitionClient := ocr.NewTesseractClient()

	botSettings := telebot.Settings{
		Token:  cfg.BotToken,
		Poller: &telebot.LongPoller{Timeout: time.Minute},
	}
	bot, err := telegram.NewBot(botSettings, chat, recognitionClient)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	bot.InitCommands()
	bot.InitHandlers(cfg.Admins)
	bot.Start()
}
