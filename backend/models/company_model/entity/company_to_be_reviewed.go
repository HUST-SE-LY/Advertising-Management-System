package entity

import (
	"backend/models/company_model"
)

type CompanyToBeReviewed struct {
	Id                    int64  `gorm:"column:id; not null; primaryKey; autoIncrement"`
	Account               string `gorm:"column:account; not null; unique"`
	Password              string `gorm:"column:password"`
	Name                  string `gorm:"column:name;"`
	Address               string `gorm:"column:address"`
	ManagerName           string `gorm:"column:manager_name"`
	ManagerTel            string `gorm:"column:manager_tel"`
	BusinessLicenseNumber string `gorm:"column:business_license_number"`
}

func (com CompanyToBeReviewed) ToCompany() Company {
	return Company{
		Id:                    com.Id,
		Account:               com.Account,
		Password:              com.Password,
		Name:                  com.Name,
		Address:               com.Address,
		ManagerName:           com.ManagerName,
		ManagerTel:            com.ManagerTel,
		BusinessLicenseNumber: com.BusinessLicenseNumber,
	}
}

func (com CompanyToBeReviewed) GetInfo() company_model.CompanyInfo {
	return company_model.CompanyInfo{
		Account:               com.Account,
		Name:                  com.Name,
		Address:               com.Address,
		ManagerName:           com.ManagerName,
		ManagerTel:            com.ManagerTel,
		BusinessLicenseNumber: com.BusinessLicenseNumber,
	}
}
