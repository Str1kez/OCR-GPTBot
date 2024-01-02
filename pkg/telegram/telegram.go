package telegram

import (
	"errors"

	"github.com/Str1kez/OCR-GPTBot/internal/config"
	log "github.com/sirupsen/logrus"
	"gopkg.in/telebot.v3"
)

func NewBot(tgSettings telebot.Settings,
	config *config.BotConfig,
	chat completion,
	recognitionClient recognition,
	settings storage,
) (*Bot, error) {
	bot, err := telebot.NewBot(tgSettings)
	if err != nil {
		return nil, err
	}
	return &Bot{
			bot:               bot,
			config:            config,
			completionClient:  chat,
			recognitionClient: recognitionClient,
			settingsStorage:   settings,
		},
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

func (b *Bot) OnShutdown() error {
	err := b.sendToAdmins("Бот закончил работу")
	errWebhook := b.bot.RemoveWebhook()
	errStorage := b.settingsStorage.Close()
	// TODO: See this issue: https://github.com/tucnak/telebot/issues/584
	b.bot.Stop()
	log.Infoln("Bot has been stopped")
	return errors.Join(err, errWebhook, errStorage)
}
