package entity

import "backend/global"

type AdvertisingSpace struct {
	Id      int64  `json:"id" gorm:"column:id;primaryKey; autoIncrement"`
	Price   int    `json:"price" gorm:"column: price;default:100"`
	Status  string `json:"status" gorm:"column: status;default:'free'"`
	Enddate string `json:"enddate" gorm:"column:enddate;example:'2023-6-1'"`
	// Indicates when the ad space will be displayed
}

func GetAdvertisementSpaces() ([]AdvertisingSpace, error) {
	var adspaces []AdvertisingSpace
	if err := global.GVA_DB.Find(&adspaces).Error; err != nil {
		return nil, err
	}
	return adspaces, nil
}

//func GetFreshSpace() {
//	var space AdvertisingSpace
//	global.GVA_DB.Last(&space)
//	for i := 1; i <= int(space.Id); i++ {
//
//	}
//
//}

// IsFreeIn : 用于判断 AdvertisingSpace 在给定的时间范围内是否空闲
//func (a *AdvertisingSpace) IsFreeIn(displayTime advertisement_model.DisplayTime) bool {
//	for _, dt := range a.DisplayTimes {
//		if (dt.StartDate < displayTime.StartDate && dt.EndDate > displayTime.StartDate) || (displayTime.StartDate < dt.StartDate && displayTime.EndDate > dt.StartDate) {
//			return false
//		}
//	}
//	return true
//}
