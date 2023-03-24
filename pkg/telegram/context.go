package telegram

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/telebot.v3"
)

func (b *Bot) contextCommandHandler(c telebot.Context) error {
	if err := b.contextStorage.Set(c.Sender().ID, c.Message().Payload); err != nil {
		log.Errorf("Error in context handler: %v\n", err)
		b.errorHandler(c.Chat().ID, errContext)
		return err
	}
	return c.Send("Контекст сохранен")
}

func (b *Bot) removeContextCommandHandler(c telebot.Context) error {
	if err := b.contextStorage.Del(c.Sender().ID); err != nil {
		log.Errorf("Error in cleaning context: %v\n", err)
		b.errorHandler(c.Chat().ID, errContext)
		return err
	}
	return c.Send("Контекст теперь пустой")
}

func (b *Bot) showContextCommandHandler(c telebot.Context) error {
	value, err := b.contextStorage.Get(c.Sender().ID)
	if err != nil {
		log.Errorf("Error in context handler: %v\n", err)
		b.errorHandler(c.Chat().ID, errContext)
		return err
	}
	if value == "" {
		return c.Send("Ваш контекст *пуст*")
	}
	return c.Send(value)
}
