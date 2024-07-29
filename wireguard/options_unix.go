//go:build darwin || linux

package wireguard

import (
	"fmt"
	"strings"
)

// PostDown generates the PostDown rules based on IPv4 and IPv6 settings
func (so *ServerOptions) PostDown() string {
	var rules []string
	if so.EnableIPv4 {
		rules = append(rules, "iptables -D FORWARD -i %i -j ACCEPT")
		rules = append(rules, fmt.Sprintf("iptables -t nat -D POSTROUTING -o %s -j MASQUERADE", so.OutInterface))
	}
	if so.EnableIPv6 {
		rules = append(rules, "ip6tables -D FORWARD -i %i -j ACCEPT")
		rules = append(rules, fmt.Sprintf("ip6tables -t nat -D POSTROUTING -o %s -j MASQUERADE", so.OutInterface))
	}

	return strings.Join(rules, "; ")
}

// PostUp generates the PostUp rules based on IPv4 and IPv6 settings
func (so *ServerOptions) PostUp() string {
	var rules []string
	if so.EnableIPv4 {
		rules = append(rules, "iptables -A FORWARD -i %i -j ACCEPT")
		rules = append(rules, fmt.Sprintf("iptables -t nat -A POSTROUTING -o %s -j MASQUERADE", so.OutInterface))
	}
	if so.EnableIPv6 {
		rules = append(rules, "ip6tables -A FORWARD -i %i -j ACCEPT")
		rules = append(rules, fmt.Sprintf("ip6tables -t nat -A POSTROUTING -o %s -j MASQUERADE", so.OutInterface))
	}

	return strings.Join(rules, "; ")
}
