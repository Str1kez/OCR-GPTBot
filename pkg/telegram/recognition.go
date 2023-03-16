package telegram

import (
	"fmt"

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
	text, err := b.textParsing(c)
	if err != nil {
		fmt.Println("error in parsing text from image")
		return
	}
	c.Message().Text = text
	b.textCompletion(c)
}
