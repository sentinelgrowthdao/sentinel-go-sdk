package wireguard

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// execFile returns the name of the executable file.
func (c *Client) execFile(name string) string {
	return name
}

// interfaceName returns the name of the WireGuard interface.
func (c *Client) interfaceName() (string, error) {
	return c.name, nil
}

// Down shuts down the WireGuard interface.
func (c *Client) Down() error {
	// Executes the 'wg-quick down' command to bring down the interface.
	cmd := exec.Command(
		c.execFile("wg-quick"),
		strings.Fields(fmt.Sprintf("down %s", c.configFilePath()))...,
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

// Up starts the WireGuard interface.
func (c *Client) Up() error {
	// Executes the 'wg-quick up' command to bring up the interface.
	cmd := exec.Command(
		c.execFile("wg-quick"),
		strings.Fields(fmt.Sprintf("up %s", c.configFilePath()))...,
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
