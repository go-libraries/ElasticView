package es_optimize

import (
	"context"
	"log"

	"github.com/olivere/elastic"
)

type Flush struct {
	indexName []string
}

func (this *Flush) SetIndexName(indexName string) {
	this.indexName = append(this.indexName, indexName)
}

func (this Flush) Do(client *elastic.Client) (err error) {
	if len(this.indexName) == 0 {
		_, err = client.Flush().Do(context.Background())
		return
	}
	log.Println("this.indexName", this.indexName)
	_, err = client.Flush(this.indexName...).Do(context.Background())
	return
}

func newFlush() OptimizeInterface {
	return &Flush{}
}
