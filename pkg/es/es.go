package es

import (
	"ji/config"

	"github.com/google/wire"
	"github.com/olivere/elastic/v7"
)

type EsClient struct {
	client *elastic.Client
}

func NewEsClient(config *config.Config) (*EsClient, error) {
	var esClient *EsClient
	esConn := "http://" + config.Es.EsHost + ":" + config.Es.Esport
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(esConn))
	if err != nil {
		return nil, err
	}
	esClient.client = client
	return esClient, nil
}

var EsClientProviderSet = wire.NewSet(NewEsClient)
