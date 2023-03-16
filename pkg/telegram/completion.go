package telegram

import (
	"fmt"
	"io"

	"gopkg.in/telebot.v3"
)

func (b *Bot) textCompletion(c telebot.Context) error {
	content, err := b.completionClient.PerformCompletion(c.Text())
	if err != nil {
		fmt.Printf("Error occured in completion: %v\n", err)
		return errCompletion
	}
	if err = c.Send(content); err != nil {
		fmt.Printf("Error occured in sending data to user: %v\n", err)
		return errSending
	}
	return nil
}

func (b *Bot) photoCompletion(c telebot.Context) error {
	file := c.Message().Photo.MediaFile()
	bytesPhoto, err := b.getPhotoInByte(file)
	if err != nil {
		fmt.Println("error in convert to bytes")
		return errConverting
	}
	text, err := b.recognitionClient.RecognitionFromBytes(bytesPhoto)
	if err != nil {
		fmt.Println("error in parsing text from image", err)
		return errParsing
	}
	c.Message().Text = text
	b.textCompletion(c)
	return nil
}

func (b *Bot) getPhotoInByte(file *telebot.File) ([]byte, error) {
	photo, err := b.bot.File(file)
	if err != nil {
		fmt.Println("error in downloading")
		return nil, err
	}
	defer photo.Close()
	return io.ReadAll(photo)
}
