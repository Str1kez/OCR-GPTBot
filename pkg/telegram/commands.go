package telegram

import (
	"fmt"

	"gopkg.in/telebot.v3"
)

func (b *Bot) InitCommands() {
	startCommand := telebot.Command{Text: "start", Description: "Начало работы"}
	codeCommand := telebot.Command{Text: "code", Description: "Получение кода для аутентификации"}
	b.bot.SetCommands([]telebot.Command{startCommand, codeCommand})
	b.bot.Handle("/start", startCommandHandler)
	b.bot.Handle("/code", codeCommandHandler)
}

func startCommandHandler(c telebot.Context) error {
	return c.Send("Здарова!\nПиши и получишь ответ")
}

func codeCommandHandler(c telebot.Context) error {
	fmt.Println(c.Sender().ID)
	return c.Send("Аххах стырил у тебя код 😎")
}
