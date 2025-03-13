package repository

import (
	"fmt"
	"soundchain/db"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	logger *zap.Logger
)

type SongRepository struct {
	DB *gorm.DB
}

func NewSong(db *gorm.DB) *SongRepository {
	return &SongRepository{DB: db}
}

func init() {
	logger, _ = zap.NewDevelopment()
	defer logger.Sync()
}

func (sr *SongRepository) SaveSong(song *db.Song) (*db.Song, error) {
	if sr.DB == nil {
		logger.Info("DB is nil")
		return nil, fmt.Errorf("DB is nil")
	}

	if err := sr.DB.Create(&song).Error; err != nil {
		logger.Error("Failed to save song:", zap.Error(err))
		return nil, fmt.Errorf("failed to save song: %v", err)
	}

	logger.Info("Song saved:", zap.Any("song", song))
	return song, nil
}

func (sr *SongRepository) GetAllSong() ([]db.Song, error) {
	var songs []db.Song

	if err := sr.DB.Find(&songs).Error; err != nil {
		logger.Error("Failed to get all songs:", zap.Error(err))
		return nil, fmt.Errorf("failed to get all songs: %v", err)
	}

	return songs, nil
}

func (sr *SongRepository) GetSongByName(name string) (*db.Song, error) {
	var song db.Song

	if err := sr.DB.Preload("Groceries").Where("name_song=?", name).Find(&song).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			logger.Warn("Song not found by name:", zap.String("name", name))
			return nil, fmt.Errorf("failed to get song by name: %v", err)
		}

		logger.Error("Failed to get song by name:", zap.Error(err))
		return nil, fmt.Errorf("failed to get song by name: %v", err)
	}

	return &song, nil
}

func (sr *SongRepository) GetSongByArtist(artist string) (*db.Song, error) {
	var song db.Song

	if err := sr.DB.Preload("Groceries").Where("artist=?", artist).Find(&song).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			logger.Warn("Failed to get song by artist:", zap.String("artist", artist))
			return nil, fmt.Errorf("failed to get song by artist: %v", err)
		}

		logger.Error("Failed to get song by artist:", zap.Error(err))
		return nil, fmt.Errorf("failed to get song by artist: %v", err)
	}

	logger.Info("Song found:", zap.Any("song", song))
	return &song, nil
}

func (sr *SongRepository) UpdateSong(song *db.Song) (*db.Song, error) {
	if err := sr.DB.First(&song, song.ID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			logger.Warn("Failed to update song:", zap.Uint("song_id", song.ID))
			return nil, fmt.Errorf("failed to update song: %v", err)
		}

		logger.Error("Failed to update song:", zap.Error(err))
		return nil, err
	}

	if err := sr.DB.Model(&db.Song{}).Where("id=?", song.ID).Updates(song).Error; err != nil {
		logger.Error("Failed to update song:", zap.Error(err))
		return nil, fmt.Errorf("failed to update song: %v", err)
	}

	logger.Info("Song updated:", zap.Any("song", song))
	return song, nil
}

func (sr *SongRepository) DeleteSong(id uint) error {
	result := sr.DB.Unscoped().Delete(&db.Song{}, id)
	if result.Error != nil {
		logger.Error("Failed to delete song", zap.Error(result.Error))
		return fmt.Errorf("failed to delete song: %v", result.Error)
	}

	if result.RowsAffected == 0 {
		logger.Warn("Song not found for deletion", zap.Uint("song_id", id))
		return fmt.Errorf("song not found for deletion")
	}

	logger.Info("Song deleted", zap.Uint("song_id", id))
	return nil
}

func (sr *SongRepository) CreateAlbum(album *db.Album) (*db.Album, error) {
	if err := sr.DB.Create(&album).Error; err != nil {
		logger.Error("Failed to create album:", zap.Error(err))
		return nil, fmt.Errorf("failed to create album: %v", err)
	}

	logger.Info("Album created:", zap.Any("album", album))
	return album, nil
}

func (sr *SongRepository) UpdateAlbum(album *db.Album) (*db.Album, error) {
	if err := sr.DB.First(&album, album.ID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			logger.Warn("Failed to update album:", zap.Uint("album_id", album.ID))
			return nil, fmt.Errorf("failed to update album: %v", err)
		}

		logger.Error("Failed to update album:", zap.Error(err))
		return nil, err
	}

	if err := sr.DB.Model(&db.Album{}).Where("id=?", album.ID).Updates(album).Error; err != nil {
		logger.Error("Failed to update album:", zap.Error(err))
		return nil, fmt.Errorf("failed to update album: %v", err)
	}

	return album, nil
}

func (sr *SongRepository) GetAlbumByID(id uint) (*db.Album, error) {
	var album db.Album
	if err := sr.DB.Preload("Songs").First(&album, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			logger.Warn("Album not found", zap.Uint("album_id", id))
			return nil, nil
		}
		logger.Error("Failed to get album", zap.Error(err))
		return nil, fmt.Errorf("failed to get album: %v", err)
	}

	return &album, nil
}

func (sr *SongRepository) GetAllAlbums() ([]db.Album, error) {
	var albums []db.Album
	if err := sr.DB.Preload("Songs").Find(&albums).Error; err != nil {
		logger.Error("Failed to get albums", zap.Error(err))
		return nil, fmt.Errorf("failed to get albums: %v", err)
	}

	return albums, nil
}

func (sr *SongRepository) GetAlbumByName(name string) (*db.Album, error) {
	var album db.Album
	if err := sr.DB.Preload("Songs").Where("title=?", name).First(&album).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			logger.Warn("Album not found", zap.String("album_name", name))
			return nil, nil
		}
		logger.Error("Failed to get album", zap.Error(err))
		return nil, fmt.Errorf("failed to get album: %v", err)
	}

	return &album, nil
}

func (sr *SongRepository) DeleteAlbum(id uint) error {
	if err := sr.DB.Delete(&db.Album{}, id).Error; err != nil {
		logger.Error("Failed to delete album", zap.Error(err))
		return fmt.Errorf("failed to delete album: %v", err)
	}

	logger.Info("Album deleted", zap.Uint("album_id", id))
	return nil
}
