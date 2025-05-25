package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"soundchain/config/config"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

const (
	AuthorizationHeader = "Authorization"
	BearerPrefix        = "Bearer "
	ContextUserIDKey    = "userID"
	ContextUserRoleKey  = "userRole"
)

var (
	ErrUnauthorized = errors.New("authentication required")
	ErrInvalidToken = errors.New("invalid token")
	ErrConfigLoad   = errors.New("could not load config")
)

func parseToken(tokenStr string) (jwt.MapClaims, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, errors.New(ErrConfigLoad.Error())
	}

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(cfg.JWT.Secret), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New(ErrInvalidToken.Error())
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid claims")
	}

	return claims, nil
}

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader(AuthorizationHeader)
		if !strings.HasPrefix(authHeader, BearerPrefix) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": ErrUnauthorized.Error()})
			c.Abort()
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, BearerPrefix)
		claims, err := parseToken(tokenStr)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		if userID, ok := claims["user_id"].(string); ok {
			c.Set(ContextUserIDKey, userID)
		}

		if userRole, ok := claims["role"].(string); ok {
			c.Set(ContextUserRoleKey, userRole)
		}

		c.Next()
	}
}
