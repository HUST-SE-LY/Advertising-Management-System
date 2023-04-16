package enum

type CompanySearchType int

const (
	ACCOUNT CompanySearchType = iota
	Name
	Address
	ManagerName
	ManagerTel
	BusinessLicenseNumber
)

func (com CompanySearchType) ToString() string {
	switch com {
	case ACCOUNT:
		return "account"
	case Name:
		return "name"
	case Address:
		return "address"
	case ManagerName:
		return "manager_name"
	case ManagerTel:
		return "manager_tel"
	case BusinessLicenseNumber:
		return "business_license_number"
	default:
		return "account"
	}
}
