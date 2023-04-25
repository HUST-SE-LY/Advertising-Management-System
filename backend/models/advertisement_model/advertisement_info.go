package advertisement_model

type AdvertisementInfo struct {
	Title     string `json:"title" gorm:"column:title"`
	Position  int    `json:"position" gorm:"column:position"`
	JumpToUrl string `json:"jump_to_url" gorm:"column:jump_to_url"`
	DisplayTime
}
