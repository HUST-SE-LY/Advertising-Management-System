package company_model

type CompanyToken struct {
	CompanyId int64  `gorm:"column:company_id; primaryKey;"`
	Token     string `gorm:"column:token; not null"`
}
