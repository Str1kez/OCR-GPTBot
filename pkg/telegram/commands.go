package telegram

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"gopkg.in/telebot.v3"
)

func (b *Bot) InitCommands() {
	startCommand := telebot.Command{Text: "start", Description: "Начало работы"}
	codeCommand := telebot.Command{Text: "code", Description: "Получение кода для аутентификации"}
	helpCommand := telebot.Command{Text: "help", Description: "Помощь"}
	b.bot.SetCommands([]telebot.Command{startCommand, codeCommand, helpCommand})
	b.bot.Handle("/start", b.startCommandHandler)
	b.bot.Handle("/code", b.codeCommandHandler)
	b.bot.Handle("/help", b.helpCommandHandler)
}

func (b *Bot) startCommandHandler(c telebot.Context) error {
	return c.Send(b.config.Messages.Start)
}

func (b *Bot) codeCommandHandler(c telebot.Context) error {
	log.Infoln(c.Sender().ID)
	message := fmt.Sprintf("Пользователь @%s - <code>%d</code>\nХочет зарегистрироваться", c.Sender().Username, c.Sender().ID)
	if err := b.sendToAdmins(message); err != nil {
		return err
	}
	return c.Send(b.config.Messages.Code)
}

func (b *Bot) helpCommandHandler(c telebot.Context) error {
	return c.Send(b.config.Messages.Help)
}
