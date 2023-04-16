package company_model

type CompanyInfo struct {
	Account               string `json:"account"`
	Name                  string `json:"name"`
	Address               string `json:"address"`
	ManagerName           string `json:"manager_name"`
	ManagerTel            string `json:"manager_tel"`
	BusinessLicenseNumber string `json:"business_license_number"`
}
