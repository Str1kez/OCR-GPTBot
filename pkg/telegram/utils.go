package telegram

import "strings"

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
