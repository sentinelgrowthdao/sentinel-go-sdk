package v2ray

// Transport is a custom type derived from byte, used to represent different transport protocols.
type Transport byte

// Constants for Transport type with automatic incrementation for each transport method.
const (
	TransportUnspecified  Transport = 0x00 + iota // starts iota at 0x00, default value for unspecified transport
	TransportTCP                                  // TransportTCP represents the TCP transport protocol
	TransportMKCP                                 // TransportMKCP represents the MKCP (modified KCP) protocol
	TransportWebSocket                            // TransportWebSocket represents the WebSocket protocol
	TransportHTTP                                 // TransportHTTP represents the HTTP protocol
	TransportDomainSocket                         // TransportDomainSocket represents a UNIX domain socket
	TransportQUIC                                 // TransportQUIC represents the QUIC protocol
	TransportGUN                                  // TransportGUN represents the GUN protocol, specific usage not standard
	TransportGRPC                                 // TransportGRPC represents gRPC, a high-performance RPC framework
)

// String returns a string representation of the Transport type.
func (t Transport) String() string {
	switch t {
	case TransportTCP:
		return "tcp"
	case TransportMKCP:
		return "mkcp"
	case TransportWebSocket:
		return "websocket"
	case TransportHTTP:
		return "http"
	case TransportDomainSocket:
		return "domainsocket"
	case TransportQUIC:
		return "quic"
	case TransportGUN:
		return "gun"
	case TransportGRPC:
		return "grpc"
	default:
		return "" // Return empty string for unspecified or unknown transport types
	}
}

// NewTransportFromString converts a string to a Transport type.
// This is often used for configuration parsing or interfacing with user inputs.
func NewTransportFromString(v string) Transport {
	switch v {
	case "tcp":
		return TransportTCP
	case "mkcp":
		return TransportMKCP
	case "websocket":
		return TransportWebSocket
	case "http":
		return TransportHTTP
	case "domainsocket":
		return TransportDomainSocket
	case "quic":
		return TransportQUIC
	case "gun":
		return TransportGUN
	case "grpc":
		return TransportGRPC
	default:
		return TransportUnspecified // Returns the default transport type if no match is found
	}
}
