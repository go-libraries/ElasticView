package es

import (
	"context"
	"errors"

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

func (this *EsClientV6) CreateIndex(indexName string, body interface{}) (interface{}, error) {
	return this.client.CreateIndex(indexName).BodyJson(body).Do(context.Background())
}

func (this *EsClientV6) CatIndices() (interface{}, error) {
	return this.client.CatIndices().Pretty(true).Do(context.Background())
}

func (this *EsClientV6) CatAliases() (interface{}, error) {
	return this.client.CatAliases().Pretty(true).Do(context.Background())
}

func (this *EsClientV6) CatAllocation() (interface{}, error) {
	return this.client.CatAllocation().Pretty(true).Do(context.Background())
}

func (this *EsClientV6) CatCount() (interface{}, error) {
	return this.client.CatCount().Pretty(true).Do(context.Background())
}

func (this *EsClientV6) CatHealth() (interface{}, error) {
	return this.client.CatHealth().Pretty(true).Do(context.Background())
}

func (this *EsClientV6) CatShards() (interface{}, error) {
	return this.client.CatShards().Pretty(true).Do(context.Background())
}

func (this *EsClientV6) DeleteIndex(indexNameList []string) (interface{}, error) {
	return this.client.DeleteIndex(indexNameList...).Do(context.Background())
}

func (this *EsClientV6) CloseIndex(indexNameList []string) error {
	var err error
	for _, indexName := range indexNameList {
		_, err = this.client.CloseIndex(indexName).Do(context.Background())
		if err != nil {
			return err
		}
	}
	return nil
}

func (this *EsClientV6) OpenIndex(indexNameList []string) error {
	var err error
	for _, indexName := range indexNameList {
		_, err = this.client.OpenIndex(indexName).Do(context.Background())
		if err != nil {
			return err
		}
	}
	return nil
}

func (this *EsClientV6) FreezeIndex(indexNameList []string) error {
	return nil
}

func (this *EsClientV6) UnfreezeIndex(indexNameList []string) error {
	return nil
}

func (this *EsClientV6) CreateMapping(indexName string, body Json) (interface{}, error) {
	return this.client.PutMapping().
		Index(indexName).
		BodyJson(body).
		IncludeTypeName(true).
		Do(context.Background())
}

func (this *EsClientV6) IndexPutSettings(indexName string, body Json) (interface{}, error) {
	return this.client.IndexPutSettings().Index(indexName).BodyJson(body).Do(context.TODO())
}

func (this *EsClientV6) Refresh(indexName ...string) (interface{}, error) {
	return this.client.Refresh(indexName...).Do(context.TODO())
}

func (this *EsClientV6) Flush(indexName ...string) (interface{}, error) {
	return this.client.Flush(indexName...).Do(context.TODO())
}

func (this *EsClientV6) Fsync(indexName ...string) (interface{}, error) {
	return this.client.SyncedFlush(indexName...).Do(context.TODO())
}
func (this *EsClientV6) Rollover(alias, indexName string) (interface{}, error) {
	return this.client.RolloverIndex(alias).NewIndex(indexName).Do(context.TODO())
}

func (this *EsClientV6) IndexSegments(indexName string) (interface{}, error) {
	return this.client.IndexSegments(indexName).Pretty(true).Do(context.TODO())
}

func (this *EsClientV6) Alias(alias, indexName string) (interface{}, error) {
	return this.client.Aliases().Index(indexName).Alias(alias).Do(context.Background())
}

func (this *EsClientV6) Reindex(sourceIndex, destinationIndex string) (interface{}, error) {
	return this.client.Reindex().SourceIndex(sourceIndex).DestinationIndex(destinationIndex).Do(context.Background())
}

func (this *EsClientV6) TasksList() (interface{}, error) {
	return this.client.TasksList().Pretty(true).
		Human(true).
		//Header("X-Opaque-Id", "123456").
		Do(context.TODO())
}

func (this *EsClientV6) IndexStats(indexName []string, metrics []string) (interface{}, error) {
	return this.client.IndexStats().Index(indexName...).Metric(metrics...).Do(context.Background())
}

func (this *EsClientV6) GetMapping(indexName string, body Json, typeName ...string) (interface{}, error) {
	return this.client.GetMapping().Index(indexName).Type(typeName...).Do(context.Background())
}

func (this *EsClientV6) PutData(indexName string, body Json, typeName ...string) (interface{}, error) {
	if len(typeName) == 0 {
		return nil, errors.New("Type 不能为空!")
	}
	return this.client.Index().Index(indexName).Type(typeName[0]).BodyJson(body).Do(context.Background())
}

func (this *EsClientV6) DeleteById(indexName, id string, typeName ...string) (interface{}, error) {
	if len(typeName) == 0 {
		return nil, errors.New("Type 不能为空!")
	}
	return this.client.Delete().Index(indexName).Type(typeName[0]).Id(id).Do(context.Background())
}

func (this *EsClientV6) Search(indexName string, query elasticV6.Query, sort *Sort, page *Page, fields []string, isInclude bool, typeName ...string) (*elasticV6.SearchResult, error) {

	search := this.client.Search(indexName).Query(query)
	if sort != nil {
		search = search.Sort(sort.Field, sort.Ascending)
	}
	if page != nil {
		search = search.From(page.PageNum).Size(page.PageSize)
	}

	if len(fields) > 0 {
		fsc := elasticV6.NewFetchSourceContext(true)
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

func (this *EsClientV6) Count(indexName string, query elasticV6.Query, typeName ...string) (int64, error) {
	if len(typeName) == 0 {
		return 0, errors.New("Type 不能为空!")
	}
	return this.client.Count(indexName).Type(typeName[0]).Query(query).Do(context.Background())
}

func (this *EsClientV6) DeleteByQuery(indexName string, query elasticV6.Query, typeName ...string) (interface{}, error) {
	if len(typeName) == 0 {
		return nil, errors.New("Type 不能为空!")
	}
	return this.client.DeleteByQuery(indexName).Type(typeName[0]).Query(query).Pretty(true).Do(context.Background())
}

func (this *EsClientV6) UpdateByID(indexName string, id string, query elasticV6.Query, typeName ...string) (interface{}, error) {
	if len(typeName) == 0 {
		return nil, errors.New("Type 不能为空!")
	}
	return this.client.Update().Index(indexName).Type(typeName[0]).Id(id).Doc(query).Do(context.Background())
}
