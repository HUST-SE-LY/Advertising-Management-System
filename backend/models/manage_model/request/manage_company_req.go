package request

type AllowCompaniesRegisterReq struct {
	CompanyAccounts []string `json:"company_accounts"`
}

type AllowCompaniesUpdateReq struct {
	CompanyAccounts []string `json:"company_accounts"`
}
