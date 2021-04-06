package es

import (
	"sync"
)

type EsCache struct {
	esConnectMap map[int]EsClient
}

var (
	once    sync.Once
	esCache *EsCache
)

func NewEsCache() *EsCache {
	once.Do(func() {
		esCache = &EsCache{esConnectMap: map[int]EsClient{}}
	})
	return esCache
}

func (this *EsCache) Set(id int, esClient EsClient) {

	this.esConnectMap[id] = esClient
}

func (this *EsCache) Get(id int) EsClient {

	if _, getConnect := this.esConnectMap[id]; getConnect {

		return this.esConnectMap[id]
	}
	return nil
}

func (this *EsCache) Rem(id int) {

	if _, getConnect := this.esConnectMap[id]; getConnect {

		delete(this.esConnectMap, id)
	}
}
