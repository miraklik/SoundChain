package db

import (
	"gorm.io/gorm"
)

type Song struct {
	gorm.Model
	NameSong string `gorm:"size:255;not null"`
	Artist   string `gorm:"size:255;not null"`
	Image    string `gorm:"size:255;not null"`
}
