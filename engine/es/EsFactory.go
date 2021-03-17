package es

import (
	"errors"
)

func GetEsClient(config EsConnect) (EsClient, error) {
	if config.Ip == "" {
		return nil, errors.New("请先选择ES连接")
	}
	switch config.Version {
	case 6:
		return NewEsClientV6(config)
	case 7:
		return NewEsClientV7(config)
	default:
		return nil, errors.New("无效的版本号")
	}
}
