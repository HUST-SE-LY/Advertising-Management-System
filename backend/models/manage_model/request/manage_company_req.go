package request

// CompaniesRegisterReq Both allow and reject
type CompaniesRegisterReq struct {
	CompanyAccounts []string `json:"company_accounts"`
}

// CompaniesUpdateReq Both allow and reject
type CompaniesUpdateReq struct {
	CompanyAccounts []string `json:"company_accounts"`
}
