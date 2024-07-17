package wireguard

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// execFile returns the name of the executable file.
func (c *Client) execFile(name string) string {
	return name
}

// interfaceName returns the path to the WireGuard executable.
func (c *Client) interfaceName() (string, error) {
	return ".\\" + filepath.Join("WireGuard", c.name+".exe"), nil
}

// Down uninstalls the WireGuard tunnel service.
func (c *Client) Down(ctx context.Context) error {
	iface, err := c.interfaceName()
	if err != nil {
		return err
	}

	// Executes the command to uninstall the WireGuard tunnel service.
	cmd := exec.CommandContext(
		ctx,
		c.execFile("wireguard"),
		strings.Fields(fmt.Sprintf("/uninstalltunnelservice %s", iface))...,
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

// Up installs the WireGuard tunnel service.
func (c *Client) Up(ctx context.Context) error {
	// Executes the command to install the WireGuard tunnel service.
	cmd := exec.CommandContext(
		ctx,
		c.execFile("wireguard"),
		strings.Fields(fmt.Sprintf("/uninstalltunnelservice %s", c.configFilePath()))...,
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
