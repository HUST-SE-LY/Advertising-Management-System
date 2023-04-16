package request

import "backend/models/company_model"

type CompanyRegisterReq struct {
	company_model.CompanyInfo
	Password string `json:"password"`
}
