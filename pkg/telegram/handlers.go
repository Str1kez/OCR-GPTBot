package telegram

import (
	"gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/middleware"
)

func (b *Bot) InitHandlers() {
	authenticatedUsers := append(b.config.Admins, b.config.Users...)
	authorizedGroup := b.bot.Group()
	authorizedGroup.Use(middleware.Whitelist(authenticatedUsers...), Logging())

	authorizedGroup.Handle(telebot.OnText, b.completionOnTextHandler)
	authorizedGroup.Handle(telebot.OnPhoto, b.recognitionOnPhotoHandler)
}

func (b *Bot) completionOnTextHandler(c telebot.Context) error {
	err := b.textCompletion(c)
	if err != nil {
		b.errorHandler(c.Chat().ID, err)
	}
	return err
}

func (b *Bot) recognitionOnPhotoHandler(c telebot.Context) error {
	err := b.textRecognition(c)
	if err != nil {
		b.errorHandler(c.Chat().ID, err)
	}
	return err
}
