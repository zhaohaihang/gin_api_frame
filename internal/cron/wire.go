package cron

import "github.com/google/wire"

var CronServerProviderSet = wire.NewSet(
	NewCronServer,
	NewTasks,
)
