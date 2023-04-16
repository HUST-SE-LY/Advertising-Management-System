package company_model

type CompanyInfo struct {
	Account               string `json:"account" gorm:"column:account; not null; unique"`
	Name                  string `json:"name" gorm:"column:name;"`
	Address               string `json:"address" gorm:"column:address"`
	ManagerName           string `json:"manager_name" gorm:"column:manager_name"`
	ManagerTel            string `json:"manager_tel" gorm:"column:manager_tel"`
	BusinessLicenseNumber string `json:"business_license_number" gorm:"column:business_license_number"`
}
