package model

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	Content string `gorm:"not null;unique;"`
	Like    int    `gorm:"not null;"`
}
