package handlers

import (
	"log"
	"net/http"
	"soundchain/config/config"
	"soundchain/db"
	"soundchain/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type RegisterUser struct {
	Email           string `json:"email"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"passwordConfirm"`
}

type LoginUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Server struct {
	db *gorm.DB
}

func NewServer(db *gorm.DB) *Server {
	return &Server{db: db}
}

func (s *Server) RegisterUser(c *gin.Context) {
	var Input RegisterUser

	cfg, err := config.Load()
	if err != nil {
		log.Printf("Failed to load config: %v", err)
		return
	}

	if err := c.ShouldBindJSON(&Input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := db.User{Email: Input.Email, Username: Input.Username, Password: Input.Password}
	user.HashedPassword()

	if user.Email == "" || user.Username == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username and password are required"})
		return
	}

	if Input.Password != Input.PasswordConfirm {
		c.JSON(http.StatusBadRequest, gin.H{"error": "passwords do not match"})
		return
	}

	if err := utils.SendEmail(cfg.Email.SmtpUser, user.Email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := utils.ValidatePassword(Input.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := s.db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user created successfully"})
}

func (s *Server) LoginUser(c *gin.Context) {
	var Input LoginUser

	if err := c.ShouldBindJSON(&Input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := db.User{Username: Input.Username, Password: Input.Password}

	token, err := s.LoginCheck(user.Username, user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (s *Server) LoginCheck(username, password string) (string, error) {
	var err error

	user := db.User{}

	if err = s.db.Model(db.User{}).Where("username=?", username).Take(&user).Error; err != nil {
		return "", err
	}

	err = db.VerifyPassword(password, user.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := utils.GenerateToken(user)

	if err != nil {
		return "", err
	}

	return token, nil
}
