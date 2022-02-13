package ethereum

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

type EthereumNetworkConfig struct {
	KeystorePath string `mapstructure:"keystore_path"`
	Endpoint     string `mapstructure:"endpoint"`
	Password     string `mapstructure:"pass_admin"`
	ChainID      int64  `mapstructure:"chain_id"`
}

type ContractAddress struct {
	Document string   `mapstructure:"document"`
	Accounts []string `mapstructure:"accounts"`
}

func NewClient(cfg *EthereumNetworkConfig) (*ethclient.Client, error) {
	client, err := ethclient.Dial(cfg.Endpoint)
	if err != nil {
		return nil, err
	}
	log.Println("we have a connection to ethereum network")
	return client, nil
}
