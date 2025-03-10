package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			NameToken   string `json:"name_token"`
			SymbolToken string `json:"symbol_token"`
			Description string `json:"description"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		if req.NameToken == "" || req.SymbolToken == "" || req.Description == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "name_token, symbol_token, and description are required"})
			return
		}

	}
}
