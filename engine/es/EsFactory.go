package es

import (
	"errors"

	"ElasticView/platform-basic-libs/request"
)

func GetEsClient(config request.EsConnect) (EsClient, error) {
	switch config.Version {
	case 6:
		return NewEsClientV6(config)
	case 7:
		return NewEsClientV7(config)
	default:
		return nil, errors.New("无效的版本号")
	}
}
