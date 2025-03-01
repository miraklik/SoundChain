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

	"github.com/gin-gonic/gin"
)

var (
	SongRepository *repository.SongRepository
)

func CreateSong() gin.HandlerFunc {
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

		song, err := SongRepository.SaveSong(&db.Song{NameSong: req.NameSong, Artist: req.Artist})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"song": song})
	}
}

func UpdateSong() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			ID       uint   `json:"id"`
			NameSong string `json:"name_song"`
			Artist   string `json:"artist"`
			Image    string `json:"image"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if req.NameSong == "" || req.Artist == "" || req.Image == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "name_song, artist, and image are required"})
			return
		}

		song, err := SongRepository.UpdateSong(&db.Song{NameSong: req.NameSong, Artist: req.Artist, Image: req.Image})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"song": song})
	}
}

func DeleteSong() gin.HandlerFunc {
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

		if err := SongRepository.DeleteSong(req.ID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "song deleted"})
	}
}

func GetAllSongs() gin.HandlerFunc {
	return func(c *gin.Context) {
		songs, err := SongRepository.GetAllSong()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"songs": songs})
	}
}

func GetSongByArtist() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			Artist string `json:"artist"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
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

	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No image uploaded:" + err.Error()})
		return
	}

	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to open file" + err.Error()})
		return
	}
	defer src.Close()

	data, err := io.ReadAll(src)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read file" + err.Error()})
		return
	}

	bucket := cfg.AWS.Bucket
	key := fmt.Sprintf("images/%s", file.Filename)

	err = utils.UploadImage(bucket, key, data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to upload image" + err.Error()})
		return
	}

	if err := database.Create(&db.Song{Image: key}).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to save image" + err.Error()})
		return
	}

	url := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucket, key)
	c.JSON(http.StatusOK, gin.H{"message": "Image uploaded successfully", "url": url})
}
