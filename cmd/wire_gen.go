// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/google/wire"
	"ji/config"
	"ji/internal/api/v1"
	"ji/internal/dao"
	"ji/internal/routes"
	"ji/internal/service"
	"ji/pkg/database"
	"ji/pkg/logger"
	"ji/pkg/redis"
	"ji/pkg/server"
	"ji/pkg/storages/localstroage"
	"ji/pkg/storages/qiniu"
)

// Injectors from wire.go:

func CreateServer() (*server.Server, error) {
	configConfig := config.NewConfig()
	loggerLogger := logger.NewLogger()
	databaseDatabase := database.NewDatabase(configConfig)
	userDao := dao.NewUserDao(databaseDatabase)
	activityDao := dao.NewActivityDao(databaseDatabase)
	pool, err := redis.NewRedisPool(configConfig)
	if err != nil {
		return nil, err
	}
	qiNiuStroage := qiniu.NewQiNiuStroage(configConfig)
	activityService := service.NewActivityService(userDao, activityDao, pool, qiNiuStroage)
	userService := service.NewUserService(userDao, activityDao, qiNiuStroage)
	activityController := v1.NewActivityContrller(loggerLogger, activityService, userService)
	userController := v1.NewUserContrller(loggerLogger, activityService, userService)
	engine := routes.NewRouter(activityController, userController)
	serverServer := server.NewServer(configConfig, engine)
	return serverServer, nil
}

// wire.go:

var providerSet = wire.NewSet(server.ServerProviderSet, config.ConfigProviderSet, routes.RouterProviderSet, v1.ActivityControllerProviderSet, v1.UserControllerProviderSet, service.UserServiceProviderSet, service.ActivityServiceProviderSet, database.DatabaseProviderSet, dao.UserDaoProviderSet, dao.ActivityDaoProviderSet, logger.LoggerProviderSet, localstroage.LocalStroageProviderSet, redis.RedisPoolProviderSet, qiniu.QiNiuStroageProviderSet)
