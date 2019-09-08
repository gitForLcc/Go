package main

import (
	"encoding/json"

	"github.com/astaxie/beego/logs"
)

func main() {
	config := make(map[string]interface{})
	config["filename"] = "./logs/test.log"
	config["debug"] = logs.LevelDebug

	conf, err := json.Marshal(config)
	if err != nil {
		panic(err)
	}
	logger := logs.NewLogger()
	logger.SetLogger("file", string(conf))

	logger.Debug("init logger success, level: %v", config["debug"])

}
