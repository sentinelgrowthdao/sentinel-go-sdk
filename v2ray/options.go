package v2ray

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/pelletier/go-toml/v2"
	"github.com/spf13/cobra"
)

// ClientOptions represents the V2Ray client configuration options.
type ClientOptions struct{}

// WriteToFile writes the ClientOptions configuration to a TOML file.
func (co *ClientOptions) WriteToFile(filepath string) error {
	data, err := toml.Marshal(co)
	if err != nil {
		return err
	}

	return os.WriteFile(filepath, data, 0644)
}

// WriteConfigToFile writes the ClientOptions configuration to a file in JSON format.
func (co *ClientOptions) WriteConfigToFile(filepath string) error {
	data, err := co.ToConfig()
	if err != nil {
		return err
	}

	return os.WriteFile(filepath, []byte(data), 0644)
}

// NewClientOptionsFromFile reads the configuration from a TOML file and unmarshals it into a ClientOptions instance.
func NewClientOptionsFromFile(filepath string) (*ClientOptions, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var co ClientOptions
	if err := toml.Unmarshal(data, &co); err != nil {
		return nil, err
	}

	return &co, nil
}

// InboundServerOptions represents the V2Ray inbound server configuration options.
type InboundServerOptions struct {
	Network     string `json:"network"`
	Port        uint16 `json:"port"`
	Protocol    string `json:"protocol"`
	Security    string `json:"security"`
	TLSCertPath string `json:"tls_cert_path"`
	TLSKeyPath  string `json:"tls_key_path"`
}

// WithNetwork sets the Network field.
func (so *InboundServerOptions) WithNetwork(network string) *InboundServerOptions {
	so.Network = network
	return so
}

// WithPort sets the Port field.
func (so *InboundServerOptions) WithPort(port uint16) *InboundServerOptions {
	so.Port = port
	return so
}

// WithProtocol sets the Protocol field.
func (so *InboundServerOptions) WithProtocol(protocol string) *InboundServerOptions {
	so.Protocol = protocol
	return so
}

// WithSecurity sets the Security field.
func (so *InboundServerOptions) WithSecurity(security string) *InboundServerOptions {
	so.Security = security
	return so
}

// WithTLSCertPath sets the TLSCertPath field.
func (so *InboundServerOptions) WithTLSCertPath(certPath string) *InboundServerOptions {
	so.TLSCertPath = certPath
	return so
}

// WithTLSKeyPath sets the TLSKeyPath field.
func (so *InboundServerOptions) WithTLSKeyPath(keyPath string) *InboundServerOptions {
	so.TLSKeyPath = keyPath
	return so
}

// Tag creates a Tag instance based on the InboundServerOptions configuration.
func (so *InboundServerOptions) Tag() *Tag {
	protocol := NewProtocolFromString(so.Protocol)
	network := NewNetworkFromString(so.Network)
	security := NewSecurityFromString(so.Security)

	return &Tag{
		p: protocol,
		n: network,
		s: security,
	}
}

// Validate validates the InboundServerOptions fields.
func (so *InboundServerOptions) Validate() error {
	network := NewNetworkFromString(so.Network)
	if !network.IsValid() {
		return fmt.Errorf("invalid network value: %s", so.Network)
	}

	protocol := NewProtocolFromString(so.Protocol)
	if !protocol.IsValid() {
		return fmt.Errorf("invalid protocol value: %s", so.Protocol)
	}

	security := NewSecurityFromString(so.Security)
	if !security.IsValid() {
		return fmt.Errorf("invalid security value: %s", so.Security)
	}

	if security == SecurityTLS {
		if so.TLSCertPath == "" || so.TLSKeyPath == "" {
			return errors.New("TLS cert path and key path cannot be empty when security is 'tls'")
		}
	}

	return nil
}

// ServerOptions represents the V2Ray server configuration options.
type ServerOptions struct {
	Inbounds []*InboundServerOptions `json:"inbounds"`
}

// WithInbounds sets the Inbounds field.
func (so *ServerOptions) WithInbounds(inbounds ...*InboundServerOptions) *ServerOptions {
	so.Inbounds = inbounds
	return so
}

