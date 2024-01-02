package telegram

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Str1kez/OCR-GPTBot/internal/config"
	"gopkg.in/telebot.v3"
)

func GetDefaultSettings() Settings {
	return Settings{
		Context:          DefaultContext,
		Stream:           false,
		Temperature:      DefaultTemperature,
		FrequencyPenalty: DefaultFrequencyPenalty,
	}
}

// IsCorrectRepresentation Maybe will be useless
func IsCorrectRepresentation(data []rune) bool {
	strData := string(data)
	return strings.HasSuffix(strData, "\n") || strings.HasPrefix(strData, "\n") ||
		strings.HasSuffix(strData, " ") || strings.HasPrefix(strData, " ")
}

func PrettySettings(s Settings) string {
	result := make([]string, 0, 4)
	translatedStream := "Выключено"
	if s.Stream {
		translatedStream = "Включено"
	}

	result = append(result, fmt.Sprintf("Контекст: %s", s.Context))
	result = append(result, fmt.Sprintf("Стрим: %v", translatedStream))
	result = append(result, fmt.Sprintf("Temperature: %.2f", s.Temperature))
	result = append(result, fmt.Sprintf("Frequency Penalty: %.2f", s.FrequencyPenalty))
	return strings.Join(result, "\n")
}

func GetPoller(cfg config.BotConfig) (telebot.Poller, error) {
	switch cfg.Poller {
	case "long":
		return &telebot.LongPoller{Timeout: time.Minute}, nil
	case "webhook":
		return &telebot.Webhook{
			Listen:      ":" + strconv.FormatUint(cfg.ListenWebhookPort, 10),
			Endpoint:    &telebot.WebhookEndpoint{PublicURL: cfg.WebhookURL},
			DropUpdates: true,
		}, nil
	}
	return nil, errors.New("unrecognized type of poller")
}
