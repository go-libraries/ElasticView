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

type EsConnectID struct {
	EsConnectID int `json:"es_connect"`
}

type EsSnapshotInfo struct {
	EsConnect        int      `json:"es_connect"`
	SnapshotInfoList []string `json:"snapshot_info_list"`
}

type EsConnect struct {
	Ip      string `json:"ip" db:"ip"`
	User    string `json:"user" db:"user"`
	Pwd     string `json:"pwd" db:"pwd"`
	Version int    `json:"version" db:"version"`
}

type EsCat struct {
	EsConnect int    `json:"es_connect"`
	Cat       string `json:"cat"`
}

type EsRest struct {
	EsConnect int    `json:"es_connect"`
	Method    string `json:"method"`
	Body      string `json:"body"`
	Path      string `json:"path"`
}

type EsIndexInfo struct {
	EsConnect int    `json:"es_connect"`
	Settings  Json   `json:"settings"`
	IndexName string `json:"index_name"`
	Types     string `json:"types"`
}

type EsMappingInfo struct {
	IndexNameList []string `json:"index_name_list"`
	EsConnect     int      `json:"es_connect"`
	Mappings      Json     `json:"mappings"`
	IndexName     string   `json:"index_name"`
}

type EsTaskInfo struct {
	EsConnect    int      `json:"es_connect"`
	TaskId       []string `json:"task_id"`
	Actions      []string `json:"actions"`
	NodeId       []string `json:"node_id"`
	ParentTaskId string   `json:"parent_task_id"`
}

type EsAliasInfo struct {
	EsConnect        int      `json:"es_connect"`
	Settings         Json     `json:"settings"`
	IndexName        string   `json:"index_name"`
	AliasName        string   `json:"alias_name"`
	NewAliasNameList []string `json:"new_alias_name_list"`
	NewIndexList     []string `json:"new_index_list"`
	Types            int      `json:"types"`
}

type EsReIndexInfo struct {
	EsConnect int `json:"es_connect"`
	UrlValues struct {
		Timeout             string `json:"timeout"`
		RequestsPerSecond   int    `json:"requests_per_second"`
		Slices              int    `json:"slices"`
		Scroll              string `json:"scroll"`
		WaitForActiveShards string `json:"wait_for_active_shards"`
		Refresh             string `json:"refresh"`
		WaitForCompletion   *bool  `json:"wait_for_completion"`
	} `json:"url_values"`
	Body map[string]interface{} `json:"body"`
}
