package entity

import (
	"backend/models/company_model"
)

type CompanyPendingReview struct {
	Company
}

func (com CompanyPendingReview) ToCompany() Company {
	return com.Company
}

func (com CompanyPendingReview) GetInfo() company_model.CompanyInfo {
	return company_model.CompanyInfo{
		Account:               com.Account,
		Name:                  com.Name,
		Address:               com.Address,
		ManagerName:           com.ManagerName,
		ManagerTel:            com.ManagerTel,
		BusinessLicenseNumber: com.BusinessLicenseNumber,
	}
}
