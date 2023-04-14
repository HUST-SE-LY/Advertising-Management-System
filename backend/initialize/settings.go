package settings

import (
	jsoniter "github.com/json-iterator/go"
	"log"
	"os"
)

type AppSetting struct {
	JwtSecret string `json:"jwt_secret"`
}

type DatabaseSetting struct {
	Type        string `json:"type"`
	User        string `json:"user"`
	Password    string `json:"password"`
	Host        string `json:"host"`
	Name        string `json:"name"`
	TablePrefix string `json:"table_prefix"`
}

type Config struct {
	App      AppSetting      `json:"app"`
	Database DatabaseSetting `json:"database"`
}

var config = &Config{}

func GetAppSetting() *AppSetting {
	return &config.App
}

func GetDatabaseSetting() *DatabaseSetting {
	return &config.Database
}

func Setup() {
	fileContent, err := os.ReadFile("./conf.json")
	if err != nil {
		log.Fatalln("Can't read conf file")
	}

	err = jsoniter.Unmarshal(fileContent, config)
	if err != nil {
		log.Fatalln("Can't parse configuration")
	}
}
