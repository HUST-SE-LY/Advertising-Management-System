package entity

import "backend/models/advertisement_model"

type Advertisement struct {
	Id        int64  `json:"id" gorm:"column:id; primaryKey"`
	CompanyId int64  `json:"company_id" gorm:"column:company_id; not null;"`
	ImageUrl  string `json:"image_url" gorm:"column:image_url"`
	advertisement_model.AdvertisementInfo
}
