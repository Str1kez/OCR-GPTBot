package telegram

import (
	"fmt"

	"gopkg.in/telebot.v3"
)

func (b *Bot) InitCommands() {
	startCommand := telebot.Command{Text: "start", Description: "Начало работы"}
	codeCommand := telebot.Command{Text: "code", Description: "Получение кода для аутентификации"}
	helpCommand := telebot.Command{Text: "help", Description: "Помощь"}
	helpSettingsCommand := telebot.Command{Text: "settings_help", Description: "Помощь в настройках"}
	settingsCommand := telebot.Command{Text: "settings", Description: "Настройки"}
	contextCommand := telebot.Command{Text: "context", Description: "Задать контекст запросов"}
	removeContextCommand := telebot.Command{Text: "context_remove", Description: "Очистить контекст запросов"}
	streamCommand := telebot.Command{Text: "stream", Description: "Установить stream-mode"}
	temperatureCommand := telebot.Command{Text: "temperature", Description: "Задать Temperature"}
	freqPenaltyCommand := telebot.Command{Text: "freq_penalty", Description: "Задать Frequency Penalty"}
	defaultSettingsCommand := telebot.Command{Text: "settings_default", Description: "Установить настройки по умолчанию"}

	b.bot.SetCommands([]telebot.Command{
		startCommand, codeCommand,
		helpCommand, helpSettingsCommand, settingsCommand, contextCommand,
		removeContextCommand, streamCommand, temperatureCommand, freqPenaltyCommand, defaultSettingsCommand,
	})
	b.bot.Handle("/start", b.startCommandHandler)
	b.bot.Handle("/code", b.codeCommandHandler)
	b.bot.Handle("/help", b.helpCommandHandler)
	b.bot.Handle("/settings_help", b.helpSettingsCommandHandler)
	b.bot.Handle("/settings", b.showSettingsHandler)
	b.bot.Handle("/context", b.contextCommandHandler)
	b.bot.Handle("/context_remove", b.removeContextCommandHandler)
	b.bot.Handle("/stream", b.setStreamCommandHandler)
	b.bot.Handle("/temperature", b.setTemperatureCommandHandler)
	b.bot.Handle("/freq_penalty", b.setFrequencyPenaltyCommandHandler)
	b.bot.Handle("/settings_default", b.setDefaultSettingsCommandHandler)
}

func (b *Bot) startCommandHandler(c telebot.Context) error {
	return c.Send(b.config.Commands.Start)
}

func (b *Bot) codeCommandHandler(c telebot.Context) error {
	message := fmt.Sprintf("Пользователь @%s - %d\nХочет зарегистрироваться",
		c.Sender().Username, c.Sender().ID)
	if err := b.sendToAdmins(message); err != nil {
		return err
	}
	return c.Send(b.config.Commands.Code)
}

func (b *Bot) helpCommandHandler(c telebot.Context) error {
	return c.Send(b.config.Commands.Help)
}

func (b *Bot) helpSettingsCommandHandler(c telebot.Context) error {
	return c.Send(b.config.Commands.SettingsHelp)
}
