package entity

import (
	"backend/global"
	"time"
)

type AdvertisingSpace struct {
	Id      int64  `json:"id" gorm:"column:id;primaryKey; autoIncrement"`
	Price   int    `json:"price" gorm:"column: price;default:100"`
	Status  string `json:"status" gorm:"column: status;default:'free'"`
	Enddate string `json:"enddate" gorm:"column:enddate;example:'2023-6-1'"`
	AdvID   int64  `json:"adv_id" gorm:"column:adv_id;default:0"`
	// Indicates when the ad space will be displayed
}

func GetAdvertisementSpaces() ([]AdvertisingSpace, error) {
	var adspaces []AdvertisingSpace
	if err := global.GVA_DB.Find(&adspaces).Error; err != nil {
		return nil, err
	}
	return adspaces, nil
}

func GetFreshSpace() error {
	var space []AdvertisingSpace
	var advs []Advertisement
	format := "2006-01-02"
	global.GVA_DB.Find(&advs)

	global.GVA_DB.Find(&space)
	settime := time.Now().Format(format)
	for _, s := range space {
		if s.Status == "free" {
			if s.Enddate > settime {
				s.Status = "busy"
			}
		} else {
			if s.Enddate < settime {
				s.Status = "free"
			}
		}
		for _, a := range advs {
			if s.Id == int64(a.Position) {
				if settime >= a.StartDate {
					if settime < a.EndDate {
						s.AdvID = a.Id
					}
				}
			}
		}
		if err := global.GVA_DB.Model(&s).Updates(s).Error; err != nil {
			return err
		}
	}
	return nil
}

// IsFreeIn : 用于判断 AdvertisingSpace 在给定的时间范围内是否空闲
//func (a *AdvertisingSpace) IsFreeIn(displayTime advertisement_model.DisplayTime) bool {
//	for _, dt := range a.DisplayTimes {
//		if (dt.StartDate < displayTime.StartDate && dt.EndDate > displayTime.StartDate) || (displayTime.StartDate < dt.StartDate && displayTime.EndDate > dt.StartDate) {
//			return false
//		}
//	}
//	return true
