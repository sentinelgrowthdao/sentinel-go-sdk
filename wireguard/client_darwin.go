package wireguard

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// execFile returns the name of the executable file.
func (c *Client) execFile(name string) string {
	return name
}

// interfaceName retrieves the name of the WireGuard interface.
func (c *Client) interfaceName() (string, error) {
	// Opens the file containing the interface name.
	nameFile, err := os.Open(fmt.Sprintf("/var/run/wireguard/%s.name", c.cfg.Name))
	if err != nil {
		return "", err
	}
	defer nameFile.Close()

	// Reads the interface name from the file.
	reader := bufio.NewReader(nameFile)
	line, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.Trim(line, "\n"), nil
}

// Down shuts down the WireGuard interface.
func (c *Client) Down(ctx context.Context) error {
	// Executes the 'wg-quick down' command to bring down the interface.
	cmd := exec.CommandContext(
		ctx,
		c.execFile("wg-quick"),
		strings.Fields(fmt.Sprintf("down %s", c.configFilePath()))...,
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

// Up starts the WireGuard interface.
func (c *Client) Up(ctx context.Context) error {
	// Executes the 'wg-quick up' command to bring up the interface.
	cmd := exec.CommandContext(
		ctx,
		c.execFile("wg-quick"),
		strings.Fields(fmt.Sprintf("up %s", c.configFilePath()))...,
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
