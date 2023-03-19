package telegram

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/telebot.v3"
)

func (b *Bot) InitCommands() {
	startCommand := telebot.Command{Text: "start", Description: "Начало работы"}
	codeCommand := telebot.Command{Text: "code", Description: "Получение кода для аутентификации"}
	b.bot.SetCommands([]telebot.Command{startCommand, codeCommand})
	b.bot.Use(PrivateChatOnly)
	// b.bot.Use(Logging(log.New()))
	b.bot.Handle("/start", b.startCommandHandler)
	b.bot.Handle("/code", b.codeCommandHandler)
}

func (b *Bot) startCommandHandler(c telebot.Context) error {
	return c.Send(b.config.Messages.Start)
}

func (b *Bot) codeCommandHandler(c telebot.Context) error {
	log.Infoln(c.Sender().ID)
	return c.Send(b.config.Messages.Code)
}
