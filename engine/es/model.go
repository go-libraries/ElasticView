package es

import "ElasticView/engine/db"

type Json map[string]interface{}

type Sort struct {
	Field     string
	Ascending bool
}

type Page struct {
	PageNum  int
	PageSize int
}

func NewPage(page int, pageSize int) *Page {
	return &Page{
		PageNum:  int(db.CreatePage(page, pageSize)),
		PageSize: pageSize,
	}
}

type EsConnect struct {
	Ip      string `json:"ip"`
	User    string `json:"user"`
	Pwd     string `json:"pwd"`
	Version int    `json:"version"`
}

type EsCat struct {
	EsConnect EsConnect `json:"es_connect"`
	Cat       string    `json:"cat"`
}

type EsRest struct {
	EsConnect EsConnect `json:"es_connect"`
	Method    string    `json:"method"`
	Body      string    `json:"body"`
	Path      string    `json:"path"`
}

type EsIndexInfo struct {
	EsConnect EsConnect `json:"es_connect"`
	Settings  Json      `json:"settings"`
	IndexName string    `json:"index_name"`
	AliasName string    `json:"alias_name"`
	Types     string    `json:"types"`
}

type EsReIndexInfo struct {
	EsConnect        EsConnect `json:"es_connect"`
	SourceIndex      string    `json:"source_index"`
	DestinationIndex string    `json:"destination_index"`
}
