package ocr

import "github.com/Str1kez/OCR-GPTBot/internal/config"

type (
	TesseractClient struct{}
	YandexOCRClient struct {
		config *config.OCRConfig
	}
)

func NewTesseractClient() *TesseractClient {
	return &TesseractClient{}
}

func NewYandexOCRClient(config *config.OCRConfig) *YandexOCRClient {
	return &YandexOCRClient{config}
}
