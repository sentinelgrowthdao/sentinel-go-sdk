package v2ray

import (
	"errors"
	"fmt"
	"os"

	"github.com/pelletier/go-toml/v2"
	"github.com/spf13/cobra"

	"github.com/sentinel-official/sentinel-go-sdk/v1/utils"
)

// ClientOptions represents the V2Ray client configuration options.
type ClientOptions struct {
}

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
	return os.WriteFile(filepath, utils.MustMarshalJSON(co), 0644)
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

// VMessServerOptions represents the V2Ray VMess server configuration options.
type VMessServerOptions struct {
	EnableTLS  bool   `json:"enable_tls"`
	ListenPort uint16 `json:"listen_port"`
	Transport  string `json:"transport"`
}

// WithEnableTLS sets the EnableTLS field and returns the modified VMessServerOptions instance.
func (so *VMessServerOptions) WithEnableTLS(v bool) *VMessServerOptions {
	so.EnableTLS = v
	return so
}

// WithListenPort sets the ListenPort field and returns the modified VMessServerOptions instance.
func (so *VMessServerOptions) WithListenPort(v uint16) *VMessServerOptions {
	so.ListenPort = v
	return so
}

// WithTransport sets the Transport field and returns the modified VMessServerOptions instance.
func (so *VMessServerOptions) WithTransport(v string) *VMessServerOptions {
	so.Transport = v
	return so
}

// Validate validates the VMessServerOptions fields.
func (so *VMessServerOptions) Validate() error {
	if so.ListenPort == 0 {
		return errors.New("listen_port cannot be zero")
	}
	if so.Transport == "" {
		return errors.New("transport cannot be empty")
	}

	t := NewTransportFromString(so.Transport)
	if !t.IsValid() {
		return errors.New("invalid transport")
	}

	return nil
}

// AddVMessServerFlagsToCmd adds VMess server-related flags to the given Cobra command.
func AddVMessServerFlagsToCmd(cmd *cobra.Command, prefix string) {
	if prefix != "" {
		prefix = prefix + "."
	}

	cmd.Flags().Bool(prefix+"enable-tls", false, "Enable TLS for the VMess server.")
	cmd.Flags().Uint16(prefix+"listen-port", 0, "Listen port for the VMess server.")
	cmd.Flags().String(prefix+"transport", "", "Transport protocol for the VMess server.")
}

// NewVMessServerOptionsFromCmd creates and returns a VMessServerOptions instance from the given Cobra command's flags.
func NewVMessServerOptionsFromCmd(cmd *cobra.Command, prefix string) (*VMessServerOptions, error) {
	if prefix != "" {
		prefix = prefix + "."
	}

	enableTLS, err := cmd.Flags().GetBool(prefix + "enable-tls")
	if err != nil {
		return nil, err
	}

	listenPort, err := cmd.Flags().GetUint16(prefix + "listen-port")
	if err != nil {
		return nil, err
	}

	transport, err := cmd.Flags().GetString(prefix + "transport")
	if err != nil {
		return nil, err
	}

	return &VMessServerOptions{
		EnableTLS:  enableTLS,
		ListenPort: listenPort,
		Transport:  transport,
	}, nil
}

// ServerOptions represents the V2Ray server configuration options.
type ServerOptions struct {
	VMess *VMessServerOptions `json:"vmess"`
}

// WithVMess sets the VMess field and returns the modified ServerOptions instance.
func (so *ServerOptions) WithVMess(v *VMessServerOptions) *ServerOptions {
	so.VMess = v
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

// WriteConfigToFile writes the ServerOptions configuration to a file in JSON format.
func (so *ServerOptions) WriteConfigToFile(filepath string) error {
	return os.WriteFile(filepath, utils.MustMarshalJSON(so), 0644)
}

// Validate validates the ServerOptions fields.
func (so *ServerOptions) Validate() error {
	if so.VMess == nil {
		return errors.New("vmess cannot be empty")
	}
	if err := so.VMess.Validate(); err != nil {
		return fmt.Errorf("invalid vmess server options: %w", err)
	}

	return nil
}

// AddServerFlagsToCmd adds server-related flags to the given Cobra command.
func AddServerFlagsToCmd(cmd *cobra.Command, prefix string) {
	if prefix != "" {
		prefix = prefix + "."
	}

	AddVMessServerFlagsToCmd(cmd, prefix+"vmess")
}

// NewServerOptionsFromCmd creates and returns a ServerOptions instance from the given Cobra command's flags.
func NewServerOptionsFromCmd(cmd *cobra.Command, prefix string) (*ServerOptions, error) {
	if prefix != "" {
		prefix = prefix + "."
	}

	vmessOpts, err := NewVMessServerOptionsFromCmd(cmd, prefix+"vmess")
	if err != nil {
		return nil, err
	}

	return &ServerOptions{
		VMess: vmessOpts,
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
