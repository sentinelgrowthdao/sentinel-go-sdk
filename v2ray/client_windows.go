package v2ray

import (
	"context"
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
func (c *Client) Up(ctx context.Context) error {
	// Constructs the command to start the V2Ray client.
	c.cmd = exec.CommandContext(
		ctx,
		c.execFile(v2ray),
		strings.Fields(fmt.Sprintf("run --config %s", c.configFilePath()))...,
	)
	c.cmd.Stdout = os.Stdout
	c.cmd.Stderr = os.Stderr

	// Starts the V2Ray client process.
	return c.cmd.Start()
}
