package telegram

import (
	"fmt"
	"strconv"

	log "github.com/sirupsen/logrus"
	"gopkg.in/telebot.v3"
)

func (b *Bot) showSettingsHandler(c telebot.Context) error {
	userId := c.Sender().ID
	settings, err := b.settingsStorage.GetAll(userId)
	if err != nil {
		log.Errorf("Couldn't handle settings: %v\n", err)
		b.errorHandler(userId, errSettings)
		return err
	}
	humanSettigns := PrettySettings(settings)
	return c.Send(humanSettigns)
}

func (b *Bot) setDefaultSettingsCommandHandler(c telebot.Context) error {
	userId := c.Sender().ID
	if err := b.settingsStorage.DelAll(userId); err != nil {
		log.Errorf("Couldn't handle settings: %v\n", err)
		b.errorHandler(userId, errSettings)
		return err
	}
	return c.Send("Установлены настройки по умолчанию")
}

func (b *Bot) setStreamCommandHandler(c telebot.Context) error {
	userId := c.Sender().ID
	data, err := b.settingsStorage.Get(userId, "stream")
	if err != nil {
		log.Errorf("Couldn't handle settings: %v\n", err)
		b.errorHandler(userId, errSettings)
		return err
	}
	isStream, err := strconv.ParseBool(string(data))
	if err != nil {
		log.Errorf("Couldn't handle settings: %v\n", err)
		b.errorHandler(userId, errSettings)
		return err
	}
	isStream = !isStream
	err = b.settingsStorage.Set(userId, "stream", isStream)
	if err != nil {
		log.Errorf("Couldn't handle settings: %v\n", err)
		b.errorHandler(userId, errSettings)
		return err
	}
	status := "Выключен"
	if isStream {
		status = "Включен"
	}
	return c.Send(fmt.Sprintf("Stream-mode %s", status))
}

func (b *Bot) setTemperatureCommandHandler(c telebot.Context) error {
	userId := c.Sender().ID
	payload := c.Message().Payload
	convertedPayload, err := strconv.ParseFloat(payload, 32)
	if err != nil {
		log.Errorf("Couldn't handle settings: %v\n", err)
		b.errorHandler(userId, errSettings)
		return err
	}
	temperature := float32(convertedPayload)
	if (temperature > 1.0) || (temperature < 0) {
		return c.Send("Допустимый интервал [0; 1]")
	}
	err = b.settingsStorage.Set(userId, "temperature", temperature)
	if err != nil {
		log.Errorf("Couldn't handle settings: %v\n", err)
		b.errorHandler(userId, errSettings)
		return err
	}
	return c.Send("Temperature установлен")
}

func (b *Bot) setFrequencyPenaltyCommandHandler(c telebot.Context) error {
	userId := c.Sender().ID
	payload := c.Message().Payload
	convertedPayload, err := strconv.ParseFloat(payload, 32)
	if err != nil {
		log.Errorf("Couldn't handle settings: %v\n", err)
		b.errorHandler(userId, errSettings)
		return err
	}
	penalty := float32(convertedPayload)
	if (penalty > 2.0) || (penalty < -2.0) {
		return c.Send("Допустимый интервал [-2.0; 2.0]")
	}
	err = b.settingsStorage.Set(userId, "frequency_penalty", penalty)
	if err != nil {
		log.Errorf("Couldn't handle settings: %v\n", err)
		b.errorHandler(userId, errSettings)
		return err
	}
	return c.Send("Frequency Penalty установлен")
}
