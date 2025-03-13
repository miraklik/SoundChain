package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"soundchain/config/config"
	"soundchain/db"
	"soundchain/repository"
	"soundchain/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

type SongHandler struct {
	SongRepository *repository.SongRepository
}

func (sh *SongHandler) CreateSong() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			NameSong string `json:"name_song"`
			Artist   string `json:"artist"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		if req.NameSong == "" || req.Artist == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "name_song, artist, and image are required"})
			return
		}

		song, err := sh.SongRepository.SaveSong(&db.Song{NameSong: req.NameSong, Artist: req.Artist})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"song": song})
	}
}

func (sh *SongHandler) UpdateSong() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			ID       uint   `json:"id"`
			NameSong string `json:"name_song"`
			Artist   string `json:"artist"`
			ImageURL string `json:"image_url"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if req.NameSong == "" || req.Artist == "" || req.ImageURL == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "name_song, artist, and image are required"})
			return
		}

		song, err := sh.SongRepository.UpdateSong(&db.Song{NameSong: req.NameSong, Artist: req.Artist, ImageURL: req.ImageURL})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"song": song})
	}
}

func (sh *SongHandler) DeleteSong() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			ID uint `json:"id"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if req.ID == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
			return
		}

		if err := sh.SongRepository.DeleteSong(req.ID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "song deleted"})
	}
}

func (sh *SongHandler) GetAllSongs() gin.HandlerFunc {
	return func(c *gin.Context) {
		songs, err := sh.SongRepository.GetAllSong()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"songs": songs})
	}
}

func (sh *SongHandler) GetSongByArtist() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			Artist string `json:"artist"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if req.Artist == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "artist is required"})
			return
		}

		data, err := sh.SongRepository.GetSongByArtist(req.Artist)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"songs": data})
	}
}

func (sh *SongHandler) CreateAlbum() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			Title  string `json:"title"`
			Artist string `json:"artist"`
			Image  string `json:"image"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		album, err := sh.SongRepository.CreateAlbum(&db.Album{Title: req.Title, Artist: req.Artist, ImageURL: req.Image})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"album": album})
	}
}

func (sh *SongHandler) UpdateAlbum() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			Title  string `json:"title"`
			Artist string `json:"artist"`
			Image  string `json:"image"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if req.Title == "" || req.Artist == "" || req.Image == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "title, artist, and image are required"})
			return
		}

		album, err := sh.SongRepository.UpdateAlbum(&db.Album{Title: req.Title, Artist: req.Artist, ImageURL: req.Image})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"album": album})
	}
}

func (sh *SongHandler) GetAllAlbums() gin.HandlerFunc {
	return func(c *gin.Context) {
		albums, err := sh.SongRepository.GetAllAlbums()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"albums": albums})
	}
}

func (sh *SongHandler) GetAlbumByName() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			Name string `json:"name_album"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if req.Name == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "name_album is required"})
			return
		}

		album, err := sh.SongRepository.GetAlbumByName(req.Name)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"album": album})
	}
}

func (sh *SongHandler) DeleteAlbum() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			ID uint `json:"id"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if req.ID == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
			return
		}

		if err := sh.SongRepository.DeleteAlbum(req.ID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "album deleted"})
	}
}

func UploadImage(c *gin.Context) {
	cfg, err := config.Load()
	if err != nil {
		log.Printf("Failed to load config: %v", err)
		return
	}

	database, err := db.ConnectDB()
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return
	}

	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	imageType := c.DefaultQuery("type", "song")

	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No image uploaded: " + err.Error()})
		return
	}

	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to open file: " + err.Error()})
		return
	}
	defer src.Close()

	data, err := io.ReadAll(src)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read file: " + err.Error()})
		return
	}

	if file.Size > 5*1024*1024 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Image too large"})
		return
	}

	if !strings.HasPrefix(file.Header.Get("Content-Type"), "image/") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Only images are allowed"})
		return
	}

	var key string
	if imageType == "album" {
		key = fmt.Sprintf("images/albums/%s", file.Filename)
	} else {
		key = fmt.Sprintf("images/songs/%s", file.Filename)
	}

	err = utils.UploadImage(cfg.AWS.Bucket, key, data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to upload image: " + err.Error()})
		return
	}

	imageURL := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", cfg.AWS.Bucket, key)

	if imageType == "album" {
		albumID := c.Query("album_id")
		if albumID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "album_id is required"})
			return
		}
		if err := database.Model(&db.Album{}).Where("id = ?", albumID).Update("ImageURL", imageURL).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save album image: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Album image uploaded successfully", "url": imageURL})
	} else {
		songID := c.Query("song_id")
		if songID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "song_id is required"})
			return
		}
		if err := database.Model(&db.Song{}).Where("id = ?", songID).Update("ImageURL", imageURL).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save song image: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Song image uploaded successfully", "url": imageURL})
	}
}
