package main

import (
	"backend/global"
	"backend/initialize"
	"backend/utils"
	"github.com/gin-gonic/gin"

	_ "backend/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	global.GVA_APP_SETTING, global.GVA_DB_SETTING, global.GVA_FILE_SETTING = initialize.Settings()
	global.GVA_DB = initialize.Gorm()
	utils.Setup()
}

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server celler server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/api/v1

//	@securityDefinitions.basic	BasicAuth

//	@externalDocs.description	OpenAPI
//	@externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	r := gin.Default()
	initialize.Routers(r)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	err := r.Run(global.GVA_APP_SETTING.Address)

	if err != nil {
		return
	}
}
