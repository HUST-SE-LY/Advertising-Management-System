package company_model

type CompanyToBeReviewed struct {
	Id                    int64  `gorm:"column:id; not null; primaryKey; autoIncrement"`
	Account               string `gorm:"column:account; not null; unique"`
	Password              string `gorm:"column:password"`
	Name                  string `gorm:"column:name;"`
	Address               string `gorm:"column:address"`
	ManagerName           string `gorm:"column:manager_name"`
	ManagerTel            string `gorm:"column:manager_tel"`
	BusinessLicenseNumber string `gorm:"column:business_license_number"`
}

func (c CompanyToBeReviewed) ToCompany() Company {
	return Company{
		Id:                    c.Id,
		Account:               c.Account,
		Password:              c.Password,
		Name:                  c.Name,
		Address:               c.Address,
		ManagerName:           c.ManagerName,
		ManagerTel:            c.ManagerTel,
		BusinessLicenseNumber: c.BusinessLicenseNumber,
	}
}
