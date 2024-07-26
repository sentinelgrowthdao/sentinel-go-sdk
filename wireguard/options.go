package wireguard

import (
	"errors"
	"fmt"
	"net/netip"
	"os"

	"github.com/pelletier/go-toml/v2"
	"github.com/spf13/cobra"

	"github.com/sentinel-official/sentinel-go-sdk/v1/third_party/wireguard/windows/conf"
)

// ClientOptions represents the WireGuard client configuration options.
type ClientOptions struct {
	conf.Config
}

// ToWgQuick converts the ClientOptions to a WireGuard quick configuration string.
func (co *ClientOptions) ToWgQuick() (string, error) {
	return co.Config.ToWgQuick(), nil
}

// WriteToFile writes the ClientOptions configuration to a TOML file.
func (co *ClientOptions) WriteToFile(filepath string) error {
	data, err := toml.Marshal(co)
	if err != nil {
		return err
	}

	return os.WriteFile(filepath, data, 0644)
}

// WriteConfigToFile writes the WireGuard configuration to a file in a format recognized by WireGuard.
func (co *ClientOptions) WriteConfigToFile(filepath string) error {
	data, err := co.ToWgQuick()
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

// ServerOptions represents the WireGuard server configuration options.
type ServerOptions struct {
	Addresses    []string `json:"addresses"`
	EnableIPv4   bool     `json:"enable_ipv4"`
	EnableIPv6   bool     `json:"enable_ipv6"`
	Interface    string   `json:"interface"`
	ListenPort   uint16   `json:"listen_port"`
	OutInterface string   `json:"out_interface"`
	PrivateKey   string   `json:"private_key"`
}

// WithAddresses sets the Addresses field and returns the modified ServerOptions instance.
func (so *ServerOptions) WithAddresses(v ...string) *ServerOptions {
	so.Addresses = v
	return so
}

// WithEnableIPv4 sets the EnableIPv4 field and returns the modified ServerOptions instance.
func (so *ServerOptions) WithEnableIPv4(v bool) *ServerOptions {
	so.EnableIPv4 = v
	return so
}

// WithEnableIPv6 sets the EnableIPv6 field and returns the modified ServerOptions instance.
func (so *ServerOptions) WithEnableIPv6(v bool) *ServerOptions {
	so.EnableIPv6 = v
	return so
}

// WithInterface sets the Interface field and returns the modified ServerOptions instance.
func (so *ServerOptions) WithInterface(v string) *ServerOptions {
	so.Interface = v
	return so
}

// WithListenPort sets the ListenPort field and returns the modified ServerOptions instance.
func (so *ServerOptions) WithListenPort(v uint16) *ServerOptions {
	so.ListenPort = v
	return so
}

// WithOutInterface sets the OutInterface field and returns the modified ServerOptions instance.
func (so *ServerOptions) WithOutInterface(v string) *ServerOptions {
	so.OutInterface = v
	return so
}

// WithPrivateKey sets the PrivateKey field and returns the modified ServerOptions instance.
func (so *ServerOptions) WithPrivateKey(v string) *ServerOptions {
	so.PrivateKey = v
	return so
}

// WriteToFile writes the ServerOptions configuration to a TOML file.
func (so *ServerOptions) WriteToFile(filepath string) error {
	data, err := toml.Marshal(so)
	if err != nil {
		return err
	}

	return os.WriteFile(filepath, data, 0644)
}

// WriteConfigToFile writes the WireGuard server configuration to a file in a format recognized by WireGuard.
func (so *ServerOptions) WriteConfigToFile(filepath string) error {
	data, err := so.ToWgQuick()
	if err != nil {
		return err
	}

	return os.WriteFile(filepath, []byte(data), 0644)
}

// Validate checks that the ServerOptions fields have valid values.
func (so *ServerOptions) Validate() error {
	if len(so.Addresses) == 0 {
		return errors.New("addresses cannot be empty")
	}
	for _, item := range so.Addresses {
		_, err := netip.ParsePrefix(item)
		if err != nil {
			return fmt.Errorf("invalid address: %w", err)
		}
	}

	if so.Interface == "" {
		return errors.New("interface cannot be empty")
	}
	if so.ListenPort == 0 {
		return errors.New("listen_port cannot be zero")
	}
	if so.OutInterface == "" {
		return errors.New("out_interface cannot be empty")
	}

	_, err := conf.NewPrivateKeyFromString(so.PrivateKey)
	if err != nil {
		return fmt.Errorf("invalid private_key: %w", err)
	}

	return nil
}

// AddServerFlagsToCmd adds server-related flags to the given Cobra command.
func AddServerFlagsToCmd(cmd *cobra.Command, prefix string) {
	if prefix != "" {
		prefix = prefix + "."
	}

	cmd.Flags().StringSlice(prefix+"addresses", nil, "Comma-separated list of addresses for the server.")
	cmd.Flags().Bool(prefix+"enable-ipv4", false, "Enable IPv4 for the server.")
	cmd.Flags().Bool(prefix+"enable-ipv6", false, "Enable IPv6 for the server.")
	cmd.Flags().String(prefix+"interface", "", "Network interface for the server.")
	cmd.Flags().Uint16(prefix+"listen-port", 0, "Listen port for the server.")
	cmd.Flags().String(prefix+"out-interface", "", "Outgoing network interface for the server.")
	cmd.Flags().String(prefix+"private-key", "", "Private key for the server.")
}

// NewServerOptionsFromCmd creates and returns a ServerOptions instance from the given Cobra command's flags.
func NewServerOptionsFromCmd(cmd *cobra.Command, prefix string) (*ServerOptions, error) {
	if prefix != "" {
		prefix = prefix + "."
	}

	addresses, err := cmd.Flags().GetStringSlice(prefix + "addresses")
	if err != nil {
		return nil, err
	}

	enableIPv4, err := cmd.Flags().GetBool(prefix + "enable-ipv4")
	if err != nil {
		return nil, err
	}

	enableIPv6, err := cmd.Flags().GetBool(prefix + "enable-ipv6")
	if err != nil {
		return nil, err
	}

	iface, err := cmd.Flags().GetString(prefix + "interface")
	if err != nil {
		return nil, err
	}

	listenPort, err := cmd.Flags().GetUint16(prefix + "listen-port")
	if err != nil {
		return nil, err
	}

	outIface, err := cmd.Flags().GetString(prefix + "out-interface")
	if err != nil {
		return nil, err
	}

	privateKey, err := cmd.Flags().GetString(prefix + "private-key")
	if err != nil {
		return nil, err
	}

	return &ServerOptions{
		Addresses:    addresses,
		EnableIPv4:   enableIPv4,
		EnableIPv6:   enableIPv6,
		Interface:    iface,
		ListenPort:   listenPort,
		OutInterface: outIface,
		PrivateKey:   privateKey,
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
