package initialize

import (
	"backend/global"
	"backend/models/admin_model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

func Gorm() *gorm.DB {
	var err error
	var db *gorm.DB
	databaseSetting := global.GVA_DB_SETTING
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", databaseSetting.User, databaseSetting.Password, databaseSetting.Host, databaseSetting.Name)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   databaseSetting.TablePrefix,
			SingularTable: true,
		},
	})

	if err != nil {
		log.Fatalf("models.Settings error: %v\n", err)
	}

	db.AutoMigrate(&admin_model.Admin{})
	return db
}
