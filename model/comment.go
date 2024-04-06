package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content string `gorm:"not null;"`
	UserID  string `gorm:"not null;"`
	ReplyID int
}
