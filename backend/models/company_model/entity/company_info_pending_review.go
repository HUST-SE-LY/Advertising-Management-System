package entity

import "backend/models/company_model"

type CompanyInfoPendingReview struct {
	Id int64 `gorm:"column:id; not null; primaryKey; autoIncrement"`
	company_model.CompanyInfo
}

func (com CompanyInfoPendingReview) ToCompanyInfo() company_model.CompanyInfo {
	return company_model.CompanyInfo{
		Account:               com.Account,
		Name:                  com.Name,
		Address:               com.Address,
		ManagerName:           com.ManagerName,
		ManagerTel:            com.ManagerTel,
		BusinessLicenseNumber: com.BusinessLicenseNumber,
	}
}
