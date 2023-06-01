package initialize

import (
	"backend/global"
	entity2 "backend/models/admin_model/entity"
	entity3 "backend/models/advertisement_model/entity"
	"backend/models/company_model/entity"
	entity4 "backend/models/record_model/entity"
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

	_ = db.AutoMigrate(&entity2.Admin{})
	_ = db.AutoMigrate(&entity2.AdminToken{})
	_ = db.AutoMigrate(&entity.Company{})
	_ = db.AutoMigrate(&entity.CompanyToken{})
	_ = db.AutoMigrate(&entity.CompanyPendingReview{})
	_ = db.AutoMigrate(&entity.CompanyInfoPendingReview{})
	_ = db.AutoMigrate(&entity3.Advertisement{})
	_ = db.AutoMigrate(&entity3.AdvertisementPendingReview{})
	_ = db.AutoMigrate(&entity3.AdvertisingSpace{})
	_ = db.AutoMigrate(&entity4.Record{})
	return db
}
