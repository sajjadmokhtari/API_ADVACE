package model

import "gorm.io/gorm"

type City struct {
	ID   uint   `gorm:"primaryKey"` // بهتره از uint استفاده بشه
	Name string `gorm:"column:name"`
}

type Description struct {
	gorm.Model
	ID          uint `gorm:"primaryKey"`
	Description string
}
