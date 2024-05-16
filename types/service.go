package types

import (
	"context"
)

// ServiceType represents the type of service as a byte.
type ServiceType byte

const (
	ServiceTypeUnspecified ServiceType = 0x00 + iota // ServiceTypeUnspecified represents an unspecified service type.
	ServiceTypeWireGuard                             // ServiceTypeWireGuard represents the WireGuard service type.
	ServiceTypeV2Ray                                 // ServiceTypeV2Ray represents the V2Ray service type.
)

// String returns the string representation of the ServiceType.
func (s ServiceType) String() string {
	switch s {
	case ServiceTypeWireGuard:
		return "wireguard"
	case ServiceTypeV2Ray:
		return "v2ray"
	default:
		return ""
	}
}

// ServiceTypeFromString converts a string to a ServiceType.
func ServiceTypeFromString(s string) ServiceType {
	switch s {
	case "wireguard":
		return ServiceTypeWireGuard
	case "v2ray":
		return ServiceTypeV2Ray
	default:
		return ServiceTypeUnspecified
	}
}

// PeerStatistic represents the download and upload statistics for a peer.
type PeerStatistic struct {
	Download int64  `json:"download"` // Download is the total download in bytes.
	Key      string `json:"key"`      // Key is the identifier for the peer.
	Upload   int64  `json:"upload"`   // Upload is the total upload in bytes.
}

// ClientService defines the interface for client-side service operations.
type ClientService interface {
	Down() error                       // Down brings down the client service.
	IsUp() (bool, error)               // IsUp checks if the client service is up.
	PostDown() error                   // PostDown performs operations after the service is brought down.
	PostUp() error                     // PostUp performs operations after the service is brought up.
	PreDown() error                    // PreDown performs operations before the service is brought down.
	PreUp(interface{}) error           // PreUp performs operations before the service is brought up.
	Statistics() (int64, int64, error) // Statistics returns the download and upload statistics.
	Up() error                         // Up brings up the client service.
}

// ServerService defines the interface for server-side service operations.
type ServerService interface {
	AddPeer(context.Context, []byte) ([]byte, error)          // AddPeer adds a peer to the server service.
	HasPeer(context.Context, []byte) (bool, error)            // HasPeer checks if a peer exists in the server service.
	Info() []byte                                             // Info returns the information of the server service.
	Init() error                                              // Init initializes the server service.
	PeerCount() int                                           // PeerCount returns the count of peers.
	PeerStatistics(context.Context) ([]*PeerStatistic, error) // PeerStatistics returns the statistics for all peers.
	RemovePeer(context.Context, []byte) error                 // RemovePeer removes a peer from the server service.
	Start() error                                             // Start starts the server service.
	Stop() error                                              // Stop stops the server service.
	Type() ServiceType                                        // Type returns the type of the server service.
}
