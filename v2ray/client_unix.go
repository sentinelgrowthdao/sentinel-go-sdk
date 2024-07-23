//go:build darwin || linux

package v2ray

// execFile returns the name of the executable file.
func (c *Client) execFile(name string) string {
	return name
}
