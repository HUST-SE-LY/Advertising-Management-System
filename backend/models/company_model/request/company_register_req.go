package request

type CompanyRegisterReq struct {
	Account               string `json:"account"`
	Password              string `json:"password"`
	Name                  string `json:"name"`
	Address               string `json:"address"`
	ManagerName           string `json:"manager_name"`
	ManagerTel            string `json:"manager_tel"`
	BusinessLicenseNumber string `json:"business_license_number"`
}
