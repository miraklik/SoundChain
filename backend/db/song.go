package db

import (
	"errors"
	"log"

	"gorm.io/gorm"
)

type Song struct {
	gorm.Model
	NameSong string `gorm:"size:255;not null"`
	Artist   string `gorm:"size:255;not null"`
	Image    string `gorm:"size:255;not null"`
}

func GetSongByName(name string) (Song, error) {
	var song Song
	db, err := ConnectDB()
	if err != nil {
		log.Println(err)
		return Song{}, err
	}

	if err := db.Preload("Groceries").Where("name_song=?", name).Find(&song).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return Song{}, errors.New("song not found")
		}
	}

	return song, nil
}

func GetSongByArtist(artist string) (Song, error) {
	var song Song
	db, err := ConnectDB()
	if err != nil {
		log.Println(err)
		return Song{}, err
	}

	if err := db.Preload("Groceries").Where("artist=?", artist).Find(&song).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return Song{}, errors.New("song not found")
		}
	}

	return song, nil
}
