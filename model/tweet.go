package model

import "gorm.io/gorm"

type Tweet struct {
	gorm.Model
	UserName string `gorm:"not null;"`
	Content  string `gorm:"type:text;not null;"`
	Image    string
	Tag      string
	Like     int `gorm:"not null;"`
}
