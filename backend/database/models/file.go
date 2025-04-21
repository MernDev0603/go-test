package models

import (
	"github.com/jinzhu/gorm"
)

type UserFile struct {
	ID        uint   `gorm:"primary_key"`
	UserID    string `gorm:"not null"`
	Filename  string `gorm:"not null"`
	Filepath  string `gorm:"not null"`
	UploadedAt string `gorm:"default:current_timestamp"`
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&UserFile{})
}
