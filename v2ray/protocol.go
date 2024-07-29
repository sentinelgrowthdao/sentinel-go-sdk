package v2ray

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