// Validate validates the ServerOptions fields.
func (so *ServerOptions) Validate() error {
	portSet := make(map[uint16]bool)
	tagSet := make(map[string]bool)

	for _, inbound := range so.Inbounds {
		if err := inbound.Validate(); err != nil {
			return err
		}

		if inbound.Port <= 1024 {
			return fmt.Errorf("port must be greater than 1024, got: %d", inbound.Port)
		}
		if portSet[inbound.Port] {
			return fmt.Errorf("port collision detected for port: %d", inbound.Port)
		}
		portSet[inbound.Port] = true

		tag := inbound.Tag().String()
		if tagSet[tag] {
			return fmt.Errorf("duplicate tag detected: %s", tag)
		}
		tagSet[tag] = true
	}

	return nil
}

// WriteToFile writes the ServerOptions configuration to a TOML file.
func (so *ServerOptions) WriteToFile(filepath string) error {
	data, err := toml.Marshal(so)
	if err != nil {
		return err
	}

	return os.WriteFile(filepath, data, 0644)
}

// WriteConfigToFile writes the ServerOptions configuration to a file in JSON format.
func (so *ServerOptions) WriteConfigToFile(filepath string) error {
	data, err := so.ToConfig()
	if err != nil {
		return err
	}

	return os.WriteFile(filepath, []byte(data), 0644)
}

// AddServerFlagsToCmd adds server-related flags to the given Cobra command.
func AddServerFlagsToCmd(cmd *cobra.Command, prefix string) {
	if prefix != "" {
		prefix = prefix + "."
	}

	cmd.Flags().StringSlice(prefix+"network", []string{}, "Comma-separated list of network types for inbound servers.")
	cmd.Flags().StringSlice(prefix+"port", []string{}, "Comma-separated list of ports for inbound servers.")
	cmd.Flags().StringSlice(prefix+"protocol", []string{}, "Comma-separated list of protocols for inbound servers.")
	cmd.Flags().StringSlice(prefix+"security", []string{}, "Comma-separated list of security settings for inbound servers.")
	cmd.Flags().StringSlice(prefix+"tls-cert-path", []string{}, "Comma-separated list of TLS certificate paths for inbound servers.")
	cmd.Flags().StringSlice(prefix+"tls-key-path", []string{}, "Comma-separated list of TLS certificate keys for inbound servers.")
}

// NewServerOptionsFromCmd creates and returns a ServerOptions instance from the given Cobra command's flags.
func NewServerOptionsFromCmd(cmd *cobra.Command, prefix string) (*ServerOptions, error) {
	if prefix != "" {
		prefix = prefix + "."
	}

	networkList, err := cmd.Flags().GetStringSlice(prefix + "network")
	if err != nil {
		return nil, err
	}

	portList, err := cmd.Flags().GetStringSlice(prefix + "port")
	if err != nil {
		return nil, err
	}

	protocolList, err := cmd.Flags().GetStringSlice(prefix + "protocol")
	if err != nil {
		return nil, err
	}

	securityList, err := cmd.Flags().GetStringSlice(prefix + "security")
	if err != nil {
		return nil, err
	}

	tlsCertPathList, err := cmd.Flags().GetStringSlice(prefix + "tls-cert-path")
	if err != nil {
		return nil, err
	}

	tlsKeyPathList, err := cmd.Flags().GetStringSlice(prefix + "tls-key-path")
	if err != nil {
		return nil, err
	}

	if len(networkList) != len(portList) || len(networkList) != len(protocolList) || len(networkList) != len(securityList) ||
		len(networkList) != len(tlsCertPathList) || len(networkList) != len(tlsKeyPathList) {
		return nil, errors.New("all inbound server flags must have the same number of values")
	}

	var inbounds []*InboundServerOptions
	for i := range networkList {
		port, err := strconv.ParseUint(portList[i], 10, 16)
		if err != nil {
			return nil, fmt.Errorf("invalid port: %w", err)
		}

		inbounds = append(inbounds, &InboundServerOptions{
			Network:     networkList[i],
			Port:        uint16(port),
			Protocol:    protocolList[i],
			Security:    securityList[i],
			TLSCertPath: tlsCertPathList[i],
			TLSKeyPath:  tlsKeyPathList[i],
		})
	}

	return &ServerOptions{
		Inbounds: inbounds,
	}, nil
}

// NewServerOptionsFromFile reads the configuration from a TOML file and unmarshals it into a ServerOptions instance.
func NewServerOptionsFromFile(filepath string) (*ServerOptions, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var so ServerOptions
	if err := toml.Unmarshal(data, &so); err != nil {
		return nil, err
	}

	return &so, nil
}
