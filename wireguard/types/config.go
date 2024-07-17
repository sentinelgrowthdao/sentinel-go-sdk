package types

import (
	"os"

	"golang.zx2c4.com/wireguard/windows/conf"
)

type ClientConfig struct {
	conf.Config
}

func (c *ClientConfig) WriteFile(filepath string) error {
	return os.WriteFile(filepath, []byte(c.ToWgQuick()), 0644)
}

type ServerConfig struct {
	conf.Config
}

func (c *ServerConfig) WriteFile(filepath string) error {
	return os.WriteFile(filepath, []byte(c.ToWgQuick()), 0644)
}
