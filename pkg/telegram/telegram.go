package telegram

import (
	"github.com/Str1kez/OCR-GPTBot/internal/config"
	log "github.com/sirupsen/logrus"
	"gopkg.in/telebot.v3"
)

type Bot struct {
	bot               *telebot.Bot
	config            *config.BotConfig
	completionClient  completion
	recognitionClient recognition
	contextStorage    storage
}

type recognition interface {
	RecognitionFromBytes(photo []byte) (string, error)
}

type completion interface {
	PerformCompletion(text, userContext string) (string, error)
}

type storage interface {
	Get(key int64) (string, error)
	Set(key int64, value string) error
	Del(key int64) error
}

func NewBot(settings telebot.Settings,
	config *config.BotConfig,
	chat completion,
	recognitionClient recognition,
	contextStorage storage) (*Bot, error) {
	bot, err := telebot.NewBot(settings)
	if err != nil {
		return nil, err
	}
	return &Bot{
			bot:               bot,
			config:            config,
			completionClient:  chat,
			recognitionClient: recognitionClient,
			contextStorage:    contextStorage},
		nil
}

func (b *Bot) Start() {
	b.bot.Start()
}

func (b *Bot) sendToAdmins(message string) error {

	for _, admin := range b.config.Admins {
		user := telebot.User{ID: admin}
		if _, err := b.bot.Send(&user, message); err != nil {
			log.Errorf("Couldn't notify admin: ID=%d\n", user.ID)
			return err
		}
	}
	return nil
}

func (b *Bot) OnStartup() error {
	b.bot.Use(PrivateChatOnly)
	// b.bot.Use(Logging(log.New()))
	if err := b.sendToAdmins("Бот начал работу"); err != nil {
		return err
	}
	log.Debugln("Admins have been notificated")
	return nil
}
