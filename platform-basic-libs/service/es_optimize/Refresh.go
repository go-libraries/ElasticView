package es_optimize

import (
	"context"

	"github.com/olivere/elastic"
)

type Refresh struct {
	indexName []string
}

func (this *Refresh) SetIndexName(indexName string) {
	this.indexName = append(this.indexName, indexName)
}

func (this Refresh) Do(client *elastic.Client) (err error) {
	_, err = client.Refresh(this.indexName...).Do(context.Background())
	return
}

func newRefresh() OptimizeInterface {
	return &Refresh{}
}
