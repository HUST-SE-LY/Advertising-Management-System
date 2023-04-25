package entity

import (
	"backend/models/company_model"
)

type CompanyPendingReview struct {
	Company
}

func NewCompanyPendingReview(company Company) *CompanyPendingReview {
	return &CompanyPendingReview{
		Company: company,
	}
}

func (com CompanyPendingReview) ToCompany() Company {
	return com.Company
}

func (com CompanyPendingReview) GetInfo() company_model.CompanyInfo {
	return *company_model.NewCompanyInfo(
		com.Account,
		com.Name,
		com.Address,
		com.ManagerName,
		com.ManagerTel,
		com.BusinessLicenseNumber,
	)
}
