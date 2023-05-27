// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/google/wire"
	"ji/app"
	"ji/config"
	"ji/internal/api/v1"
	"ji/internal/cron"
	"ji/internal/dao"
	"ji/internal/http"
	"ji/internal/routes"
	"ji/internal/service"
	"ji/pkg/database"
	"ji/pkg/logger"
	"ji/pkg/redis"
	"ji/pkg/storages/localstroage"
	"ji/pkg/storages/qiniu"
)

// Injectors from wire.go:

func CreateApp() (*app.App, error) {
	configConfig := config.NewConfig()
	logrusLogger := logger.NewLogger()
	databaseDatabase := database.NewDatabase(configConfig)
	userDao := dao.NewUserDao(databaseDatabase)
	activityDao := dao.NewActivityDao(databaseDatabase)
	pool, err := redis.NewRedisPool(configConfig)
	if err != nil {
		return nil, err
	}
	qiNiuStroage := qiniu.NewQiNiuStroage(configConfig)
	activityService := service.NewActivityService(logrusLogger, userDao, activityDao, pool, qiNiuStroage)
	userService := service.NewUserService(logrusLogger, userDao, activityDao, qiNiuStroage)
	activityController := v1.NewActivityContrller(logrusLogger, activityService, userService)
	userController := v1.NewUserContrller(logrusLogger, activityService, userService)
	engine := routes.NewRouter(activityController, userController)
	httpServer := http.NewHttpServer(configConfig, engine)
	tasks := cron.NewTasks(userDao, activityDao)
	cronServer := cron.NewCronServer(tasks)
	appApp := app.NewApp(configConfig, engine, httpServer, cronServer)
	return appApp, nil
}

// wire.go:

var providerSet = wire.NewSet(app.AppProviderSet, http.HttpServerProviderSet, config.ConfigProviderSet, routes.RouterProviderSet, v1.ControllerProviderSet, service.ServiceProviderSet, database.DatabaseProviderSet, dao.DaoProviderSet, logger.LoggerProviderSet, localstroage.LocalStroageProviderSet, redis.RedisPoolProviderSet, qiniu.QiNiuStroageProviderSet, cron.CronServerProviderSet)
