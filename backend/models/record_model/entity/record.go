package entity

import (
	"backend/models/advertisement_model"
)

type Record struct {
	Id         int64  `json:"id" gorm:"column:id; primaryKey; autoIncrement"`
	AdvID      int64  `json:"advID" gorm:"column:adv_id"`
	ComAccount string `json:"ComAccount" gorm:"ComAccount"`
	Status     string `json:"status" gorm:"column:status;default:'pending'"`
	Duration   int    `json:"duration" gorm:"column:duration"`
	Cost       int    `json:"cost" gorm:"column:cost"`
	advertisement_model.AdvertisementInfo
}

//func (record *Record) GetIfo {
//
//}
