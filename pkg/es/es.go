package es

import (
	"context"
	"ji/config"

	"github.com/google/wire"
	"github.com/olivere/elastic/v7"
)

type EsClient struct {
	client *elastic.Client
}

func NewEsClient(config *config.Config) (*EsClient, error) {
	var esClient EsClient
	esConn := "http://" + config.Es.EsHost + ":" + config.Es.Esport
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(esConn))
	if err != nil {
		return nil, err
	}
	esClient.client = client
	return &esClient, nil
}

var EsClientProviderSet = wire.NewSet(NewEsClient)

func (cli *EsClient) Create(Params map[string]string) (string, error) {
	var (
		res *elastic.IndexResponse
		err error
	)
	 
	cli.client.Index()
	res, err = cli.client.Index().
		Index(Params["index"]).
		Id(Params["id"]).BodyJson(Params["bodyJson"]).
		Do(context.Background())

	if err != nil {
		return "", err
	}
	return res.Result, nil
}

func (cli *EsClient) Update(Params map[string]string, Doc map[string]string) string {
	var (
		res *elastic.UpdateResponse
		err error
	)
	res, err = cli.client.Update().
		Index(Params["index"]).
		Id(Params["id"]).
		Doc(Doc).
		Do(context.Background())

	if err != nil {
		return ""
	}
	return res.Result
}

func (cli *EsClient) Delete(Params map[string]string) (string, error) {
	var (
		res *elastic.DeleteResponse
		err error
	)

	res, err = cli.client.Delete().
		Index(Params["index"]).
		Id(Params["id"]).
		Do(context.Background())

	if err != nil {
		return "", err
	}

	return res.Result, nil
}