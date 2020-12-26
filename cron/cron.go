package cron

import (
	"github.com/go-redis/redis/v8"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"github.com/yoelsusanto/akseleran-notifier/config"
	"github.com/yoelsusanto/akseleran-notifier/cron/jobs"
	"github.com/yoelsusanto/akseleran-notifier/discord"
	logwrapper "github.com/yoelsusanto/akseleran-notifier/log"
)

const notifyAkseleran string = "akseleran_notify"

type Options struct {
	DiscordModule *discord.Module
	RedisClient   *redis.Client
	CronConfigs   map[string]*config.CronConfig
	Log           *logwrapper.StandardLogger
}

type Module struct {
	Log  *logwrapper.StandardLogger
	Cron *cron.Cron
}

func CreateCronModule(options *Options) (*Module, error) {
	cronInstance := cron.New()
	module := &Module{Log: options.Log, Cron: cronInstance}

	jobDependencies := &jobs.JobDependencies{
		DiscordModule: options.DiscordModule,
		RedisClient:   options.RedisClient,
		CronConfigs:   options.CronConfigs,
		Log:           options.Log,
	}

	_, err := cronInstance.AddFunc(options.CronConfigs[notifyAkseleran].Rule, jobs.CreateNewAkseleranCampaignHandler(jobDependencies))
	if err != nil {
		module.Log.WithFields(logrus.Fields{
			"error":        err.Error(),
			"cron_handler": "CreateNewAkseleranCampaignHandler",
		}).Error("[CreateCronModule] Failed to add handler to cron instance")
		return nil, err
	}

	return module, nil
}

func (m Module) Start() {
	m.Cron.Start()
}

func (m Module) Stop() {
	m.Cron.Stop()
}
