package admin

type Admin struct {
	Id       int64  `gorm:"column: id; not null; primary key; auto_increment"`
	Account  string `gorm:"column: account; not null; unique"`
	Password string `gorm:"column: password; not null;"`
}
