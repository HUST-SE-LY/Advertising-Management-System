package response

type CompanyLoginResp struct {
	Account string `json:"account" example:"联创"`
	Token   string `json:"token" example:"114514"`
}
