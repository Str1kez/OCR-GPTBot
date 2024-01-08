package config

import (
	"errors"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	OpenAIToken string    `mapstructure:"OPENAI_TOKEN"`
	OCR         OCRConfig `mapstructure:"ocr"`
	Bot         BotConfig `mapstructure:"bot"`
	Storage     StorageConfig
}

type BotConfig struct {
	Token             string        `mapstructure:"BOT_TOKEN"`
	Admins            []int64       `mapstructure:"ADMINS"`
	Users             []int64       `mapstructure:"USERS"`
	Poller            string        `mapstructure:"POLLER"`
	WebhookURL        string        `mapstructure:"WEBHOOK_URL"`
	ListenWebhookPort uint64        `mapstructure:"LISTEN_WEBHOOK_PORT"`
	Commands          CommandConfig `mapstructure:"commands"`
	Errors            ErrorConfig   `mapstructure:"errors"`
}

type OCRConfig struct {
	YandexToken  string   `mapstructure:"YANDEX_OCR_TOKEN"`
	OCRUrl       string   `mapstructure:"ocr_url"`
	JQTemplate   string   `mapstructure:"jq_template"`
	OCRLanguages []string `mapstructure:"ocr_languages"`
}

type StorageConfig struct {
	URL      string `mapstructure:"STORAGE_URL"`
	Password string `mapstructure:"STORAGE_PASSWORD"`
	DB       int    `mapstructure:"STORAGE_DB"`
}

type CommandConfig struct {
	Start        string `mapstructure:"start"`
	Code         string `mapstructure:"code"`
	Help         string `mapstructure:"help"`
	SettingsHelp string `mapstructure:"settings_help"`
}

type ErrorConfig struct {
	Completion string `mapstructure:"completion"`
	Sending    string `mapstructure:"sending"`
	Converting string `mapstructure:"converting"`
	Parsing    string `mapstructure:"parsing"`
	Context    string `mapstructure:"context"`
	Settings   string `mapstructure:"settings"`
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
	err = viper.Unmarshal(&config.Storage)
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
	if os.Getenv("MODE") == "dev" {
		viper.SetConfigFile(".local.env")
	} else {
		// if you want to start without docker
		viper.SetConfigFile(".env")
	}
	if err := viper.ReadInConfig(); err == nil {
		return nil
	}

	// for docker
	err := errors.Join(
		viper.BindEnv("OPENAI_TOKEN"),
		viper.BindEnv("BOT_TOKEN"),
		viper.BindEnv("ADMINS"),
		viper.BindEnv("USERS"),
		viper.BindEnv("POLLER"),
		viper.BindEnv("WEBHOOK_URL"),
		viper.BindEnv("LISTEN_WEBHOOK_PORT"),
		viper.BindEnv("YANDEX_OCR_TOKEN"),
		viper.BindEnv("STORAGE_URL"),
	)
	return err
}

func parseConfig() error {
	viper.SetConfigName("main")
	viper.AddConfigPath("config")
	return viper.ReadInConfig()
}
