package telegram

import (
	"errors"
	"io"
	"strings"

	"gopkg.in/telebot.v3"

	log "github.com/sirupsen/logrus"
)

func (b *Bot) textCompletion(c telebot.Context) error {
	userSettings, err := b.settingsStorage.GetAll(c.Sender().ID)
	if err != nil {
		log.Errorf("Couldn't get user settings. User: %d\n%v\n", c.Sender().ID, err)
		return errContext
	}

	if userSettings.Stream {
		return b.streamTextCompletion(c, userSettings)
	}
	return b.fullTextCompletion(c, userSettings)
}

func (b *Bot) fullTextCompletion(c telebot.Context, settings Settings) error {
	content, err := b.completionClient.GetCompletionText(c.Text(), settings)
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

func (b *Bot) streamTextCompletion(c telebot.Context, settings Settings) error {
	stream, err := b.completionClient.GetCompletionStream(c.Text(), settings)
	if err != nil {
		log.Errorf("Error occured in completion: %v\n", err)
		return errCompletion
	}
	defer stream.Close()

	var msg *telebot.Message = nil
	buf := make([]rune, 0, 40)
	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			if len(buf) > 0 {
				_, err := b.bot.Edit(msg, msg.Text+string(buf))
				if err != nil && !strings.Contains(err.Error(), "message is not modified") {
					log.Errorf("Editing message failed: %v\n", err)
					return err
				}
			}
			log.Debugln("Stream finished")
			return nil
		} else if err != nil {
			log.Errorf("Stream error: %v\n", err)
			return errCompletion
		}

		newData := response.Choices[0].Delta.Content
		if len(buf) >= 30 && !(strings.HasSuffix(newData, "\n") || strings.HasSuffix(newData, " ")) {
			messageText := string(buf) + newData
			if msg == nil {
				msg, err = b.bot.Send(c.Chat(), messageText)
				if err != nil {
					log.Errorf("Error occured in sending data to user: %v\n", err)
					return errSending
				}
				log.Debugln("Message has been sent")
			} else {
				if messageText == "" {
					log.Debugln("empty message")
				}
				msg, err = b.bot.Edit(msg, msg.Text+messageText)
				if err != nil && !strings.Contains(err.Error(), "message is not modified") {
					log.Errorf("Editing message failed: %v\n", err)
					return err
				}
			}
			buf = make([]rune, 0, 40)
		} else {
			buf = append(buf, []rune(newData)...)
		}
	}
}
