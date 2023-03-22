package telegram

import (
	log "github.com/sirupsen/logrus"

	"gopkg.in/telebot.v3"
)

func (b *Bot) textCompletion(c telebot.Context) error {
	content, err := b.completionClient.PerformCompletion(c.Text())
	if err != nil {
		log.Errorf("Error occured in completion: %v\n", err)
		return errCompletion
	}
	log.Debugln("Completion success")
	if err = c.Send(content); err != nil {
		log.Errorf("Error occured in sending data to user: %v\n", err)
		return errSending
	}
	log.Debugln("Message has been sent")
	return nil
}
