package cron

import (
	"github.com/robfig/cron"
	"github.com/sirupsen/logrus"
)

type CronServer struct {
	cron  *cron.Cron
	tasks *Tasks
}

func NewCronServer(t *Tasks) *CronServer {
	return &CronServer{
		cron:  cron.New(),
		tasks: t,
	}
}

func (cs *CronServer) Start() {
	addCronFunc(cs.cron, "@every 1m", func() {
		cs.tasks.Task1()
	})

	addCronFunc(cs.cron, "@every 2m", func() {
		cs.tasks.Task2()
	})

	cs.cron.Start()
}

func (cs *CronServer) Stop() {
	cs.cron.Stop()
}

func addCronFunc(c *cron.Cron, sepc string, cmd func()) {
	err := c.AddFunc(sepc, cmd)
	if err != nil {
		logrus.Error(err)
	}
}
