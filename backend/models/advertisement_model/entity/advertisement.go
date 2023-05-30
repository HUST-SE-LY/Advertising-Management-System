package entity

import (
	"backend/models/advertisement_model"
)

type Advertisement struct {
	Id        int64  `json:"id" gorm:"column:id; primaryKey; autoIncrement"`
	CompanyId int64  `json:"company_id" gorm:"column:company_id; not null;"`
	ImageUrl  string `json:"image_url" gorm:"column:image_url"`
	advertisement_model.AdvertisementInfo
}

func (a Advertisement) GetInfo() advertisement_model.AdvertisementToBePreviewedInfo {
	return *advertisement_model.NewAdvertisementInfo(
		a.Id,
		a.CompanyId,
		a.ImageUrl,
		a.Title,
		a.Position,
		a.JumpToUrl,
		a.DisplayTime,
	)
}

func NewAdvertisementWithoutId(companyid int64, imageurl string, advertisementInfo advertisement_model.AdvertisementInfo) *Advertisement {
	return &Advertisement{
		CompanyId:         companyid,
		ImageUrl:          imageurl,
		AdvertisementInfo: advertisementInfo,
	}
}
