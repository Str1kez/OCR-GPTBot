package ocr

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

type BatchAnalyzeRequest struct {
	AnalyzeSpecs []AnalyzeSpec `json:"analyzeSpecs"`
}

type AnalyzeSpec struct {
	Content  string    `json:"content"`
	Features []Feature `json:"features"`
}

type Feature struct {
	Type                string `json:"type"`
	TextDetectionConfig `json:"textDetectionConfig"`
}

type TextDetectionConfig struct {
	LanguageCodes []string `json:"languageCodes"`
}

func (y *YandexOCRClient) generateAnalyzePayload(photo64 string) *BatchAnalyzeRequest {
	return &BatchAnalyzeRequest{
		AnalyzeSpecs: []AnalyzeSpec{
			{
				Content: photo64,
				Features: []Feature{
					{
						Type: "TEXT_DETECTION",
						TextDetectionConfig: TextDetectionConfig{
							LanguageCodes: y.config.OCRLanguages,
						},
					},
				},
			},
		},
	}
}

func (y *YandexOCRClient) AnalyzeRequest(photo []byte) ([]byte, error) {
	photo64 := base64.RawStdEncoding.EncodeToString(photo)
	payload := y.generateAnalyzePayload(photo64)
	json, err := json.Marshal(payload)
	if err != nil {
		return nil, errors.New("failed to marshalling payload")
	}

	request, err := http.NewRequest(http.MethodPost, y.config.OCRUrl, bytes.NewReader(json))
	if err != nil {
		return nil, errors.New("preparing request has been failed")
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Api-Key "+y.config.YandexToken)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, errors.New("failed request on recognition")
	}
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("yandex returned error http code %v", response.StatusCode)
	}
	defer func() {
		err = response.Body.Close()
		if err != nil {
			log.Panic(err)
		}
	}()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, errors.New("failed representation body in bytes")
	}

	return responseBody, nil
}
