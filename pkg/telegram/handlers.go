package telegram

import (
	"gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/middleware"
)

func (b *Bot) InitHandlers() {
	authorizedGroup := b.bot.Group()
	authorizedGroup.Use(middleware.Whitelist(b.config.Admins...))

	authorizedGroup.Handle(telebot.OnText, b.completionOnTextHandler)
	authorizedGroup.Handle(telebot.OnPhoto, b.completionOnPhotoHandler)
}

func (b *Bot) completionOnTextHandler(c telebot.Context) error {
	go func() {
		err := b.textCompletion(c)
		if err != nil {
			b.errorHandler(c.Chat().ID, err)
		}
	}()
	return nil
}

func (b *Bot) completionOnPhotoHandler(c telebot.Context) error {
	go func() {
		err := b.photoCompletion(c)
		if err != nil {
			b.errorHandler(c.Chat().ID, err)
		}
	}()
	return nil
}
