package es_optimize

import (
	"context"

	"github.com/olivere/elastic"
)

type Forcemerge struct {
	indexName []string
}

func (this *Forcemerge) SetIndexName(indexName string) {
	this.indexName = append(this.indexName, indexName)
}

func (this *Forcemerge) Do(client *elastic.Client) (err error) {
	_, err = client.Forcemerge(this.indexName...).MaxNumSegments(1).Do(context.Background())
	return
}

func newForcemerge() OptimizeInterface {
	return &Forcemerge{}
}
