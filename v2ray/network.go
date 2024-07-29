package v2ray

// Network is a custom type used to represent different network protocols.
type Network byte

// Constants for Network type with automatic incrementation for each network method.
const (
	NetworkUnspecified  Network = iota // Default value for unspecified network
	NetworkDomainSocket                // NetworkDomainSocket represents a UNIX domain socket
	NetworkGUN                         // NetworkGUN represents the GUN protocol
	NetworkGRPC                        // NetworkGRPC represents gRPC, a high-performance RPC framework
	NetworkHTTP                        // NetworkHTTP represents the HTTP protocol
	NetworkMKCP                        // NetworkMKCP represents the MKCP (modified KCP) protocol
	NetworkQUIC                        // NetworkQUIC represents the QUIC protocol
	NetworkTCP                         // NetworkTCP represents the TCP network protocol
	NetworkWebSocket                   // NetworkWebSocket represents the WebSocket protocol
)

// String returns a string representation of the Network type.
func (t Network) String() string {
	switch t {
	case NetworkDomainSocket:
		return "domainsocket"
	case NetworkGUN:
		return "gun"
	case NetworkGRPC:
		return "grpc"
	case NetworkHTTP:
		return "http"
	case NetworkMKCP:
		return "mkcp"
	case NetworkQUIC:
		return "quic"
	case NetworkTCP:
		return "tcp"
	case NetworkWebSocket:
		return "websocket"
	default:
		return "" // Return empty string for unspecified or unknown network types
	}
}

// IsValid checks if the Network value is valid.
func (t Network) IsValid() bool {
	return t.String() != ""
}

// NewNetworkFromString converts a string to a Network type.
func NewNetworkFromString(v string) Network {
	switch v {
	case "domainsocket":
		return NetworkDomainSocket
	case "gun":
		return NetworkGUN
	case "grpc":
		return NetworkGRPC
	case "http":
		return NetworkHTTP
	case "mkcp":
		return NetworkMKCP
	case "quic":
		return NetworkQUIC
	case "tcp":
		return NetworkTCP
	case "websocket":
		return NetworkWebSocket
	default:
		return NetworkUnspecified // Returns the default network type if no match is found
	}
}
