package application

var GlobConfig Config

type Config struct {
	Log struct{
		StorageDays int `json:"storage_days"`
		LogDir string `json:"log_dir"`
	} `json:"log"`
	Port int `json:"port"`
	Mysql struct {
		Username string `json:"username"`
		Pwd string `json:"pwd"`
		IP string `json:"ip"`
		Port string `json:"port"`
		DbName string `json:"db_name"`
		MaxOpenConns int    `json:"max_open_conns"`
		MaxIdleConns int    `json:"max_idle_conns"`
	} `json:"mysql"`
}



