package response

import "backend/models/company_model"

type GetCompaniesCountResp struct {
	Count int64 `json:"count"`
}

type GetCompaniesResp struct {
	CompanyInfos []company_model.CompanyInfo `json:"company_infos"`
}

type GetCompaniesToBeReviewedResp struct {
	CompanyInfos []company_model.CompanyInfo `json:"company_infos"`
}
