package config

import (
	"github.com/spf13/viper"
)

type CronConfig struct {
	Rule string `mapstructure:"rule"`
}

type DiscordConfig struct {
	Token     string `mapstructure:"token"`
	ChannelID string `mapstructure:"channel_id"`
}

type RedisConfig struct {
	Address  string `mapstructure:"address"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db_number"`
}

type ObsScene struct {
	Name     string `mapstructure:"name"`
	Image    string `mapstructure:"image"`
	ButtonId int
}

type GeneralConfig struct {
	ServiceName string `mapstructure:"service_name"`
	Environment string `mapstructure:"environment"`
	LogPath 	string `mapstructure:"log_path"`
}

type Config struct {
	General       *GeneralConfig         `mapstructure:"general"`
	RedisConfig   *RedisConfig           `mapstructure:"redis"`
	DiscordConfig *DiscordConfig         `mapstructure:"discord"`
	CronConfigs   map[string]*CronConfig `mapstructure:"crons"`
}

func ReadConfig() (*Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
