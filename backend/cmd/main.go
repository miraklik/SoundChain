package main

import (
	"log"
	"soundchain/config/config"
	"soundchain/db"
	"soundchain/handlers"
	"soundchain/middleware"
	"soundchain/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DBInit() *gorm.DB {
	db, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	r := gin.Default()

	db := DBInit()

	server := handlers.NewServer(db)
	songRespository := repository.NewSong(db)
	SongHandler := handlers.SongHandler{SongRepository: songRespository}
	router := r.Group("/api")

	router.POST("/register", server.RegisterUser)
	router.POST("/login", server.LoginUser)

	auth := router.Group("/")
	auth.Use(middleware.JwtAuthMiddleware())
	{
		auth.POST("/songs", SongHandler.CreateSong())
		auth.GET("/songs", SongHandler.GetAllSongs())
		auth.POST("/upload", handlers.UploadImage)
		auth.GET("/artist/:artist", SongHandler.GetSongByArtist())
		auth.PUT("/song/:id", SongHandler.UpdateSong())
		auth.DELETE("/song/:id", SongHandler.DeleteSong())

		auth.POST("/create-album", SongHandler.CreateAlbum())
		auth.GET("/albums", SongHandler.GetAllAlbums())
		auth.PUT("/album/:id", SongHandler.UpdateAlbum())
		auth.GET("/album/:name", SongHandler.GetAlbumByName())
		auth.DELETE("/album/id", SongHandler.DeleteAlbum())

		auth.POST("/create-token", handlers.CreateToken())
	}
	if err := r.Run(":" + cfg.Server.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
