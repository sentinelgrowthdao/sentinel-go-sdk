package wireguard

// AddPeerRequest represents a request to add a new peer in WireGuard.
type AddPeerRequest struct {
	PublicKey *Key `json:"public_key"`
}

// Key returns the public key as a string.
func (r *AddPeerRequest) Key() string {
	return r.PublicKey.String()
}

// Validate checks if the AddPeerRequest is valid.
func (r *AddPeerRequest) Validate() error {
	return nil
}

// NewAddPeerRequestFromKey creates a new AddPeerRequest from a base64-encoded public key string.
func NewAddPeerRequestFromKey(s string) (*AddPeerRequest, error) {
	// Use NewKeyFromString to create the Key
	key, err := NewKeyFromString(s)
	if err != nil {
		return nil, err
	}

	// Create and return the AddPeerRequest
	return &AddPeerRequest{
		PublicKey: key,
	}, nil
}

// HasPeerRequest represents a request to check if a peer exists in WireGuard.
type HasPeerRequest struct {
	PublicKey *Key `json:"public_key"`
}

// Key returns the public key as a string.
func (r *HasPeerRequest) Key() string {
	return r.PublicKey.String()
}

// Validate checks if the HasPeerRequest is valid.
func (r *HasPeerRequest) Validate() error {
	return nil
}

// NewHasPeerRequestFromKey creates a new HasPeerRequest from a base64-encoded public key string.
func NewHasPeerRequestFromKey(s string) (*HasPeerRequest, error) {
	// Use NewKeyFromString to create the Key
	key, err := NewKeyFromString(s)
	if err != nil {
		return nil, err
	}

	// Create and return the HasPeerRequest
	return &HasPeerRequest{
		PublicKey: key,
	}, nil
}

// RemovePeerRequest represents a request to remove an existing peer from WireGuard.
type RemovePeerRequest struct {
	PublicKey *Key `json:"public_key"`
}

// Key returns the public key as a string.
func (r *RemovePeerRequest) Key() string {
	return r.PublicKey.String()
}

// Validate checks if the RemovePeerRequest is valid.
func (r *RemovePeerRequest) Validate() error {
	return nil
}

// NewRemovePeerRequestFromKey creates a new RemovePeerRequest from a base64-encoded public key string.
func NewRemovePeerRequestFromKey(s string) (*RemovePeerRequest, error) {
	// Use NewKeyFromString to create the Key
	key, err := NewKeyFromString(s)
	if err != nil {
		return nil, err
	}

	// Create and return the RemovePeerRequest
	return &RemovePeerRequest{
		PublicKey: key,
	}, nil
}
