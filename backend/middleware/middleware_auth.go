package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"soundchain/config/config"
	"soundchain/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

const (
	AuthorizationHeader    = "Authorization"
	BearerPrefix           = "Bearer "
	UserRoleHeader         = "X-User-Role"
	BuyerIDHeader          = "X-Buyer-ID"
	ErrorUnauthorized      = "Unauthorized"
	ErrorInvalidToken      = "Invalid token"
	ErrorDatabase          = "Database error"
	ErrorConfig            = "Config error"
	ErrorNotFound          = "No NFTs found"
	ErrorInsufficientFunds = "Insufficient funds"
)

func parseToken(tokenStr string) (jwt.MapClaims, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, errors.New(ErrorConfig)
	}

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(cfg.JWT.Secret), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New(ErrorInvalidToken)
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid claims")
	}

	return claims, nil
}

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := utils.ValidateToken(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"Unauthorized": "Authentication required"})
			fmt.Println(err)
			c.Abort()
			return
		}
		c.Next()
	}
}
