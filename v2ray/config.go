package v2ray

import (
	"os"

	"github.com/sentinel-official/sentinel-go-sdk/v1/utils"
)

type ClientConfig struct {
}

func (c *ClientConfig) WriteFile(filepath string) error {
	return os.WriteFile(filepath, utils.MustMarshalJSON(c), 0644)
}

type ServerConfig struct {
}

func (c *ServerConfig) WriteFile(filepath string) error {
	return os.WriteFile(filepath, utils.MustMarshalJSON(c), 0644)
}
