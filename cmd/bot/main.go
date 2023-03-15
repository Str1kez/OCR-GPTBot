package main

import (
	"fmt"
	"os"

	"github.com/Str1kez/chatGPT-bot/internal/config"
)

func main() {
	// TODO: implement architecture

	cfg, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v\n", cfg)
	os.Exit(0)

	// botConfig := telebot.Settings{
	// 	Token:  os.Getenv("BOT_TOKEN"),
	// 	Poller: &telebot.LongPoller{Timeout: time.Minute},
	// }
	// admins, err := MapStringToInt64(strings.Split(os.Getenv("ADMINS"), ","))
	// if err != nil {
	// 	fmt.Printf("Error occured in parsing Admin list: %v\n", err)
	// 	return
	// }
	//
	// bot, err := telebot.NewBot(botConfig)
	// if err != nil {
	// 	fmt.Printf("Error occured in Bot Init: %v\n", err)
	// 	return
	// }
	//
	// bot.Handle("/start", func(c telebot.Context) error {
	// 	return c.Send("Здарова!\nПиши и получишь ответ")
	// })
	//
	// bot.Handle("/code", func(c telebot.Context) error {
	// 	fmt.Println(c.Sender().ID)
	// 	return c.Send("Аххах стырил у тебя код 😎")
	// })
	//
	// bot.Use(middleware.Whitelist(admins...))
	//
	// bot.Handle(telebot.OnText, func(c telebot.Context) error {
	// 	// TODO: НАПИСАТЬ MIDDLEWARE
	// 	if c.Chat().Type != telebot.ChatPrivate {
	// 		return nil
	// 	}
	// 	go func() {
	// 		ctx := context.Background()
	// 		openaiRequest := openai.CompletionRequest{
	// 			Model:            openai.GPT3TextDavinci003,
	// 			Prompt:           c.Text(),
	// 			MaxTokens:        MAX_TOKENS,
	// 			Temperature:      0.5,
	// 			TopP:             1.0,
	// 			FrequencyPenalty: 0,
	// 			PresencePenalty:  0,
	// 		}
	// 		resp, err := openaiClient.CreateCompletion(ctx, openaiRequest)
	// 		if err != nil {
	// 			fmt.Printf("Error occured in completion: %v\n", err)
	// 			return
	// 		}
	// 		if err = c.Send(resp.Choices[0].Text); err != nil {
	// 			fmt.Printf("Error occured in sending data to user: %v\n", err)
	// 			return
	// 		}
	// 	}()
	//
	// 	return nil
	// })
	//
	// bot.Start()
}
