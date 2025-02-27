package repository

import (
	"errors"
	"log"
	"soundchain/db"

	"gorm.io/gorm"
)

type SongRepository struct {
	DB *gorm.DB
}

func NewSong(db *gorm.DB) *SongRepository {
	return &SongRepository{DB: db}
}

func (sr *SongRepository) SaveSong(song *db.Song) (*db.Song, error) {
	if sr.DB == nil {
		log.Println("DB is nil")
		return nil, errors.New("DB is nil")
	}

	if err := sr.DB.Create(&song).Error; err != nil {
		log.Println("Failed to save song:", err)
		return nil, err
	}

	log.Printf("Song saved: %v", song)
	return song, nil
}

func (sr *SongRepository) GetAllSong() ([]db.Song, error) {
	var song []db.Song

	if err := sr.DB.Find(&song).Error; err != nil {
		log.Println("Failed to get all songs:", err)
		return nil, err
	}

	return song, nil
}

func (sr *SongRepository) GetSongByName(name string) (*db.Song, error) {
	var song db.Song

	if err := sr.DB.Preload("Groceries").Where("name_song=?", name).Find(&song).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("song not found")
		}

		return nil, err
	}

	return &song, nil
}

func (sr *SongRepository) GetSongByArtist(artist string) (*db.Song, error) {
	var song db.Song

	if err := sr.DB.Preload("Groceries").Where("artist=?", artist).Find(&song).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("Failed to get song by artist: %v", err)
			return nil, errors.New("song not found")
		}

		log.Printf("Failed to get song by artist: %v", err)
		return nil, err
	}

	log.Printf("Song found: %v", song)
	return &song, nil
}

func (sr *SongRepository) UpdateSong(song *db.Song) (*db.Song, error) {
	if err := sr.DB.First(&song, song.ID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("Failed to update song: %v", err)
			return nil, errors.New("song not found")
		}

		log.Printf("Failed to update song: %v", err)
		return nil, err
	}

	if err := sr.DB.Model(&db.Song{}).Where("id=?", song.ID).Updates(song).Error; err != nil {
		log.Printf("Failed to update song: %v", err)
		return nil, err
	}

	log.Printf("Song updated: %v", song)
	return song, nil
}

func (sr *SongRepository) DeleteSong(id uint) error {
	if err := sr.DB.Delete(&db.Song{}, id).Error; err != nil {
		log.Printf("Failed to delete song: %v", err)
		return err
	}

	return nil
}
