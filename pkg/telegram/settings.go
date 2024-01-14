package telegram

import (
	"errors"
	"fmt"
	"strconv"

	"gopkg.in/telebot.v3"
)

func (b *Bot) showSettingsHandler(c telebot.Context) error {
	userId := c.Sender().ID
	settings, err := b.settingsStorage.Get(userId)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			settings = GetDefaultSettings()
		} else {
			return err
		}
	}
	humanSettigns := PrettySettings(settings)
	return c.Send(humanSettigns)
}

func (b *Bot) setDefaultSettingsCommandHandler(c telebot.Context) error {
	userId := c.Sender().ID
	if err := b.settingsStorage.Set(userId, GetDefaultSettings()); err != nil {
		return err
	}
	return c.Send("Установлены настройки по умолчанию")
}

func (b *Bot) setStreamCommandHandler(c telebot.Context) error {
	userId := c.Sender().ID
	settings, err := b.settingsStorage.Get(userId)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			settings = GetDefaultSettings()
		} else {
			return err
		}
	}

	settings.Stream = !settings.Stream
	err = b.settingsStorage.Set(userId, settings)
	if err != nil {
		return err
	}

	status := "Выключен"
	if settings.Stream {
		status = "Включен"
	}
	return c.Send(fmt.Sprintf("Stream-mode %s", status))
}

func (b *Bot) setTemperatureCommandHandler(c telebot.Context) error {
	userId := c.Sender().ID
	payload := c.Message().Payload
	convertedPayload, err := strconv.ParseFloat(payload, 32)
	if err != nil {
		return err
	}
	temperature := float32(convertedPayload)
	if (temperature > 1.0) || (temperature < 0) {
		return c.Send("Допустимый интервал [0; 1]")
	}

	settings, err := b.settingsStorage.Get(userId)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			settings = GetDefaultSettings()
		} else {
			return err
		}
	}
	settings.Temperature = temperature
	if err = b.settingsStorage.Set(userId, settings); err != nil {
		return err
	}
	return c.Send("Temperature установлен")
}

func (b *Bot) setFrequencyPenaltyCommandHandler(c telebot.Context) error {
	userId := c.Sender().ID
	payload := c.Message().Payload
	convertedPayload, err := strconv.ParseFloat(payload, 32)
	if err != nil {
		return err
	}
	penalty := float32(convertedPayload)
	if (penalty > 2.0) || (penalty < -2.0) {
		return c.Send("Допустимый интервал [-2.0; 2.0]")
	}

	settings, err := b.settingsStorage.Get(userId)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			settings = GetDefaultSettings()
		} else {
			return err
		}
	}
	settings.FrequencyPenalty = penalty
	if err = b.settingsStorage.Set(userId, settings); err != nil {
		return err
	}
	return c.Send("Frequency Penalty установлен")
}
