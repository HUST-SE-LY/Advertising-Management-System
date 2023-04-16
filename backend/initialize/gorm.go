package initialize

import (
	"backend/global"
	"backend/models/admin_model"
	"backend/models/company_model/entity"
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

	_ = db.AutoMigrate(&admin_model.Admin{})
	_ = db.AutoMigrate(&admin_model.AdminToken{})
	_ = db.AutoMigrate(&entity.Company{})
	_ = db.AutoMigrate(&entity.CompanyToBeReviewed{})

	return db
}
