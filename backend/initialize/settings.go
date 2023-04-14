package initialize

import (
	"backend/global"
	jsoniter "github.com/json-iterator/go"
	"log"
	"os"
)

type Config struct {
	App      global.AppSetting      `json:"app"`
	Database global.DatabaseSetting `json:"database"`
}

var config = &Config{}

func Settings() (*global.AppSetting, *global.DatabaseSetting) {
	fileContent, err := os.ReadFile("./conf.json")
	if err != nil {
		log.Fatalln("Can't read conf file")
	}

	err = jsoniter.Unmarshal(fileContent, config)
	if err != nil {
		log.Fatalln("Can't parse configuration")
	}
	return &config.App, &config.Database
}
