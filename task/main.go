package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/1340691923/ElasticView/engine/es"
	_ "github.com/go-sql-driver/mysql"
	"github.com/olivere/elastic"
)

var (
	appName        string
	configFileDir  string
	configFileName string
	configFileExt  string
)

func init() {
	flag.StringVar(&appName, "appName", "ReportPlatform", "应用名")
	flag.StringVar(&configFileDir, "configFileDir", "config", "配置文件夹名")
	flag.StringVar(&configFileName, "configFileName", "config", "配置文件名")
	flag.StringVar(&configFileExt, "configFileExt", "json", "配置文件后缀")
}

// By 肖文龙
func main() {
	/*app := application.NewApp(
		application.WithAppName(appName),
		application.WithConfigFileDir(configFileDir),
		application.WithConfigFileName(configFileName),
		application.WithConfigFileExt(configFileExt),
		application.RegisterInitFnObserver(application.InitLogs),
	)

	err := app.InitConfig().NotifyInitFnObservers().Error()

	if err != nil {
		logs.Logger.Error("ElasticView 初始化失败", zap.String("err.Error()", err.Error()))
		panic(err)
	}*/

	foEs, err := es.NewEsClientV6(es.EsConnect{
		Ip:      "http://elasticsearch.lynlzqy.com:9200",
		User:    "elastic",
		Pwd:     "lyn1314",
		Version: 0,
	})
	e(err)

	/*toEs, err := es.NewEsClientV6(es.EsConnect{
		Ip:      "http://127.0.0.1:9200",
		Version: 0,
	})
	e(err)*/
	query := elastic.NewBoolQuery().
		Must(elastic.NewMatchPhraseQuery("game", 7)).
		Filter(elastic.NewRangeQuery("create_time").
			From("2021-06-01 00:00:00").
			To("2021-07-13 00:00:00"))
	foEsRes, err := foEs.(*es.EsClientV6).
		Client.
		Search().
		Index("bsqb_customize_v4").
		Type("bsqb_customize").
		Query(query).Sort("create_time", true).Size(40000000).
		Do(context.Background())
	e(err)

	for i, r := range foEsRes.Hits.Hits {
		fmt.Println(i, string(*r.Source))
	}
}

func e(err error) {
	if err != nil {
		panic(err)
	}
}
