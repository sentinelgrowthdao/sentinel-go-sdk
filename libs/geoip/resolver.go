package geoip

import (
	"time"
)

// Location represents geographical location information associated with an IP address.
type Location struct {
	City      string  `json:"city,omitempty"`      // City where the IP address is located.
	Country   string  `json:"country,omitempty"`   // Country where the IP address is located.
	IP        string  `json:"ip,omitempty"`        // IP address that was resolved.
	Latitude  float64 `json:"latitude,omitempty"`  // Latitude of the location.
	Longitude float64 `json:"longitude,omitempty"` // Longitude of the location.
}

// Resolver is an interface for resolving IP addresses into location data.
type Resolver interface {
	Resolve(ip string) (*Location, error)
}

// NewDefaultResolver creates a new default Resolver instance using the default IPAPIClient.
func NewDefaultResolver() Resolver {
	return NewIPAPIClient(15 * time.Second)
}
