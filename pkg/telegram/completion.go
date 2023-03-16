package telegram

import (
	"fmt"
	"io"

	"gopkg.in/telebot.v3"
)

func (b *Bot) completionRequest(text string) (string, error) {
	response, err := b.openAI.CreateChatCompletion(text)
	if err != nil {
		return "", err
	}
	return response.Choices[0].Message.Content, nil
}

func (b *Bot) textCompletion(c telebot.Context) {
	content, err := b.completionRequest(c.Text())
	if err != nil {
		fmt.Printf("Error occured in completion: %v\n", err)
		return
	}
	// if err = c.Send(response.Choices[0].Text); err != nil {
	if err = c.Send(content); err != nil {
		fmt.Printf("Error occured in sending data to user: %v\n", err)
		return
	}
}

func (b *Bot) photoCompletion(c telebot.Context) {
	file := c.Message().Photo.MediaFile()
	bytesPhoto, err := b.getPhotoInByte(file)
	if err != nil {
		fmt.Println("error in convert to bytes")
		return
	}
	text, err := b.recognitionClient.RecognitionFromBytes(bytesPhoto)
	if err != nil {
		fmt.Println("error in parsing text from image")
		return
	}
	c.Message().Text = text
	b.textCompletion(c)
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
