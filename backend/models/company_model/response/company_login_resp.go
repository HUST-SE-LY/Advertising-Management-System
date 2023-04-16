package response

import "backend/models/company_model"

type CompanyLoginResp struct {
	company_model.CompanyInfo
}
