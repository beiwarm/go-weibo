package dto

type Comment struct {
	Content string `gorm:"not null;"`
	UserID  string `gorm:"not null;"`
	ReplyID int
}
