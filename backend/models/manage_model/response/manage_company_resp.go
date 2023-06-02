package response

import "backend/models/company_model"

type GetCompaniesCountResp struct {
	Count int `json:"count"`
}

type GetCompaniesResp struct {
	CompanyInfos []company_model.CompanyInfo `json:"company_infos"`
}

type GetCompaniesToBeReviewedResp struct {
	CompanyInfos []company_model.CompanyInfo `json:"company_infos"`
}

type CompaniesUpdateInfo struct {
	Account                       string `json:"account" example:"联创"`
	PreviousName                  string `json:"previous_name"  example:"联小创在线科技有限公司"`
	PreviousAddress               string `json:"previous_address" example:"启明学院亮胜楼八楼"`
	PreviousManagerName           string `json:"previous_manager_name"  example:"汉堡"`
	PreviousManagerTel            string `json:"previous_manager_tel"  example:"1919810"`
	PreviousBusinessLicenseNumber string `json:"previous_business_license_number"  example:"114514"`
	NewName                       string `json:"new_name"  example:"联小创在线科技有限公司11"`
	NewAddress                    string `json:"new_address" example:"启明学院亮胜楼八楼11"`
	NewManagerName                string `json:"new_manager_name"  example:"汉堡11"`
	NewManagerTel                 string `json:"new_manager_tel"  example:"191981011"`
	NewBusinessLicenseNumber      string `json:"new_business_license_number"  example:"11451411"`
}

type GetCompaniesUpdateInfoResp struct {
	CompaniesUpdateInfoList []CompaniesUpdateInfo
}
