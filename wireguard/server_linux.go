package wireguard

// interfaceName returns the name of the WireGuard interface.
func (s *Server) interfaceName() (string, error) {
	return s.name, nil
}
