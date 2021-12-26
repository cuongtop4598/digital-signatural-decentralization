package config

import (
	"digitalsignature/internal/pkg/ethereum"
	"log"
	"path/filepath"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// Load returns Configuration struct
func Load(env string) *Configuration {
	_, filePath, _, _ := runtime.Caller(0)
	configName := "config." + env + ".yaml"
	configPath := filePath[:len(filePath)-14] + "files" + string(filepath.Separator)
	viper.SetConfigName(configName)
	viper.AddConfigPath(configPath)
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	var config Configuration
	viper.Unmarshal(&config)
	setGinMode(config.Server.Mode)

	return &config
}

// Configuration holds data necessery for configuring application
type Configuration struct {
	Server   *Server                 `yaml:"server"`
	Ethereum *ethereum.NetworkConfig `yaml:"ethereum"`
}

// Server holds data necessary for server configuration
type Server struct {
	Mode string `yaml:"mode"`
	Host string `yaml:"host"`
	Port string `yaml:"port"`
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
