package telegram

import (
	"io"

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

func (b *Bot) textRecognition(c telebot.Context) error {
	file := c.Message().Photo.MediaFile()
	bytesPhoto, err := b.getPhotoInByte(file)
	if err != nil {
		log.Errorf("error in convert to bytes: %v\n", err)
		return errConverting
	}
	log.Debugln("Photo has been converted in bytes")
	text, err := b.recognitionClient.RecognitionFromBytes(bytesPhoto)
	if err != nil {
		log.Errorf("error in parsing text from image: %v\n", err)
		return errParsing
	}
	log.Debugln("Text has been parsed")
	log.Debugf("Recognized text: %v\n", text)
	if err := c.Send(text); err != nil {
		log.Errorf("error in text recognition response to client: %v\n", err)
	}
	log.Debugln("Text has been sent to user")
	return nil
}

func (b *Bot) getPhotoInByte(file *telebot.File) ([]byte, error) {
	photo, err := b.bot.File(file)
	if err != nil {
		log.Errorf("error in downloading: %v\n", err)
		return nil, err
	}
	log.Debugln("File has been downloaded")
	defer photo.Close()
	return io.ReadAll(photo)
}
