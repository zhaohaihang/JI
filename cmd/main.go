package main

import (
	"ji/config"
	"ji/loading"
	"ji/routes"
)

func main() {
	config.LoadConfig()
	loading.Init()
	r := routes.NewRouter()
	_ = r.Run(config.Conf.Server.ServerPort)
}
