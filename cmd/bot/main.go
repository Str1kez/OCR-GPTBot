package main

import (
	"os"
	"path"
	"runtime"
	"strconv"
	"time"

	"github.com/Str1kez/chatGPT-bot/internal/config"
	chatgpt "github.com/Str1kez/chatGPT-bot/pkg/chatGPT"
	"github.com/Str1kez/chatGPT-bot/pkg/ocr"
	"github.com/Str1kez/chatGPT-bot/pkg/storage"
	"github.com/Str1kez/chatGPT-bot/pkg/telegram"
	log "github.com/sirupsen/logrus"
	"gopkg.in/telebot.v3"
)

func callerPrettyfier(frame *runtime.Frame) (function string, file string) {
	fileName := path.Base(frame.File) + ":" + strconv.Itoa(frame.Line)
	return "", fileName
}

func initLogger() {
	mode := os.Getenv("MODE")
	if mode == "prod" {
		log.SetFormatter(&log.JSONFormatter{CallerPrettyfier: callerPrettyfier})
		log.SetLevel(log.InfoLevel)
	} else {
		log.SetFormatter(&log.TextFormatter{FullTimestamp: true, CallerPrettyfier: callerPrettyfier})
		log.SetLevel(log.DebugLevel)
	}
	log.SetReportCaller(true)
	log.SetOutput(os.Stdout)
}

func main() {
	initLogger()

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Couldn't initialize config", err)
	}
	log.Infoln("Config has been parsed")

	chatCompletionClient := chatgpt.NewChatCompletionClient(cfg.OpenAIToken)
	recognitionClient := ocr.NewYandexOCRClient(&cfg.OCR)
	storageClient, err := storage.NewRedisStorage(&cfg.Storage)

	if err != nil {
		log.Fatalf("Couldn't connect to storage: %v\n", err)
	}

	botSettings := telebot.Settings{
		Token:     cfg.Bot.Token,
		Poller:    &telebot.LongPoller{Timeout: time.Minute},
		// ParseMode: telebot.ModeMarkdown, // https://core.telegram.org/bots/api#markdown-style
	}
	bot, err := telegram.NewBot(botSettings, &cfg.Bot, chatCompletionClient, recognitionClient, storageClient)
	if err != nil {
		log.Fatalf("Couldn't start bot: %v\n", err)
	}

	if err = bot.OnStartup(); err != nil {
		log.Errorf("Couldn't handle startup routines: %v\n", err)
	}

	bot.InitCommands()
	bot.InitHandlers()
	log.Infoln("Bot is working")
	bot.Start()
}
