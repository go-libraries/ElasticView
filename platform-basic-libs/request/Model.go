package request

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
