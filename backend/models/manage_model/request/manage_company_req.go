package request

// CompaniesRegisterReq Both allow and reject
type CompaniesRegisterReq struct {
	CompanyAccounts []string `json:"company_accounts"`
}

type AllowCompaniesUpdateReq struct {
	CompanyAccounts []string `json:"company_accounts"`
}
