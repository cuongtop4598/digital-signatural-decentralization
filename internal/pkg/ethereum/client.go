package ethereum

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

type NetworkConfig struct {
	Endpoint string `yaml:"endpoint"`
}

func NewClient(cfg *NetworkConfig) (*ethclient.Client, error) {
	client, err := ethclient.Dial(cfg.Endpoint)
	if err != nil {
		return nil, err
	}
	log.Println("we have a connection to ethereum network")
	return client, nil
}
