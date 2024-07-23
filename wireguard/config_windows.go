package wireguard

import (
	"net/netip"
	"strings"

	"github.com/sentinel-official/sentinel-go-sdk/v1/third_party/wireguard/windows/conf"
)

func (sc *ServerConfig) ToWgQuick() (string, error) {
	privateKey, err := conf.NewPrivateKeyFromString(sc.PrivateKey)
	if err != nil {
		return "", err
	}

	var (
		addresses []netip.Prefix
		postUp    []string
		postDown  []string
	)

	for _, item := range sc.Addresses {
		address, err := netip.ParsePrefix(item)
		if err != nil {
			return "", err
		}

		addresses = append(addresses, address)
	}

	cfg := &conf.Config{
		Name: sc.Interface,
		Interface: conf.Interface{
			PrivateKey: *privateKey,
			Addresses:  addresses,
			ListenPort: sc.ListenPort,
			PostUp:     strings.Join(postUp, " "),
			PostDown:   strings.Join(postDown, " "),
		},
	}

	return cfg.ToWgQuick(), nil
}
