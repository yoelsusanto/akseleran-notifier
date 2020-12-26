package jobs

import (
	"github.com/go-redis/redis/v8"
	"github.com/yoelsusanto/akseleran-notifier/config"
	"github.com/yoelsusanto/akseleran-notifier/discord"
	logwrapper "github.com/yoelsusanto/akseleran-notifier/log"
)

type JobDependencies struct {
	DiscordModule *discord.Module
	RedisClient   *redis.Client
	CronConfigs   map[string]*config.CronConfig
	Log           *logwrapper.StandardLogger
}
