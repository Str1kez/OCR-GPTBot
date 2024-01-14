package telegram

import (
	"errors"

	"github.com/Str1kez/OCR-GPTBot/internal/config"

	log "github.com/sirupsen/logrus"
	"gopkg.in/telebot.v3"
)

var (
	errCompletion = errors.New("error occured in completion")
	errSending    = errors.New("error occured in sending data to user")
	errConverting = errors.New("error in convert photo to bytes")
	errParsing    = errors.New("error in parsing text from image")
	errContext    = errors.New("error in interaction with context storage")
	errSettings   = errors.New("error in interaction with storage settings")
	ErrNotFound   = errors.New("data in storage not found")
)

func NewErrorHandler(errorConfig config.ErrorConfig) func(err error, c telebot.Context) {
	return func(err error, c telebot.Context) {
		if c == nil {
			log.Warningln(err)
			return
		}

		var e error

		switch {
		case errors.Is(err, errCompletion):
			e = c.Send(errorConfig.Completion)
		case errors.Is(err, errSending):
			e = c.Send(errorConfig.Sending)
		case errors.Is(err, errConverting):
			e = c.Send(errorConfig.Converting)
		case errors.Is(err, errParsing):
			e = c.Send(errorConfig.Parsing)
		case errors.Is(err, errContext):
			e = c.Send(errorConfig.Context)
		case errors.Is(err, errSettings):
			e = c.Send(errorConfig.Settings)
		case errors.Is(err, ErrNotFound):
			e = c.Send(errorConfig.Settings)
		default:
			e = c.Send("Непредвиденная ошибка")
		}
		log.Errorln("Update ID:", c.Update().ID, err)

		if e != nil {
			log.Fatalf("Error on error handling: %v\n", e)
		}
	}
}
