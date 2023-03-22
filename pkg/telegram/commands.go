package telegram

import (
	"fmt"

	"gopkg.in/telebot.v3"
)

func (b *Bot) InitCommands() {
	startCommand := telebot.Command{Text: "start", Description: "Начало работы"}
	codeCommand := telebot.Command{Text: "code", Description: "Получение кода для аутентификации"}
	helpCommand := telebot.Command{Text: "help", Description: "Помощь"}
	contextCommand := telebot.Command{Text: "context", Description: "Задать контекст запросов"}
	removeContextCommand := telebot.Command{Text: "context_remove", Description: "Очистить контекст запросов"}
	showContextCommand := telebot.Command{Text: "context_show", Description: "Показать контекст запросов"}

	b.bot.SetCommands([]telebot.Command{
		startCommand, codeCommand,
		helpCommand, contextCommand,
		removeContextCommand, showContextCommand})
	b.bot.Handle("/start", b.startCommandHandler)
	b.bot.Handle("/code", b.codeCommandHandler)
	b.bot.Handle("/help", b.helpCommandHandler)
	b.bot.Handle("/context", b.contextCommandHandler)
	b.bot.Handle("/context_remove", b.removeContextCommandHandler)
	b.bot.Handle("/context_show", b.showContextCommandHandler)
}

func (b *Bot) startCommandHandler(c telebot.Context) error {
	return c.Send(b.config.Commands.Start)
}

func (b *Bot) codeCommandHandler(c telebot.Context) error {
	message := fmt.Sprintf("Пользователь @%s - <code>%d</code>\nХочет зарегистрироваться",
		c.Sender().Username, c.Sender().ID)
	if err := b.sendToAdmins(message); err != nil {
		return err
	}
	return c.Send(b.config.Commands.Code)
}

func (b *Bot) helpCommandHandler(c telebot.Context) error {
	return c.Send(b.config.Commands.Help)
}
