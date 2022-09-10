package ethereum

import (
	"context"
	"log"

	"github.com/caarlos0/env/v6"
	"github.com/ethereum/go-ethereum/ethclient"
)

func NewClient() *ethclient.Client {
	config := getNetworkConfig()
	retry := 5
RETRY:
	client, err := ethclient.DialContext(context.Background(), config.Url)
	if err != nil {
		log.Fatal(err)
	}
	if client == nil {
		if retry > 0 {
			retry--
			goto RETRY
		}
	}
	return client
}

type Network struct {
	Name    string `env:"NAME"`
	Url     string `env:"CHAIN_ENDPOINT"`
	ChainID int    `env:"CHAIN_ID"`
}

func getNetworkConfig() *Network {
	n := Network{}
	if err := env.Parse(&n); err != nil {
		log.Fatal(err)
		return nil
	}
	return &n
}
