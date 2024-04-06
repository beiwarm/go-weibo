package model

import "gorm.io/gorm"

type TokenRecord struct {
	gorm.Model
	UserName string `gorm:"not null;"`
	Token    string `gorm:"not null;unique"`
}
