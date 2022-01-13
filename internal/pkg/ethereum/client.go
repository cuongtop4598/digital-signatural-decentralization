package ethereum

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

type NetworkConfig struct {
	Wallets  string `yaml:"wallets"`
	Endpoint string `yaml:"endpoint"`
	Password string `yaml:"passadmin"`
}

func NewClient(cfg *NetworkConfig) (*ethclient.Client, error) {
	client, err := ethclient.Dial(cfg.Endpoint)
	if err != nil {
		return nil, err
	}
	log.Println("we have a connection to ethereum network")
	return client, nil
}