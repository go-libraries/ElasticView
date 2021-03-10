package es

type EsClient interface {
	Ping() (interface{}, int, error)
}
