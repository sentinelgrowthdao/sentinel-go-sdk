package v2ray

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	v2ray = "v2ray.exe"
)

// execFile returns the name of the executable file.
func (c *Client) execFile(name string) string {
	return ".\\" + filepath.Join("V2Ray", name)
}

// Up starts the V2Ray client process.
func (c *Client) Up() error {
	// Constructs the command to start the V2Ray client.
	c.cmd = exec.Command(
		c.execFile(v2ray),
		strings.Fields(fmt.Sprintf("run --config %s", c.configFilePath()))...,
	)
	c.cmd.Stdout = os.Stdout
	c.cmd.Stderr = os.Stderr

	// Starts the V2Ray client process.
	if err := c.cmd.Start(); err != nil {
		return err
	}

	return nil
}
