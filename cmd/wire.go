// +build wireinject

package main

import (
	"ji/config"
	v1 "ji/internal/api/v1"
	"ji/internal/dao"
	"ji/internal/routes"
	"ji/internal/service"
	"ji/pkg/database"
	"ji/pkg/logger"
	"ji/pkg/redis"
	"ji/pkg/server"
	"ji/pkg/storages/localstroage"
	"ji/pkg/storages/qiniu"

	"github.com/google/wire"
)

var providerSet = wire.NewSet(
	server.ServerProviderSet,
	config.ConfigProviderSet,
	routes.RouterProviderSet,
	v1.ActivityControllerProviderSet,
	v1.UserControllerProviderSet,
	service.UserServiceProviderSet,
	service.ActivityServiceProviderSet,
	database.DatabaseProviderSet,
	dao.UserDaoProviderSet,
	dao.ActivityDaoProviderSet,
	logger.LoggerProviderSet,
	localstroage.LocalStroageProviderSet,
	redis.RedisPoolProviderSet,
	qiniu.QiNiuStroageProviderSet,

)

func CreateServer() (*server.Server, error) {
	panic(wire.Build(providerSet))
}
