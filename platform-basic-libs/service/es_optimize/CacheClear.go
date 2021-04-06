package es_optimize

import (
	"context"

	"github.com/olivere/elastic"
)

type CacheClear struct {
	indexName []string
}

func (this *CacheClear) SetIndexName(indexName string) {
	this.indexName = append(this.indexName, indexName)
}

func (this CacheClear) Do(client *elastic.Client) (err error) {
	_, err = client.ClearCache(this.indexName...).Do(context.Background())
	return
}

func newCacheClear() OptimizeInterface {
	return &CacheClear{}
}
