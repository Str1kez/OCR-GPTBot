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
	return b.textCompletion(c)
}

func (b *Bot) recognitionOnPhotoHandler(c telebot.Context) error {
	return b.textRecognition(c)
}
