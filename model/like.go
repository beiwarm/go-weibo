package model

import "gorm.io/gorm"

type Like struct {
	gorm.Model
	TweetID int    `gorm:"not null;"`
	UserID  string `gorm:"not null"`
}
