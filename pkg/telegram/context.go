package telegram

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/telebot.v3"
)

func (b *Bot) contextCommandHandler(c telebot.Context) error {
	if err := b.settingsStorage.Set(c.Sender().ID, "context", c.Message().Payload); err != nil {
		log.Errorf("Error in context handler: %v\n", err)
		b.errorHandler(c.Chat().ID, errContext)
		return err
	}
	return c.Send("Контекст сохранен")
}

func (b *Bot) removeContextCommandHandler(c telebot.Context) error {
	if err := b.settingsStorage.Del(c.Sender().ID, "context"); err != nil {
		log.Errorf("Error in cleaning context: %v\n", err)
		b.errorHandler(c.Chat().ID, errContext)
		return err
	}
	return c.Send("Контекст теперь пустой")
}
