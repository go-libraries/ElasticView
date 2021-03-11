package es

import (
	"context"

	"ElasticView/platform-basic-libs/request"

	elasticV6 "github.com/olivere/elastic"
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

func (this *EsClientV7) CreateIndex(indexName string, body interface{}) (interface{}, error) {
	return this.client.CreateIndex(indexName).BodyJson(body).Do(context.Background())
}

func (this *EsClientV7) CatIndices() (interface{}, error) {
	return this.client.CatIndices().Pretty(true).Do(context.Background())
}

func (this *EsClientV7) CatAliases() (interface{}, error) {
	return this.client.CatAliases().Pretty(true).Do(context.Background())
}

func (this *EsClientV7) CatAllocation() (interface{}, error) {
	return this.client.CatAllocation().Pretty(true).Do(context.Background())
}

func (this *EsClientV7) CatCount() (interface{}, error) {
	return this.client.CatCount().Pretty(true).Do(context.Background())
}

func (this *EsClientV7) CatHealth() (interface{}, error) {
	return this.client.CatHealth().Pretty(true).Do(context.Background())
}

func (this *EsClientV7) CatShards() (interface{}, error) {
	return this.client.CatShards().Pretty(true).Do(context.Background())
}

func (this *EsClientV7) DeleteIndex(indexNameList []string) (interface{}, error) {
	return this.client.DeleteIndex(indexNameList...).Do(context.Background())
}

func (this *EsClientV7) CloseIndex(indexNameList []string) error {
	var err error
	for _, indexName := range indexNameList {
		_, err = this.client.CloseIndex(indexName).Do(context.Background())
		if err != nil {
			return err
		}
	}
	return nil
}

func (this *EsClientV7) OpenIndex(indexNameList []string) error {
	var err error
	for _, indexName := range indexNameList {
		_, err = this.client.OpenIndex(indexName).Do(context.Background())
		if err != nil {
			return err
		}
	}
	return nil
}

func (this *EsClientV7) FreezeIndex(indexNameList []string) error {
	var err error
	for _, indexName := range indexNameList {
		_, err = this.client.FreezeIndex(indexName).Do(context.Background())
		if err != nil {
			return err
		}
	}
	return nil
}

func (this *EsClientV7) UnfreezeIndex(indexNameList []string) error {
	var err error
	for _, indexName := range indexNameList {
		_, err = this.client.UnfreezeIndex(indexName).Do(context.Background())
		if err != nil {
			return err
		}
	}
	return nil
}

func (this *EsClientV7) CreateMapping(indexName string, body Json) (interface{}, error) {
	return this.client.PutMapping().
		Index(indexName).
		BodyJson(body).
		Do(context.Background())
}

func (this *EsClientV7) IndexPutSettings(indexName string, body Json) (interface{}, error) {
	return this.client.IndexPutSettings().Index(indexName).BodyJson(body).Do(context.TODO())
}

func (this *EsClientV7) Reindex(sourceIndex, destinationIndex string) (interface{}, error) {
	return this.client.Reindex().SourceIndex(sourceIndex).DestinationIndex(destinationIndex).Do(context.Background())
}

func (this *EsClientV7) TasksList() (interface{}, error) {
	return this.client.TasksList().Pretty(true).
		Human(true).
		//Header("X-Opaque-Id", "123456").
		Do(context.TODO())
}

func (this *EsClientV7) Refresh(indexName ...string) (interface{}, error) {
	return this.client.Refresh(indexName...).Do(context.TODO())
}

func (this *EsClientV7) Flush(indexName ...string) (interface{}, error) {
	return this.client.Flush(indexName...).Do(context.TODO())
}

func (this *EsClientV7) Fsync(indexName ...string) (interface{}, error) {
	return this.client.SyncedFlush(indexName...).Do(context.TODO())
}

func (this *EsClientV7) Rollover(alias, indexName string) (interface{}, error) {
	return this.client.RolloverIndex(alias).NewIndex(indexName).Do(context.TODO())
}

func (this *EsClientV7) IndexSegments(indexName string) (interface{}, error) {
	return this.client.IndexSegments(indexName).Pretty(true).Do(context.TODO())
}

func (this *EsClientV7) Alias(alias, indexName string) (interface{}, error) {
	return this.client.Aliases().Index(indexName).Alias(alias).Do(context.Background())
}

func (this *EsClientV7) IndexStats(indexName []string, metrics []string) (interface{}, error) {
	return this.client.IndexStats().Index(indexName...).Metric(metrics...).Do(context.Background())
}

func (this *EsClientV7) GetMapping(indexName string, body Json, typeName ...string) (interface{}, error) {
	return this.client.GetMapping().Index(indexName).Type("_doc").Do(context.Background())
}

func (this *EsClientV7) PutData(indexName string, body Json, typeName ...string) (interface{}, error) {
	return this.client.Index().Index(indexName).BodyJson(body).Do(context.Background())
}

func (this *EsClientV7) DeleteById(indexName, id string, typeName ...string) (interface{}, error) {
	return this.client.Delete().Index(indexName).Id(id).Do(context.Background())
}

func (this *EsClientV7) Search(indexName string, query elasticV6.Query, sort *Sort, page *Page, fields []string, isInclude bool, typeName ...string) (*elasticV6.SearchResult, error) {

	search := this.client.Search(indexName).Query(query)
	if sort != nil {
		search = search.Sort(sort.Field, sort.Ascending)
	}
	if page != nil {
		search = search.From(page.PageNum).Size(page.PageSize)
	}

	if len(fields) > 0 {
		fsc := elasticV7.NewFetchSourceContext(true)
		if isInclude {
			search.FetchSourceContext(fsc.Include(fields...))
		} else {
			search.FetchSourceContext(fsc.Exclude(fields...))
		}
	}

	var T interface{}

	T, err := search.Do(context.Background())
	return T.(*elasticV6.SearchResult), err
}

func (this *EsClientV7) Count(indexName string, query elasticV6.Query, typeName ...string) (int64, error) {
	return this.client.Count(indexName).Query(query).Do(context.Background())
}

func (this *EsClientV7) DeleteByQuery(indexName string, query elasticV6.Query, typeName ...string) (interface{}, error) {
	return this.client.DeleteByQuery(indexName).Query(query).Pretty(true).Do(context.Background())
}

func (this *EsClientV7) UpdateByID(indexName string, id string, query elasticV6.Query, typeName ...string) (interface{}, error) {
	return this.client.Update().Index(indexName).Id(id).Doc(query).Do(context.Background())
}
