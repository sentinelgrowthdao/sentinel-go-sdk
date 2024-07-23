//go:build darwin || linux

package v2ray

// execFile returns the name of the executable file for the V2Ray server.
func (s *Server) execFile(name string) string {
	return name
}
