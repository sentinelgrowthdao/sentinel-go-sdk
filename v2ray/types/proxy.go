package types

import (
	"github.com/v2fly/v2ray-core/v5/common/protocol"
	"github.com/v2fly/v2ray-core/v5/common/serial"
	"github.com/v2fly/v2ray-core/v5/common/uuid"
	"github.com/v2fly/v2ray-core/v5/proxy/vmess"
	"google.golang.org/protobuf/types/known/anypb"
)

// Proxy is a custom type derived from byte to represent different proxy protocols.
type Proxy byte

// Constants for Proxy type, automatically incremented for each proxy type.
const (
	ProxyUnspecified Proxy = 0x00 + iota // ProxyUnspecified represents an unspecified proxy type.
	ProxyVMess                           // ProxyVMess represents the VMess proxy protocol
)

// String converts a Proxy type into its string representation.
func (p Proxy) String() string {
	switch p {
	case ProxyVMess:
		return "vmess"
	default:
		return "" // Return an empty string for unspecified proxy types
	}
}

// Tag returns a user-friendly string name of the Proxy, used for display purposes.
func (p Proxy) Tag() string {
	return p.String()
}

// Account returns a protocol buffer (Any type) containing the account information based on the proxy type.
// It serializes account data specific to the proxy protocol used.
func (p Proxy) Account(uid uuid.UUID) *anypb.Any {
	switch p {
	case ProxyVMess:
		// For VMess protocol, serialize the VMess account data with ID and default security settings.
		return serial.ToTypedMessage(
			&vmess.Account{
				AlterId: 0,
				Id:      uid.String(),
				SecuritySettings: &protocol.SecurityConfig{
					Type: protocol.SecurityType_AUTO,
				},
				TestsEnabled: "",
			},
		)
	default:
		return nil // Return nil for unspecified or unsupported proxy types
	}
}

// ProxyFromString converts a string representation of a proxy protocol to its corresponding Proxy type.
func ProxyFromString(s string) Proxy {
	switch s {
	case "vmess":
		return ProxyVMess
	default:
		return ProxyUnspecified // Default to ProxyUnspecified for unrecognized strings
	}
}
