package response

import "backend/models/company_model"

type GetCompaniesResp struct {
	CompanyInfos []company_model.CompanyInfo `json:"company_infos"`
}

type GetCompaniesToBeReviewedResp struct {
	CompanyInfos []company_model.CompanyInfo `json:"company_infos"`
}
