package services

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	memcoin "soundchain"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type EthereumService struct {
	Client         *ethclient.Client
	PrivateKey     *ecdsa.PrivateKey
	MemCoin        *memcoin.Memcoin
	MemCoinFactory *memcoin.Memcoinfactory
}

func NewEthereumService(rpcURL, memcoinAddr, MemcoinfactoryAddr, privateKeyHex string) (*EthereumService, error) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Ethereum: %w", err)
	}

	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, fmt.Errorf("invalid private key: %w", err)
	}

	memcoinAddress := common.HexToAddress(memcoinAddr)
	memCoinInstance, err := memcoin.NewMemcoin(memcoinAddress, client)
	if err != nil {
		return nil, fmt.Errorf("failed to bind MemCoin: %w", err)
	}

	memcoinFactoryAddress := common.HexToAddress(MemcoinfactoryAddr)
	memCoinFactoryInstance, err := memcoin.NewMemcoinfactory(memcoinFactoryAddress, client)
	if err != nil {
		return nil, fmt.Errorf("failed to bind MemCoinFactory: %w", err)
	}

	return &EthereumService{
		Client:         client,
		PrivateKey:     privateKey,
		MemCoin:        memCoinInstance,
		MemCoinFactory: memCoinFactoryInstance,
	}, nil
}

func (es *EthereumService) createAuth(ctx context.Context, gasLimit uint64, gasPriceWei *big.Int, valueEther float64) (*bind.TransactOpts, error) {
	chainID, err := es.Client.ChainID(ctx)
	if err != nil {
		log.Printf("Failed to get chain ID: %v", err)
		return nil, fmt.Errorf("failed to get chain ID: %w", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(es.PrivateKey, chainID)
	if err != nil {
		log.Printf("Failed to create auth: %v", err)
		return nil, fmt.Errorf("failed to create auth: %w", err)
	}

	auth.GasLimit = gasLimit

	if gasPriceWei != nil && gasPriceWei.Cmp(big.NewInt(0)) > 0 {
		auth.GasPrice = gasPriceWei
	} else {
		auth.GasPrice, err = es.Client.SuggestGasPrice(ctx)
		if err != nil {
			log.Printf("Failed to suggest gas price: %v", err)
			return nil, fmt.Errorf("failed to suggest gas price: %w", err)
		}
	}

	if valueEther > 0 {
		valueWei := new(big.Float).Mul(big.NewFloat(valueEther), big.NewFloat(1e18))
		valueWeiInt, _ := valueWei.Int(nil)
		auth.Value = valueWeiInt
	} else {
		auth.Value = big.NewInt(0)
	}

	return auth, nil
}

func (es *EthereumService) CreateToken(name, symbol string, totalSupply *big.Int) error {
	auth, err := es.createAuth(context.Background(), 0, nil, 0)
	if err != nil {
		log.Printf("Failed to create auth: %v", err)
		return fmt.Errorf("failed to create auth: %v", err)
	}

	tx, err := es.MemCoinFactory.CreateMemCoin(auth, name, symbol, totalSupply)
	if err != nil {
		log.Printf("Failed to create token: %v", err)
		return fmt.Errorf("failed to create token: %w", err)
	}

	log.Printf("Token created: %s", tx.Hash().Hex())
	return nil
}

func (es *EthereumService) Transfer(to string, amount *big.Int) error {
	auth, err := es.createAuth(context.Background(), 0, nil, 0)
	if err != nil {
		log.Printf("Failed to create auth: %v", err)
		return fmt.Errorf("failed to create auth: %v", err)
	}

	tx, err := es.MemCoin.Transfer(auth, common.HexToAddress(to), amount)
	if err != nil {
		log.Printf("Failed to transfer: %v", err)
		return fmt.Errorf("failed to transfer: %w", err)
	}

	log.Printf("Transfer successful: %s", tx.Hash().Hex())
	return nil
}

func (es *EthereumService) Withdraw(amount *big.Int) error {
	auth, err := es.createAuth(context.Background(), 0, nil, 0)
	if err != nil {
		log.Printf("Failed to create auth: %v", err)
		return fmt.Errorf("failed to create auth: %v", err)
	}

	tx, err := es.MemCoin.Withdraw(auth, amount)
	if err != nil {
		log.Printf("Failed to withdraw: %v", err)
		return fmt.Errorf("failed to withdraw: %w", err)
	}

	log.Printf("Withdrawal successful: %s", tx.Hash().Hex())
	return nil
}

func (es *EthereumService) AddLiquidity(amount *big.Int) error {
	auth, err := es.createAuth(context.Background(), 0, nil, 0)
	if err != nil {
		log.Printf("Failed to create auth: %v", err)
		return fmt.Errorf("failed to create auth: %v", err)
	}

	tx, err := es.MemCoin.AddLiquidity(auth, amount)
	if err != nil {
		log.Printf("Failed to add liquidity: %v", err)
		return fmt.Errorf("failed to add liquidity: %w", err)
	}

	log.Printf("Liquidity added: %s", tx.Hash().Hex())
	return nil
}

func (es *EthereumService) ExcludeFromTax(account string, excluded bool) error {
	addr := common.HexToAddress(account)

	auth, err := es.createAuth(context.Background(), 0, nil, 0)
	if err != nil {
		log.Printf("Failed to create auth: %v", err)
		return fmt.Errorf("failed to create auth: %v", err)
	}

	tx, err := es.MemCoin.ExcludeFromTax(auth, addr, excluded)
	if err != nil {
		log.Printf("Failed to exclude from tax: %v", err)
		return fmt.Errorf("failed to exclude from tax: %w", err)
	}

	log.Printf("Excluded from tax: %s", tx.Hash().Hex())
	return nil
}

func (es *EthereumService) SetTaxes(buyTax, sellTax *big.Int) error {
	auth, err := es.createAuth(context.Background(), 0, nil, 0)
	if err != nil {
		log.Printf("Failed to create auth: %v", err)
		return fmt.Errorf("failed to create auth: %v", err)
	}

	tx, err := es.MemCoin.SetTaxes(auth, buyTax, sellTax)
	if err != nil {
		log.Printf("Failed to set taxes: %v", err)
		return fmt.Errorf("failed to set taxes: %w", err)
	}

	log.Printf("Taxes set: %s", tx.Hash().Hex())
	return nil
}

func (es *EthereumService) WatchEvent() error {
	events := make(chan *memcoin.MemcoinTaxCollected)
	sub, err := es.MemCoin.WatchTaxCollected(nil, events)
	if err != nil {
		log.Printf("Failed to watch event: %v", err)
		return fmt.Errorf("failed to watch event: %w", err)
	}
	defer sub.Unsubscribe()

	for event := range events {
		log.Printf("Tax collected: %s", event.Amount.String())
	}

	return nil
}
