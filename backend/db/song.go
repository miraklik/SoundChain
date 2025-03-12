package db

import (
	"gorm.io/gorm"
)

type Song struct {
	gorm.Model
	NameSong string `gorm:"size:255;not null"`
	Artist   string `gorm:"size:255;not null"`
	AlbumID  uint
	Album    *Album `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	ImageURL string `gorm:"size:500;not null"`
}

type Album struct {
	gorm.Model
	Title       string `gorm:"size:255;not null"`
	Artist      string `gorm:"size:255;not null"`
	ReleaseDate string `gorm:"size:255;not null"`
	Songs       []Song `gorm:"foreginKey:AlbumID"`
	ImageURL    string `gorm:"size:500;not null"`
}
