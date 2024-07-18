package wireguard

import (
	"os"

	"github.com/sentinel-official/sentinel-go-sdk/v1/third_party/wireguard/windows/conf"
)

type ClientConfig struct {
	conf.Config
}

func (cc *ClientConfig) ToWgQuick() (string, error) {
	return cc.Config.ToWgQuick(), nil
}

func (cc *ClientConfig) WriteFile(filepath string) error {
	data, err := cc.ToWgQuick()
	if err != nil {
		return err
	}

	return os.WriteFile(filepath, []byte(data), 0644)
}

type ServerConfig struct {
	Addresses    []string `json:"addresses"`
	EnableIPv4   bool     `json:"enable_ipv4"`
	EnableIPv6   bool     `json:"enable_ipv6"`
	Interface    string   `json:"interface"`
	ListenPort   uint16   `json:"listen_port"`
	OutInterface string   `json:"out_interface"`
	PrivateKey   string   `json:"private_key"`
}

func (sc *ServerConfig) WriteFile(filepath string) error {
	data, err := sc.ToWgQuick()
	if err != nil {
		return err
	}

	return os.WriteFile(filepath, []byte(data), 0644)
}
