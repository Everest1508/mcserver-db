package services

import (
	"github.com/everest1508/mcserver-db/constants"
	utils "github.com/everest1508/mcserver-db/utils/api"
	"github.com/robfig/cron"
)

type CronJob struct {
	Client *utils.APIClient
}

func NewCronJob(client *utils.APIClient) *CronJob {
	cron := &CronJob{Client: client}
	return cron
}

func (cj CronJob) StartCron() {
	c := cron.New()

	c.AddFunc(constants.CRON_STRING, func() { FetchAndStoreJarData(cj.Client) })
	c.Start()
}
