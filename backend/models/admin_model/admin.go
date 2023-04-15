package admin_model

const ADMIN_ROLE = "role_admin"

type Admin struct {
	Id       int64  `gorm:"column:id; not null; primarykey; auto_increment"`
	Account  string `gorm:"column:account; not null; unique"`
	Password string `gorm:"column:password; not null;"`
}
