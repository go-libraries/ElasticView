package es

import (
	"context"

	"ElasticView/platform-basic-libs/request"

	elasticV6 "github.com/olivere/elastic"
)

type EsClientV6 struct {
	client          *elasticV6.Client
	esConnectConfig request.EsConnect
}

func NewEsClientV6(esConnectConfig request.EsConnect) (client EsClient, err error) {

	optList := []elasticV6.ClientOptionFunc{elasticV6.SetSniff(false)}

	optList = append(optList, elasticV6.SetURL(esConnectConfig.Ip))

	if esConnectConfig.User != "" || esConnectConfig.Pwd != "" {
		optList = append(optList, elasticV6.SetBasicAuth(esConnectConfig.User, esConnectConfig.Pwd))
	}

	esClient, err := elasticV6.NewClient(optList...)

	if err != nil {
		return
	}

	client = &EsClientV6{
		client:          esClient,
		esConnectConfig: esConnectConfig,
	}

	return
}

func (this *EsClientV6) Ping() (interface{}, int, error) {
	return this.client.Ping(this.esConnectConfig.Ip).Do(context.Background())
}

func (this *EsClientV6) Index(indexName string) (interface{}, error) {
	return this.client.CreateIndex(indexName).Do(context.Background())
}
