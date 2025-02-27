package main

import (
	"log"
	"soundchain/config/config"
	"soundchain/db"
	"soundchain/handlers"
	"soundchain/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	db, err := db.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	songRespository := repository.NewSong(db)

	handlers.SongRepository = songRespository

	r := gin.Default()

	r.POST("/Create", handlers.CreateSong())
	/*r.GET("/songs", handlers.GetAllSongs)
	r.GET("/song/:id", handlers.GetSongById)
	r.GET("/artist/:artist", handlers.GetSongByArtist)*/
	r.PUT("/update", handlers.UpdateSong())
	r.DELETE("/delete", handlers.DeleteSong())

	if err := r.Run(":" + cfg.Server.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
