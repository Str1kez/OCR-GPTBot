package telegram

import (
	"fmt"

	"gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/middleware"
)

func (b *Bot) InitCommands() {
	startCommand := telebot.Command{Text: "start", Description: "Начало работы"}
	codeCommand := telebot.Command{Text: "code", Description: "Получение кода для аутентификации"}
	b.bot.SetCommands([]telebot.Command{startCommand, codeCommand})
}

func (b *Bot) InitHandlers(admins []int64) {
	authorizedGroup := b.bot.Group()
	authorizedGroup.Use(middleware.Whitelist(admins...))
	authorizedGroup.Use(PrivateChatOnly)

	b.bot.Handle("/start", startHandler)
	b.bot.Handle("/code", codeHandler)

	authorizedGroup.Handle(telebot.OnText, b.completionOnTextHandler)
	authorizedGroup.Handle(telebot.OnPhoto, b.completionOnPhotoHandler)
}

func startHandler(c telebot.Context) error {
	return c.Send("Здарова!\nПиши и получишь ответ")
}

func codeHandler(c telebot.Context) error {
	fmt.Println(c.Sender().ID)
	return c.Send("Аххах стырил у тебя код 😎")
}

func (b *Bot) completionOnTextHandler(c telebot.Context) error {
	go b.textCompletion(c)
	return nil
}

func (b *Bot) completionOnPhotoHandler(c telebot.Context) error {
	go b.photoCompletion(c)
	return nil
}
