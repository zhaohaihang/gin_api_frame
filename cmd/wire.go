// +build wireinject

package main

import (
	"gin_api_frame/internal/app"
	"gin_api_frame/config"
	v1 "gin_api_frame/internal/api/v1"
	"gin_api_frame/internal/cron"
	"gin_api_frame/internal/dao"
	"gin_api_frame/internal/http"
	"gin_api_frame/internal/routes"
	"gin_api_frame/internal/service"
	"gin_api_frame/pkg/database"
	"gin_api_frame/pkg/logger"
	"gin_api_frame/pkg/mail"
	"gin_api_frame/pkg/redis"
	"gin_api_frame/pkg/storages/qiniu"

	"github.com/google/wire"
)

var providerSet = wire.NewSet(
	app.AppProviderSet,
	http.HttpServerProviderSet,
	config.ConfigProviderSet,
	routes.RouterProviderSet,
	v1.ControllerProviderSet,
	service.ServiceProviderSet,
	database.DatabaseProviderSet,
	dao.DaoProviderSet,
	logger.LoggerProviderSet,
	redis.RedisPoolProviderSet,
	qiniu.QiNiuStroageProviderSet,
	cron.CronServerProviderSet,
	mail.MailPoolProviderSet,
)

func CreateApp() (*app.App, error) {
	panic(wire.Build(providerSet))
}

