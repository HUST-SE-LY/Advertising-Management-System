package entity

import "backend/models/advertisement_model"

type AdvertisingSpace struct {
	Id int64 `json:"id" gorm:"column:id;primaryKey; autoIncrement"`
	// Indicates when the ad space will be displayed
	DisplayTimes []advertisement_model.DisplayTime `json:"display_times" gorm:"column:display_times"`
}

// IsFreeIn : 用于判断 AdvertisingSpace 在给定的时间范围内是否空闲
func (a *AdvertisingSpace) IsFreeIn(displayTime advertisement_model.DisplayTime) bool {
	for _, dt := range a.DisplayTimes {
		if (dt.StartDate < displayTime.StartDate && dt.EndDate > displayTime.StartDate) || (displayTime.StartDate < dt.StartDate && displayTime.EndDate > dt.StartDate) {
			return false
		}
	}
	return true
}
