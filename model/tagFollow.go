package model

import "gorm.io/gorm"

type TagFollow struct {
	gorm.Model
	FollowerID string `gorm:"not null;"`
	TagID      int    `gorm:"not null;"`
}
