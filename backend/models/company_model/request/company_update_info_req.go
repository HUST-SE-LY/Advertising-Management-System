package request

import "backend/models/company_model"

type CompanyUpdateInfoReq struct {
	company_model.CompanyInfo
}
