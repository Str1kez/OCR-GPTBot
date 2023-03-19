package telegram

import (
	"errors"

	log "github.com/sirupsen/logrus"
	"gopkg.in/telebot.v3"
)

var (
	errCompletion = errors.New("error occured in completion")
	errSending    = errors.New("error occured in sending data to user")
	errConverting = errors.New("error in convert photo to bytes")
	errParsing    = errors.New("error in parsing text from image")
)

func (b *Bot) errorHandler(chatId int64, e error) {
	chat := telebot.ChatID(chatId)
	switch e {
	case errCompletion:
		b.bot.Send(chat, b.config.Errors.Completion)
	case errSending:
		b.bot.Send(chat, b.config.Errors.Sending)
	case errConverting:
		b.bot.Send(chat, b.config.Errors.Converting)
	case errParsing:
		b.bot.Send(chat, b.config.Errors.Parsing)
	}
	log.Debugln("error has been handled")
}
