package telegram

import (
	"errors"

	"gopkg.in/telebot.v3"
)

func (b *Bot) contextCommandHandler(c telebot.Context) error {
	userId := c.Sender().ID
	settings, err := b.settingsStorage.Get(userId)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			settings = GetDefaultSettings()
		} else {
			return err
		}
	}
	settings.Context = c.Message().Payload
	if err = b.settingsStorage.Set(userId, settings); err != nil {
		return err
	}
	return c.Send("Контекст сохранен")
}

func (b *Bot) removeContextCommandHandler(c telebot.Context) error {
	userId := c.Sender().ID
	settings, err := b.settingsStorage.Get(userId)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			settings = GetDefaultSettings()
		} else {
			return err
		}
	}
	settings.Context = ""
	if err = b.settingsStorage.Set(userId, settings); err != nil {
		return err
	}
	return c.Send("Контекст теперь пустой")
}
