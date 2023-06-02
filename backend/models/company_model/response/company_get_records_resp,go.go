package response

import entity3 "backend/models/record_model/entity"

type CompanyGetRecordsResp struct {
	Company_records []entity3.ConsumeRecord `json:"company_records"`
}
