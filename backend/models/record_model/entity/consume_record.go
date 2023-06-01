package entity

import (
	"backend/models/record_model/enum"
)

type ConsumeRecord struct {
	Id       int64                 `gorm:"column:id; primaryKey; autoIncrement"`
	Account  string                `gorm:"column:account;not null;"`
	Date     int                   `gorm:"column:date;default:20230601;not null"`
	Start    string                `gorm:"column:start;default:'2023-06-01';not null"`
	End      string                `gorm:"column:end;default:'2023-06-01';not null"`
	Cost     int                   `gorm:"column:cost;default:0"`
	Position int                   `gorm:"column:position;not null"`
	Status   enum.ReviewStatusType `gorm:"column:status;not null"`
}
