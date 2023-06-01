package entity

import "backend/models/record_model/enum"

type ConsumeRecord struct {
	Id        int64                 `gorm:"column:id; primaryKey; autoIncrement"`
	Account   string                `gorm:"column:account;not null;"`
	Startdate string                `gorm:"column:start_date;not null"`
	Enddate   string                `gorm:"column:ecd_date;not null"`
	Cost      int                   `gorm:"column:cost;default:0"`
	Status    enum.ReviewStatusType `gorm:"column:status;not null"`
}
