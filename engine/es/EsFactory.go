package es

import (
	"errors"

	"ElasticView/engine/db"
	"ElasticView/platform-basic-libs/util"
)

func GetEsClientV6ByID(id int) (esClient EsClient, err error) {

	esConnect := EsConnect{}
	sql, args, err := db.SqlBuilder.
		Select("ip", "user", "pwd", "version").
		From("es_link").
		Where(db.Eq{"id": id}).
		Limit(1).ToSql()
	if err != nil {
		return
	}
	err = db.Sqlx.Get(&esConnect, sql, args...)

	if util.FilterMysqlNilErr(err) {
		return
	}
	if esConnect.Ip == "" {
		return nil, errors.New("请先选择ES连接")
	}
	return NewEsClientV6(esConnect)
}

func GetEsClient(esConnect EsConnect) (EsClient, error) {
	if esConnect.Ip == "" {
		return nil, errors.New("请先选择ES连接")
	}
	switch esConnect.Version {
	case 6:
		return NewEsClientV6(esConnect)
	case 7:
		return NewEsClientV7(esConnect)
	default:
		return nil, errors.New("无效的版本号")
	}
}
