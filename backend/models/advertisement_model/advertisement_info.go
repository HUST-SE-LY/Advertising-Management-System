package advertisement_model

type AdvertisementInfo struct {
	Title     string `json:"title" gorm:"column:title" example:"advertisement example"`
	Position  int    `json:"position" gorm:"column:position" example:"1"`
	JumpToUrl string `json:"jump_to_url" gorm:"column:jump_to_url" example:"example.com/1919810.jpg"`
	DisplayTime
}
type AdvertisementToBePreviewedInfo struct {
	Id        int64  `json:"id" gorm:"column:id; primaryKey; autoIncrement"`
	CompanyId int64  `json:"company_id" gorm:"column:company_id; not null;"`
	ImageUrl  string `json:"image_url" gorm:"column:image_url"`
	Title     string `json:"title" gorm:"column:title"`
	Position  int    `json:"position" gorm:"column:position"`
	JumpToUrl string `json:"jump_to_url" gorm:"column:jump_to_url"`
	DisplayTime
}

func NewAdvertisementInfo(id int64, companyid int64, uml string, title string, position int, jump_to_url string, displaytime DisplayTime) *AdvertisementToBePreviewedInfo {
	return &AdvertisementToBePreviewedInfo{
		Id:          id,
		CompanyId:   companyid,
		ImageUrl:    uml,
		Title:       title,
		Position:    position,
		JumpToUrl:   jump_to_url,
		DisplayTime: displaytime,
	}

}
