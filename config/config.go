package config

import (
	"digitalsignature/internal/pkg/ethereum"
	"fmt"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func NewConfig(path, stage string) (config *Configuration, err error) {
	conf := Configuration{}
	name := fmt.Sprintf("config.%s", stage)

	vp := viper.New()
	vp.AddConfigPath(path)
	vp.SetConfigName(name)
	vp.AutomaticEnv()
	vp.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := vp.ReadInConfig(); err != nil {
		return nil, err
	}
	if err := vp.Unmarshal(&conf); err != nil {
		log.Println("Failed to unmarshal config: ", err)
		return nil, err
	}
	setGinMode("debug")
	log.Printf("Config: %+v", conf)
	return &conf, nil
}

// Configuration holds data necessery for configuring application
type Configuration struct {
	Ethereum        *ethereum.EthereumNetworkConfig `mapstructure:"ethereum"`
	ContractAddress *ethereum.ContractAddress       `mapstructure:"address"`
}

func setGinMode(mode string) {
	switch mode {
	case "production":
		gin.SetMode(gin.ReleaseMode)
	case "test":
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.DebugMode)
	}
}
