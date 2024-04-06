package model

import "gorm.io/gorm"

type UserFollow struct {
	gorm.Model
	FollowerID string `gorm:"not null;"`
	FolloweeID string `gorm:"not null;"`
}
