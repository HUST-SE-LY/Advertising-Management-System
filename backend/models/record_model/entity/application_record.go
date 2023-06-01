package entity

import "backend/models/record_model/enum"

type ApplicationRecord struct {
	Id      int64                 `gorm:"column:id; primaryKey; autoIncrement"`
	Account string                `gorm:"column:account;not null"`
	Status  enum.ReviewStatusType `gorm:"column:status;not null"`
	Type    enum.ApplicationType  `gorm:"column:type; not null"`
	Date    int                   `gorm:"column:date; not null"`
}
