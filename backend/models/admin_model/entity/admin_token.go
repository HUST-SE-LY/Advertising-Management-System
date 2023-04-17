package entity

type AdminToken struct {
	AdminId int64  `gorm:"column:admin_id; primarykey;"`
	Token   string `gorm:"column:token;not null;"`
}

func NewAdminToken(adminId int64, token string) *AdminToken {
	return &AdminToken{
		AdminId: adminId,
		Token:   token,
	}
}
