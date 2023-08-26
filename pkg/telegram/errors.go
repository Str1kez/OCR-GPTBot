package telegram

import (
	"errors"

	log "github.com/sirupsen/logrus"
	"gopkg.in/telebot.v3"
)

var (
	errCompletion          = errors.New("error occured in completion")
	errSending             = errors.New("error occured in sending data to user")
	errConverting          = errors.New("error in convert photo to bytes")
	errParsing             = errors.New("error in parsing text from image")
	errContext             = errors.New("error in interaction with context storage")
	errInterfaceConversion = errors.New("error in interface conversion")
)

func (b *Bot) errorHandler(chatId int64, e error) {
	chat := telebot.ChatID(chatId)
	var err error = nil

	switch {
	case errors.Is(e, errCompletion):
		_, err = b.bot.Send(chat, b.config.Errors.Completion)
	case errors.Is(e, errSending):
		_, err = b.bot.Send(chat, b.config.Errors.Sending)
	case errors.Is(e, errConverting):
		_, err = b.bot.Send(chat, b.config.Errors.Converting)
	case errors.Is(e, errParsing):
		_, err = b.bot.Send(chat, b.config.Errors.Parsing)
	case errors.Is(e, errContext):
		_, err = b.bot.Send(chat, b.config.Errors.Context)
	default:
		_, err = b.bot.Send(chat, "Непредвиденная ошибка")
	}

	if err != nil {
		log.Fatalf("Error on error handling: %v\n", err)
	}
	log.Debugf("error has been handled: %v\n", e)
}
