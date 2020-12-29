package jobs

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"github.com/yoelsusanto/akseleran-notifier/internal"
)

func CreateNewAkseleranCampaignHandler(deps *JobDependencies) func() {
	return func() {
		deps.Log.Info("Start check akseleran campaign")

		redisClient := deps.RedisClient
		discordModule := deps.DiscordModule

		resp, err := internal.GetAkseleranCampaign()
		if err != nil {
			deps.Log.WithFields(logrus.Fields{
				"error": err.Error(),
			}).Error("Start check akseleran campaign")
			return
		}

		ctx := context.Background()

		campaigns := resp.Data
		for _, campaign := range campaigns {
			campaignID := strconv.Itoa(campaign.ID)
			// Check if id already existed in db
			_, err := redisClient.Get(ctx, campaignID).Result()
			dataExisted := err == nil

			if dataExisted {
				continue
			} else if err != nil && err != redis.Nil {
				deps.Log.WithFields(logrus.Fields{
					"error":    err.Error(),
					"data_key": campaignID,
				}).Error("Error when getting data from redis")
				continue
			}

			_, err = redisClient.Set(ctx, campaignID, campaign, 0).Result()
			if err != nil {
				deps.Log.WithFields(logrus.Fields{
					"error": err.Error(),
					"data":  campaign,
				}).Error("Error when saving data to redis")
			}

			message, err := internal.CampaignToMessage(campaign)

			if err != nil {
				log.Println(err)
			}

			messageInBlock := fmt.Sprintf("```%s```", message)
			messageWithMention := discordModule.AddMentionEveryOne(messageInBlock)

			err = discordModule.SendMessageToAkseleranChannel(messageWithMention)
			if err != nil {
				log.Println(err)
			}
		}
	}
}
