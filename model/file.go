package model

type FileRecord struct {
	//ID的长度取uuid.IDLength
	ID         string `gorm:"type:char(14);primary_key;"`
	FileName   string `gorm:"unique"`
	ServerPath string
	NetPath    string
	FileExt    string
}
