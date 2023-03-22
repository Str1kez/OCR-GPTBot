package telegram

import (
	log "github.com/sirupsen/logrus"

	"gopkg.in/telebot.v3"
)

func (b *Bot) textCompletion(c telebot.Context) error {
	userContext, err := b.contextStorage.Get(c.Sender().ID)
	if err != nil {
		log.Errorf("Error in storage: %v\n", err)
		return errContext
	}
	content, err := b.completionClient.PerformCompletion(c.Text(), userContext)
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
