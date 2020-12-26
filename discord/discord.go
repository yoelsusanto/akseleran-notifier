package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
	"github.com/yoelsusanto/akseleran-notifier/config"
	logwrapper "github.com/yoelsusanto/akseleran-notifier/log"
)

type Options struct {
	Cfg *config.DiscordConfig
	Log *logwrapper.StandardLogger
}

type Module struct {
	Cfg    *config.DiscordConfig
	Client *discordgo.Session
	Log    *logwrapper.StandardLogger
}

func CreateDiscordModule(options *Options) (*Module, error) {
	module := &Module{Log: options.Log, Cfg: options.Cfg}

	client, err := discordgo.New("Bot " + module.Cfg.Token)
	if err != nil {
		module.Log.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("[CreateDiscordModule] Failed to create discord module")
		return nil, err
	}

	module.Client = client
	return module, nil
}

func (m *Module) SendMessageToAkseleranChannel(message string) error {
	_, err := m.Client.ChannelMessageSend(m.Cfg.ChannelID, message)
	if err != nil {
		m.Log.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("[SendMessageToAkseleranChannel] Failed to send message to akseleran channel")
		return err
	}
	return nil
}
