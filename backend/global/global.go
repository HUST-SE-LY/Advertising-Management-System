package global

import (
	"gorm.io/gorm"
)

var (
	GVA_DB          *gorm.DB
	GVA_APP_SETTING *AppSetting
	GVA_DB_SETTING  *DatabaseSetting
)
