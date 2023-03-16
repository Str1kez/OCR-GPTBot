package ocr

import "github.com/Str1kez/chatGPT-bot/internal/config"

type TesseractClient struct{}
type YandexOCRClient struct {
	config *config.OCRConfig
}

func NewTesseractClient() *TesseractClient {
	return &TesseractClient{}
}

func NewYandexOCRClient(config *config.OCRConfig) *YandexOCRClient {
	return &YandexOCRClient{config}
}
