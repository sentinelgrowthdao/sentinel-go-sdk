package wireguard

// interfaceName returns the name of the WireGuard interface.
func (c *Client) interfaceName() (string, error) {
	return c.name, nil
}
