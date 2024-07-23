package v2ray

import (
	"path/filepath"
)

// execFile returns the name of the executable file for the V2Ray server.
func (s *Server) execFile(name string) string {
	return ".\\" + filepath.Join("V2Ray", name+".exe")
}
