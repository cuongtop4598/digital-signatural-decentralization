package ipfs

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"

	shell "github.com/ipfs/go-ipfs-api"
	"go.uber.org/zap"
)

type IpfsService struct {
	log   *zap.Logger
	shell *shell.Shell
}

type IPFSConfig struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

func NewIpfsService(c IPFSConfig, l *zap.Logger) *IpfsService {
	log.Println("ipfs API config:", c)
	sh := shell.NewShell(c.Host + ":" + c.Port)
	return &IpfsService{
		log:   l,
		shell: sh,
	}
}

func (s *IpfsService) AddFile(data []byte) (string, error) {
	reader := bytes.NewReader(data)
	cid, err := s.shell.Add(reader)
	if err != nil {
		s.log.Error("Add file error :", zap.String("detail", err.Error()))
		return "", err
	}
	s.log.Info("Added %s ", zap.String("cid", cid))
	return cid, nil
}

func (s *IpfsService) GetFile(cid string, fileName string) ([]byte, error) {
	os.Chdir(".")
	err := s.shell.Get(cid, "./assets/documents/"+fileName)
	if err != nil {
		s.log.Error("Get file error :", zap.String("detail", err.Error()))
	}
	f, err := os.Open("./assets/documents/" + fileName)
	if err != nil {
		s.log.Error("Read save file error :", zap.String("detail", err.Error()))
	}
	return ioutil.ReadAll(f)
}
