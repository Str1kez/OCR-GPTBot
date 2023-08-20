package telegram

import (
	"gopkg.in/telebot.v3"

	log "github.com/sirupsen/logrus"
)

func (b *Bot) textCompletion(c telebot.Context) error {
	// TODO: Надо добавить в сторедж хранение выбора юзера типа отображения ответа
	userContext, err := b.contextStorage.Get(c.Sender().ID)
	if err != nil {
		log.Errorf("Error in storage: %v\n", err)
		return errContext
	}
	// TODO: Добавить по выбору тип ответа (вынос в отдельную)
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

	// stream, err := b.completionClient.GetCompletionStream(c.Text(), userContext)
	// if err != nil {
	// 	log.Errorf("Error occured in completion: %v\n", err)
	// 	return errCompletion
	// }
	// defer stream.Close()
	// var msg *telebot.Message = nil
	// for {
	// 	response, err := stream.Recv()
	// 	if errors.Is(err, io.EOF) {
	// 		log.Debugln("Stream finished")
	// 		return nil
	// 	} else if err != nil {
	// 		log.Errorf("Stream error: %v\n", err)
	// 		return errCompletion
	// 	}
	//
	// 	// TODO: Надо сделать чанками по словам, иначе тг ддосится
	// 	newData := response.Choices[0].Delta.Content
	// 	if msg == nil {
	// 		if newData == "" {
	// 			continue
	// 		}
	// 		msg, err = b.bot.Send(c.Chat(), newData)
	// 		if err != nil {
	// 			log.Errorf("Error occured in sending data to user: %v\n", err)
	// 			return errSending
	// 		}
	// 	} else {
	// 		msg, err = b.bot.Edit(msg, msg.Text+newData)
	// 		if err != nil {
	// 			log.Errorf("Editing message failed: %v\n", err)
	// 			return err
	// 		}
	// 	}
	// }
	log.Debugln("Message has been sent")

	return nil
}
