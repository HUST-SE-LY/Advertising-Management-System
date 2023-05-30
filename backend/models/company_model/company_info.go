package company_model

type CompanyInfo struct {
	Account               string `json:"account" gorm:"column:account; not null; unique" example:"联创"`
	Name                  string `json:"name" gorm:"column:name;" example:"联小创在线科技有限公司"`
	Address               string `json:"address" gorm:"column:address" example:"启明学院亮胜楼八楼"`
	ManagerName           string `json:"manager_name" gorm:"column:manager_name" example:"汉堡"`
	ManagerTel            string `json:"manager_tel" gorm:"column:manager_tel" example:"1919810"`
	BusinessLicenseNumber string `json:"business_license_number" gorm:"column:business_license_number" example:"114514"`
}

func NewCompanyInfo(account, name, address, managerName, managerTel, businessLicenseNumber string) *CompanyInfo {
	return &CompanyInfo{
		Account:               account,
		Name:                  name,
		Address:               address,
		ManagerName:           managerName,
		ManagerTel:            managerTel,
		BusinessLicenseNumber: businessLicenseNumber,
	}
}
