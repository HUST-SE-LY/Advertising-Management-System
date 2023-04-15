package response

type CompanyLoginResp struct {
	Account               string `gorm:"column:account; not null; unique"`
	Name                  string `gorm:"column:name;"`
	Address               string `gorm:"column:address"`
	ManagerName           string `gorm:"column:manager_name"`
	ManagerTel            string `gorm:"column:manager_tel"`
	BusinessLicenseNumber string `gorm:"column:business_license_number"`
}
