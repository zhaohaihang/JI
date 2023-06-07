// +build wireinject

package main

import (
	"ji/app"
	"ji/config"
	v1 "ji/internal/api/v1"
	"ji/internal/cron"
	"ji/internal/backproc"
	"ji/internal/dao"
	"ji/internal/http"
	"ji/internal/routes"
	"ji/internal/service"
	"ji/pkg/database"
	"ji/pkg/es"
	"ji/pkg/mq"
	"ji/pkg/mail"
	"ji/pkg/logger"
	"ji/pkg/redis"
	"ji/pkg/storages/localstroage"
	"ji/pkg/storages/qiniu"

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
	localstroage.LocalStroageProviderSet,
	redis.RedisPoolProviderSet,
	qiniu.QiNiuStroageProviderSet,
	cron.CronServerProviderSet,
	backproc.BackProcServerProviderSet,
	mail.MailPoolProviderSet,
	es.EsClientProviderSet,
	mq.RabbitMQClientProviderSet,
)

func CreateApp() (*app.App, error) {
	panic(wire.Build(providerSet))
}

