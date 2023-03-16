package main

import (
	"fmt"
	"os"

	"github.com/Str1kez/chatGPT-bot/internal/config"
	chatgpt "github.com/Str1kez/chatGPT-bot/pkg/chatGPT"
	"github.com/Str1kez/chatGPT-bot/pkg/telegram"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	chat := chatgpt.NewChatGPT(cfg.OpenAIToken)
	bot, err := telegram.NewBot(telegram.GenerateSettings(cfg.BotToken), chat)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	bot.InitCommands()
	bot.InitHandlers(cfg.Admins)
	bot.Start()
}
