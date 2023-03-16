package telegram

import (
	"gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/middleware"
)

func (b *Bot) InitHandlers(admins []int64) {
	authorizedGroup := b.bot.Group()
	authorizedGroup.Use(middleware.Whitelist(admins...))
	authorizedGroup.Use(PrivateChatOnly)

	authorizedGroup.Handle(telebot.OnText, b.completionOnTextHandler)
	authorizedGroup.Handle(telebot.OnPhoto, b.completionOnPhotoHandler)
}

func (b *Bot) completionOnTextHandler(c telebot.Context) error {
	go b.textCompletion(c)
	return nil
}

func (b *Bot) completionOnPhotoHandler(c telebot.Context) error {
	go b.photoCompletion(c)
	return nil
}
