package entity

import "backend/models/advertisement_model"

type AdvertisingSpace struct {
	Id int64 `json:"id" gorm:"column:id;primaryKey; autoIncrement"`
	// Indicates when the ad space will be displayed
	DisplayTimes []advertisement_model.DisplayTime `json:"display_times" gorm:"column:display_times"`
}
