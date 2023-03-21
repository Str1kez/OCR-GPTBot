package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	OpenAIToken string    `mapstructure:"OPENAI_TOKEN"`
	OCR         OCRConfig `mapstructure:"ocr"`
	Bot         BotConfig `mapstructure:"bot"`
}

type BotConfig struct {
	Token    string        `mapstructure:"BOT_TOKEN"`
	Admins   []int64       `mapstructure:"ADMINS"`
	Users    []int64       `mapstructure:"USERS"`
	Messages MessageConfig `mapstructure:"messages"`
	Errors   ErrorConfig   `mapstructure:"errors"`
}

type OCRConfig struct {
	YandexToken  string   `mapstructure:"YANDEX_OCR_TOKEN"`
	OCRUrl       string   `mapstructure:"ocr_url"`
	JQTemplate   string   `mapstructure:"jq_template"`
	OCRLanguages []string `mapstructure:"ocr_languages"`
}

type MessageConfig struct {
	Start string `mapstructure:"start"`
	Code  string `mapstructure:"code"`
}

type ErrorConfig struct {
	Completion string `mapstructure:"completion"`
	Sending    string `mapstructure:"sending"`
	Converting string `mapstructure:"converting"`
	Parsing    string `mapstructure:"parsing"`
}

func NewConfig() (*Config, error) {
	var config Config

	err := parseEnv()
	if err != nil {
		log.Errorf("Error in parsing env data: %v\n", err)
		return nil, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Errorf("Error in unmarshalling env data: %v\n", err)
		return nil, err
	}
	err = viper.Unmarshal(&config.Bot)
	if err != nil {
		log.Errorf("Error in unmarshalling env data: %v\n", err)
		return nil, err
	}
	err = viper.Unmarshal(&config.OCR)
	if err != nil {
		log.Errorf("Error in unmarshalling env data: %v\n", err)
		return nil, err
	}

	err = parseConfig()
	if err != nil {
		log.Errorf("Error in parsing yaml config: %v\n", err)
		return nil, err
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Errorf("Error in unmarshalling yaml config: %v\n", err)
		return nil, err
	}
	return &config, nil
}

func parseEnv() error {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err == nil {
		return nil
	}
	if err := viper.BindEnv("OPENAI_TOKEN"); err != nil {
		return err
	}
	if err := viper.BindEnv("BOT_TOKEN"); err != nil {
		return err
	}
	if err := viper.BindEnv("ADMINS"); err != nil {
		return err
	}
	if err := viper.BindEnv("USERS"); err != nil {
		return err
	}
	if err := viper.BindEnv("YANDEX_OCR_TOKEN"); err != nil {
		return err
	}
	return nil

}

func parseConfig() error {
	viper.SetConfigName("main")
	viper.AddConfigPath("config")
	return viper.ReadInConfig()
}
