package es

import (
	"ji/config"

	"github.com/google/wire"
	"github.com/olivere/elastic/v7"
)

var EsClient *elastic.Client

func NewEsClient(config *config.Config) (*elastic.Client, error) {
	esConn := "http://" + config.Es.EsHost + ":" + config.Es.Esport
	esClient, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(esConn))
	if err != nil {
		return nil,err
	}
	EsClient = esClient
	return esClient,nil
}

var EsClientProviderSet = wire.NewSet(NewEsClient)
