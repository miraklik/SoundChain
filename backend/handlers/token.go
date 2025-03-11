package handlers

import (
	"math/big"
	"net/http"
	"soundchain/config/config"
	"soundchain/services"

	"github.com/gin-gonic/gin"
)

func CreateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			NameToken   string   `json:"name_token"`
			SymbolToken string   `json:"symbol_token"`
			TotalSupply *big.Int `json:"totalSupply"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if req.NameToken == "" || req.SymbolToken == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "name_token и symbol_token обязательны"})
			return
		}

		if req.TotalSupply == nil {
			req.TotalSupply = big.NewInt(0)
		}

		cfg, err := config.Load()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		service, err := services.NewEthereumService(
			cfg.Blockchain.RPC_URL,
			cfg.Blockchain.MemcoinAddr,
			cfg.Blockchain.MemcoinFactoryAddr,
			cfg.Blockchain.PrivateKey,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if err := service.CreateToken(req.NameToken, req.SymbolToken, req.TotalSupply); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "token created"})
	}
}
