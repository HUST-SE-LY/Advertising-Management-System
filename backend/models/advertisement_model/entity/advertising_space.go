package entity

type AdvertisingSpace struct {
	Id      int64  `json:"id" gorm:"column:id;primaryKey; autoIncrement"`
	Price   int    `json:"price" gorm:"column: price;default:100"`
	Status  string `json:"status" gorm:"column: status;default:'free'"`
	Enddate string `json:"enddate" gorm:"column:enddate;example:'2023-6-1'"`
	// Indicates when the ad space will be displayed
}

// IsFreeIn : 用于判断 AdvertisingSpace 在给定的时间范围内是否空闲
//func (a *AdvertisingSpace) IsFreeIn(displayTime advertisement_model.DisplayTime) bool {
//	for _, dt := range a.DisplayTimes {
//		if (dt.StartDate < displayTime.StartDate && dt.EndDate > displayTime.StartDate) || (displayTime.StartDate < dt.StartDate && displayTime.EndDate > dt.StartDate) {
//			return false
//		}
//	}
//	return true
//}
