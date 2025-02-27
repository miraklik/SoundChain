package handlers

import (
	"net/http"
	"soundchain/db"
	"soundchain/repository"

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
			Image    string `json:"image"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		if req.NameSong == "" || req.Artist == "" || req.Image == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "name_song, artist, and image are required"})
			return
		}

		song, err := SongRepository.SaveSong(&db.Song{NameSong: req.NameSong, Artist: req.Artist, Image: req.Image})
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
