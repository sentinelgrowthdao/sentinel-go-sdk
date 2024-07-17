package v2ray

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// execFile returns the name of the executable file for the V2Ray server.
func (s *Server) execFile(name string) string {
	return name
}

// Up starts the V2Ray server process.
func (s *Server) Up(ctx context.Context) error {
	// Constructs the command to start the V2Ray server.
	s.cmd = exec.CommandContext(
		ctx,
		s.execFile(v2ray),
		strings.Fields(fmt.Sprintf("run --config %s", s.configFilePath()))...,
	)
	s.cmd.Stdout = os.Stdout
	s.cmd.Stderr = os.Stderr

	// Starts the V2Ray server process.
	return s.cmd.Start()
}
