package main

import (
	"backend/global"
	"backend/initialize"
	"backend/utils"
	"github.com/gin-gonic/gin"
)

func init() {
	global.GVA_APP_SETTING, global.GVA_DB_SETTING = initialize.Settings()
	global.GVA_DB = initialize.Gorm()
	utils.Setup()
}

func main() {

	r := gin.Default()
	initialize.Routers(r)

	err := r.Run(global.GVA_APP_SETTING.Address)

	if err != nil {
		return
	}
}
