package ocr

import (
	"fmt"

	"github.com/otiai10/gosseract/v2"
)

func (t *TesseractClient) RecognitionFromBytes(photo []byte) (string, error) {
	client := gosseract.NewClient()
	defer client.Close()
	client.SetLanguage("eng", "rus")
	err := client.SetImageFromBytes(photo)
	if err != nil {
		fmt.Println("error in setting image")
		return "", err
	}
	return client.Text()
}
