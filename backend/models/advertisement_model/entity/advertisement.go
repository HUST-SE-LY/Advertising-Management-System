package entity

import "backend/models/advertisement_model"

type Advertisement struct {
	Id          int64  `json:"id" gorm:"column:id; primaryKey"`
	CompanyId   int64  `json:"company_id" gorm:"column:company_id; not null;"`
	Position    int    `json:"position" gorm:"column:position"`
	Name        string `json:"name" gorm:"column:name"`
	Description string `json:"description" gorm:"column:description"`
	ImageUrl    string `json:"image_url" gorm:"column:image_url"`
	JumpToUrl   string `json:"jump_to_url" gorm:"column:jump_to_url"`
	advertisement_model.DisplayTime
}
