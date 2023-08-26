package ocr

import (
	"encoding/json"

	"github.com/itchyny/gojq"
)

func (y *YandexOCRClient) GetTextFromImage(image []byte) (string, error) {
	responseBody, err := y.AnalyzeRequest(image)
	if err != nil {
		return "", err
	}
	var parsedJson map[string]interface{}
	err = json.Unmarshal(responseBody, &parsedJson)
	if err != nil {
		return "", err
	}

	query, err := gojq.Parse(y.config.JQTemplate)
	if err != nil {
		return "", err
	}
	iter := query.Run(parsedJson)
	for {
		v, ok := iter.Next()
		if !ok {
			break
		}
		if err, ok := v.(error); ok {
			return "", err
		}
		if s, ok := v.(string); ok {
			return s, nil
		}
	}
	return "", nil
}
