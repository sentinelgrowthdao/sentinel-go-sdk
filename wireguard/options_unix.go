//go:build darwin || linux

package wireguard

import (
	"fmt"
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

	if so.EnableIPv4 {
		postUp = append(postUp, "iptables -A FORWARD -i %i -j ACCEPT;")
		postUp = append(postUp, fmt.Sprintf("iptables -t nat -A POSTROUTING -o %s -j MASQUERADE;", so.OutInterface))

		postDown = append(postDown, "iptables -D FORWARD -i %i -j ACCEPT;")
		postDown = append(postDown, fmt.Sprintf("iptables -t nat -D POSTROUTING -o %s -j MASQUERADE;", so.OutInterface))
	}
	if so.EnableIPv6 {
		postUp = append(postUp, "ip6tables -A FORWARD -i %i -j ACCEPT;")
		postUp = append(postUp, fmt.Sprintf("ip6tables -t nat -A POSTROUTING -o %s -j MASQUERADE;", so.OutInterface))

		postDown = append(postDown, "ip6tables -D FORWARD -i %i -j ACCEPT;")
		postDown = append(postDown, fmt.Sprintf("ip6tables -t nat -D POSTROUTING -o %s -j MASQUERADE;", so.OutInterface))
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