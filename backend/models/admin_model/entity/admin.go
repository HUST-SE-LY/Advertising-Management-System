package entity

const ADMIN_ROLE = "role_admin"

type Admin struct {
	Id       int64  `gorm:"column:id; not null; primaryKey; autoIncrement"`
	Account  string `gorm:"column:account; not null; unique"`
	Password string `gorm:"column:password; not null;"`
}

func NewAdminWithoutId(account, password string) *Admin {
	return &Admin{
		Account:  account,
		Password: password,
	}
}

func NewAdmin(id int64, account, password string) *Admin {
	return &Admin{
		Id:       id,
		Account:  account,
		Password: password,
	}
}
