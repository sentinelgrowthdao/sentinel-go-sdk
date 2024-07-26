package wireguard

import (
	"net/netip"
	"strings"

	"github.com/sentinel-official/sentinel-go-sdk/v1/third_party/wireguard/windows/conf"
)

// ToWgQuick converts the ServerOptions to a WireGuard quick configuration string.
func (so *ServerOptions) ToWgQuick() (string, error) {
	privateKey, err := conf.NewPrivateKeyFromString(so.PrivateKey)
	if err != nil {
		return "", err
	}

	var (
		addresses []netip.Prefix
		postUp    []string
		postDown  []string
	)

	for _, item := range so.Addresses {
		address, err := netip.ParsePrefix(item)
		if err != nil {
			return "", err
		}

		addresses = append(addresses, address)
	}

	cfg := &conf.Config{
		Name: so.Interface,
		Interface: conf.Interface{
			PrivateKey: *privateKey,
			Addresses:  addresses,
			ListenPort: so.ListenPort,
			PostUp:     strings.Join(postUp, " "),
			PostDown:   strings.Join(postDown, " "),
		},
	}

	return cfg.ToWgQuick(), nil
}
