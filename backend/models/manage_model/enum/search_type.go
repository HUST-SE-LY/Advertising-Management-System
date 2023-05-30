package enum

type CompanySearchType int

type AdvertisementType int

const (
	ACCOUNT CompanySearchType = iota
	Name
	Address
	ManagerName
	ManagerTel
	BusinessLicenseNumber
)

const (
	Id AdvertisementType = iota
	Companyid
	Imageurl
	Title
	Position
	Jumptourl
	startdate
	enddate
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

func (adv AdvertisementType) ToString() string {
	switch adv {
	case Id:
		return "id"
	case Companyid:
		return "company_id"
	case Imageurl:
		return "image_url"
	case Title:
		return "title"
	case Position:
		return "position"
	case Jumptourl:
		return "jump_to_url"
	case startdate:
		return "start_date"
	case enddate:
		return "end_date"
	default:
		return "id"
	}
}
