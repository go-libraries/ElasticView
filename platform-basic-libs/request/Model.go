package request

type EsConnect struct {
	Ip      string `json:"ip"`
	User    string `json:"user"`
	Pwd     string `json:"pwd"`
	Version int    `json:"version"`
}
