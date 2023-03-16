package telegram

import (
	"fmt"
	"io"

	"github.com/otiai10/gosseract/v2"
	"gopkg.in/telebot.v3"
)

func (b *Bot) textParsing(c telebot.Context) (string, error) {
	client := gosseract.NewClient()
	defer client.Close()
	client.SetLanguage("eng", "rus")
	file := c.Message().Photo.MediaFile()
	bytesPhoto, err := b.getPhotoInByte(file)
	if err != nil {
		fmt.Println("error in convert to bytes")
		return "", err
	}
	err = client.SetImageFromBytes(bytesPhoto)
	if err != nil {
		fmt.Println("error in setting image")
		return "", err
	}
	return client.Text()
	// if err != nil {
	//   fmt.Println("error in parsing text from image")
	//   return "", err
	// }
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
