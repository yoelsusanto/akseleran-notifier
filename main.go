package main

import (
	defaultlog "log"
	"os"
	"os/signal"

	"github.com/sirupsen/logrus"
	logwrapper "github.com/yoelsusanto/akseleran-notifier/log"

	"github.com/yoelsusanto/akseleran-notifier/config"
	"github.com/yoelsusanto/akseleran-notifier/cron"
	"github.com/yoelsusanto/akseleran-notifier/discord"
	"github.com/yoelsusanto/akseleran-notifier/redis"
)

func main() {
	cfg, err := config.ReadConfig()
	if err != nil {
		defaultlog.Fatalf("Failed to read config: %+v", err)
	}

	log, err := logwrapper.CreateLogger(&logwrapper.Options{ServiceName: cfg.General.ServiceName, Environment: cfg.General.Environment, LogPath: cfg.General.LogPath})
	if err != nil {
		defaultlog.Fatalf("Failed to initiate logger: %+v", err)
	}

	redisClient := redis.CreateRedisClient(cfg.RedisConfig)

	discordModule, err := discord.CreateDiscordModule(&discord.Options{Cfg: cfg.DiscordConfig, Log: log})
	if err != nil {
		log.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Fatal("[Main] Failed to create discord module")
	}

	err = discordModule.Client.Open()
	if err != nil {
		log.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Fatal("[Main] Discord client failed to open connection")
	}

	cronModule, err := cron.CreateCronModule(&cron.Options{RedisClient: redisClient, DiscordModule: discordModule, CronConfigs: cfg.CronConfigs, Log: log})
	if err != nil {
		log.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Fatal("[Main] Failed to create cron module")
	}
	cronModule.Start()

	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig
}
