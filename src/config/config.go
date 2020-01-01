package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type conf struct {
	MysqlCfg string `json:"mysql_cfg"` // user:password@/dbname?charset=utf8&parseTime=True&loc=Local
}

var globalCfg conf

func GetMysqlCfg() string {
	return globalCfg.MysqlCfg
}

func init() {
	confDir := os.Getenv("taurus_conf_file")
	if confDir == "" {
		panic("taurus_conf_file not found")
	}
	file, err := os.Open(confDir)
	if err != nil {
		panic("open conf_file fail" + err.Error())
	}
	defer file.Close()
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic("read conf_file fail" + err.Error())
	}
	if err = json.Unmarshal(bytes, &globalCfg); err != nil {
		panic(err)
	}
}
