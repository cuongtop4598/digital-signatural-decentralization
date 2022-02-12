package config

import (
	"digitalsignature/internal/pkg/database"
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
	// binding
	if err := vp.Unmarshal(&conf); err != nil {
		log.Println("Failed to unmarshal config: ", err)
		return nil, err
	}
	setGinMode(conf.Server.Mode)
	log.Printf("Config: %+v", conf)
	return &conf, nil
}

// Configuration holds data necessery for configuring application
type Configuration struct {
	Server          *Server                         `mapstructure:"server"`
	Ethereum        *ethereum.EthereumNetworkConfig `mapstructure:"ethereum"`
	Database        *database.DBConfig              `mapstructure:"database"`
	ContractAddress string                          `mapstructure:"contract"`
}

// Server holds data necessary for server configuration
type Server struct {
	Mode string `mapstructure:"mode"`
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
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
