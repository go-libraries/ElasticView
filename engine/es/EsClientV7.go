package es

import (
	"context"

	"ElasticView/platform-basic-libs/request"

	elasticV7 "github.com/olivere/elastic/v7"
)

type EsClientV7 struct {
	client          *elasticV7.Client
	esConnectConfig request.EsConnect
}

func NewEsClientV7(esConnectConfig request.EsConnect) (client EsClient, err error) {

	optList := []elasticV7.ClientOptionFunc{elasticV7.SetSniff(false)}

	optList = append(optList, elasticV7.SetURL(esConnectConfig.Ip))

	if esConnectConfig.User != "" || esConnectConfig.Pwd != "" {
		optList = append(optList, elasticV7.SetBasicAuth(esConnectConfig.User, esConnectConfig.Pwd))
	}

	esClient, err := elasticV7.NewClient(optList...)

	if err != nil {
		return
	}

	client = &EsClientV7{
		client:          esClient,
		esConnectConfig: esConnectConfig,
	}

	return
}

func (this *EsClientV7) Ping() (interface{}, int, error) {
	return this.client.Ping(this.esConnectConfig.Ip).Do(context.Background())
}

func (this *EsClientV7) Index(indexName string) (interface{}, error) {
	return this.client.CreateIndex(indexName).Do(context.Background())
}
