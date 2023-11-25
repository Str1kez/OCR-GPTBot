package ocr

import (
	"errors"

	"github.com/otiai10/gosseract/v2"
	log "github.com/sirupsen/logrus"
)

func (t *TesseractClient) GetTextFromImage(image []byte) (result string, err error) {
	client := gosseract.NewClient()
	defer func() {
		if deferErr := client.Close(); deferErr != nil {
			log.Errorln("error in closing tesseract client")
			err = errors.Join(err, deferErr)
		}
	}()

	err = client.SetLanguage("eng", "rus")
	if err != nil {
		log.Errorln("error in setting language")
		return
	}
	err = client.SetImageFromBytes(image)
	if err != nil {
		log.Errorln("error in setting image")
		return
	}
	return client.Text()
}
