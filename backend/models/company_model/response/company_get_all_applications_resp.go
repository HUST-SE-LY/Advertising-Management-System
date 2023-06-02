package response

import "backend/models/record_model/enum"

type ApplicationContent struct {
	Date   string                `json:"date" example:"2023-06-04"`
	Type   enum.ApplicationType  `json:"type" example:1"`
	Status enum.ReviewStatusType `json:"status" example:2"`
}

type CompanyGetAllApplicationsResp struct {
	ApplicationList []ApplicationContent
}
