package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/sashabaranov/go-openai"
	"gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/middleware"
)

const MAX_TOKENS = 1024

func MapStringToInt64(str []string) ([]int64, error) {
	// TODO: IN DIR TOOLS
	result := make([]int64, len(str))

	for i, s := range str {
		num, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return nil, err
		}
		result[i] = num
	}
	return result, nil
}

func main() {
	// TODO: implement architecture
	// ? add viper for env
	openaiClient := openai.NewClient(os.Getenv("OPENAI_TOKEN"))
	botConfig := telebot.Settings{
		Token:  os.Getenv("BOT_TOKEN"),
		Poller: &telebot.LongPoller{Timeout: time.Minute},
	}
	admins, err := MapStringToInt64(strings.Split(os.Getenv("ADMINS"), ","))
	if err != nil {
		fmt.Printf("Error occured in parsing Admin list: %v\n", err)
		return
	}

	bot, err := telebot.NewBot(botConfig)
	if err != nil {
		fmt.Printf("Error occured in Bot Init: %v\n", err)
		return
	}

	bot.Handle("/start", func(c telebot.Context) error {
		return c.Send("Здарова!\nПиши и получишь ответ")
	})

	bot.Handle("/code", func(c telebot.Context) error {
		fmt.Println(c.Sender().ID)
		return c.Send("Аххах стырил у тебя код 😎")
	})

	bot.Use(middleware.Whitelist(admins...))

	bot.Handle(telebot.OnText, func(c telebot.Context) error {
		// TODO: НАПИСАТЬ MIDDLEWARE
		if c.Chat().Type != telebot.ChatPrivate {
			return nil
		}
		go func() {
			ctx := context.Background()
			openaiRequest := openai.CompletionRequest{
				Model:            openai.GPT3TextDavinci003,
				Prompt:           c.Text(),
				MaxTokens:        MAX_TOKENS,
				Temperature:      0.5,
				TopP:             1.0,
				FrequencyPenalty: 0,
				PresencePenalty:  0,
			}
			resp, err := openaiClient.CreateCompletion(ctx, openaiRequest)
			if err != nil {
				fmt.Printf("Error occured in completion: %v\n", err)
				return
			}
			if err = c.Send(resp.Choices[0].Text); err != nil {
				fmt.Printf("Error occured in sending data to user: %v\n", err)
				return
			}
		}()

		return nil
	})

	bot.Start()
}
