package telegram

import (
	"gopkg.in/telebot.v3"
)

func (b *Bot) contextCommandHandler(c telebot.Context) error {
	if err := b.settingsStorage.Set(c.Sender().ID, "context", c.Message().Payload); err != nil {
		return err
	}
	return c.Send("Контекст сохранен")
}

func (b *Bot) removeContextCommandHandler(c telebot.Context) error {
	if err := b.settingsStorage.Del(c.Sender().ID, "context"); err != nil {
		return err
	}
	return c.Send("Контекст теперь пустой")
}
