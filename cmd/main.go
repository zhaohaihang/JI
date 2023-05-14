package main

import (
	"ji/config"
	"ji/loading"
	"ji/routes"
)

// @title JI API
// @version 1.0
// @description The api docs of JI project
// @termsOfService http://swagger.io/terms/

// @contact.name Zhao Haihang
// @contact.url http://www.swagger.io/support
// @contact.email 1932859223@qq.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:4000
// @BasePath /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	config.LoadConfig()
	loading.Init()
	r := routes.NewRouter()
	_ = r.Run(config.Conf.Server.ServerPort)
}
