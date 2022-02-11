package ethereum

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

type EthereumNetworkConfig struct {
	KeystorePath string `yaml:"keystore_path"`
	Endpoint     string `yaml:"endpoint"`
	Password     string `yaml:"pass_admin"`
	ChainID      string `yaml:"chain_id"`
}

func NewEthereumNetworkConfig() {

}

func NewClient(cfg *EthereumNetworkConfig) (*ethclient.Client, error) {
	client, err := ethclient.Dial(cfg.Endpoint)
	if err != nil {
		return nil, err
	}
	log.Println("we have a connection to ethereum network")
	return client, nil
}
