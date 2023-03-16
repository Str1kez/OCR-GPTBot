package telegram

import (
	"gopkg.in/telebot.v3"
)

func PrivateChatOnly(next telebot.HandlerFunc) telebot.HandlerFunc {
	return func(c telebot.Context) error {
		if c.Chat().Type == telebot.ChatPrivate {
			return next(c)
		}
		return nil
	}
}
