package advertisement_model

type AdvertisementInfo struct {
	Title     string `json:"title" gorm:"column:title"`
	Position  int    `json:"position" gorm:"column:position"`
	ImageUrl  string `json:"image_url" gorm:"column:image_url"`
	JumpToUrl string `json:"jump_to_url" gorm:"column:jump_to_url"`
	StartDate string `json:"start_date" gorm:"column:start_date"`
	EndDate   string `json:"end_date" gorm:"column:end_date"`
}
