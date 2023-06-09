package entity

import "backend/models/company_model"

const COMPANY_ROLE = "role_company"

type Company struct {
	Id       int64  `gorm:"column:id; not null; primaryKey; autoIncrement"`
	Password string `gorm:"column:password"`
	company_model.CompanyInfo
	//Account               string `gorm:"column:account; not null; unique"`
	//Password              string `gorm:"column:password"`
	//Name                  string `gorm:"column:name;"`
	//Address               string `gorm:"column:address"`
	//ManagerName           string `gorm:"column:manager_name"`
	//ManagerTel            string `gorm:"column:manager_tel"`
	//BusinessLicenseNumber string `gorm:"column:business_license_number"`
}

func NewCompanyWithoutId(password string, companyInfo company_model.CompanyInfo) *Company {
	return &Company{
		Password:    password,
		CompanyInfo: companyInfo,
	}
}

func NewCompany(id int64, password string, companyInfo company_model.CompanyInfo) *Company {
	return &Company{
		Id:          id,
		Password:    password,
		CompanyInfo: companyInfo,
	}
}

func (c Company) GetInfo() company_model.CompanyInfo {
	return c.CompanyInfo
}
