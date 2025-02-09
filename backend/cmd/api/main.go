package main

import (
	"log"
	"soundchain/config/config"
	"soundchain/db"
	"soundchain/handlers"

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

	router := r.Group("/api")

	router.POST("/register", server.RegisterUser)
	router.POST("/login", server.LoginUser)

	if err := r.Run(":" + cfg.Server.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
