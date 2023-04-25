package entity

type CompanyToken struct {
	CompanyId int64  `gorm:"column:company_id; primaryKey;"`
	Token     string `gorm:"column:token; not null"`
}

func NewCompanyToken(companyId int64, token string) *CompanyToken {
	return &CompanyToken{
		CompanyId: companyId,
		Token:     token,
	}
}
