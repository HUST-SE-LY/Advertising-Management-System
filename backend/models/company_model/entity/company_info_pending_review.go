package entity

import "backend/models/company_model"

type CompanyInfoPendingReview struct {
	Id int64 `gorm:"column:id; not null; primaryKey; autoIncrement"`
	company_model.CompanyInfo
}

func NewCompanyInfoPendingReview(id int64, companyInfo company_model.CompanyInfo) *CompanyInfoPendingReview {
	return &CompanyInfoPendingReview{
		Id:          id,
		CompanyInfo: companyInfo,
	}
}

func (com CompanyInfoPendingReview) ToCompanyInfo() company_model.CompanyInfo {
	return com.CompanyInfo
}
