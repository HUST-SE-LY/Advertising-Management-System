package admin_model

type AdminToken struct {
	AdminId int64  `gorm:"column:admin_id; primarykey;"`
	Token   string `gorm:"column:token;not null;"`
}
