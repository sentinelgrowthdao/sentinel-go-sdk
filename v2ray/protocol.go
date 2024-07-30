package v2ray

import (
	"github.com/v2fly/v2ray-core/v5/common/serial"
	"github.com/v2fly/v2ray-core/v5/common/uuid"
	"github.com/v2fly/v2ray-core/v5/proxy/vless"
	"github.com/v2fly/v2ray-core/v5/proxy/vmess"
	"google.golang.org/protobuf/types/known/anypb"
)

// Protocol is a custom type used to represent different protocols.
type Protocol byte

// Constants for Protocol type with automatic incrementation for each protocol method.
const (
	ProtocolUnspecified Protocol = iota // Default value for unspecified protocol
	ProtocolVLess                       // ProtocolVLess represents the VLess protocol
	ProtocolVMess                       // ProtocolVMess represents the VMess protocol
)

// String returns a string representation of the Protocol type.
func (p Protocol) String() string {
	switch p {
	case ProtocolVLess:
		return "vless"
	case ProtocolVMess:
		return "vmess"
	default:
		return "" // Return empty string for unspecified or unknown protocols
	}
}

// IsValid checks if the Protocol value is valid.
func (p Protocol) IsValid() bool {
	return p.String() != ""
}

// Account generates an account message based on the Protocol type.
func (p Protocol) Account(uid uuid.UUID) *anypb.Any {
	switch p {
	case ProtocolVLess:
		return serial.ToTypedMessage(
			&vless.Account{
				Id: uid.String(),
			},
		)
	case ProtocolVMess:
		return serial.ToTypedMessage(
			&vmess.Account{
				Id: uid.String(),
			},
		)
	default:
		return nil
	}
}

// NewProtocolFromString converts a string to a Protocol type.
func NewProtocolFromString(v string) Protocol {
	switch v {
	case "vless":
		return ProtocolVLess
	case "vmess":
		return ProtocolVMess
	default:
		return ProtocolUnspecified // Returns the default protocol if no match is found
	}
}
