package types

import (
	"sync"
)

// Peer represents an entity with an Email field.
type Peer struct {
	Email string
}

// Key returns the unique identifier (email) associated with the Peer.
func (p *Peer) Key() string {
	return p.Email
}

// PeerManager is a thread-safe map-like structure that stores Peer objects.
type PeerManager struct {
	*sync.RWMutex
	m map[string]*Peer
}

// NewPeerManager creates and returns a new instance of PeerManager.
func NewPeerManager() *PeerManager {
	return &PeerManager{
		RWMutex: &sync.RWMutex{},
		m:       make(map[string]*Peer),
	}
}

// Get retrieves a Peer from the PeerManager based on the provided key.
// It returns the corresponding Peer if found, or nil otherwise.
func (p *PeerManager) Get(v string) *Peer {
	p.RLock()
	defer p.RUnlock()

	value, ok := p.m[v]
	if !ok {
		return nil
	}

	return value
}

// Put adds a Peer to the PeerManager.
// If a Peer with the same key already exists, it does nothing.
func (p *PeerManager) Put(v *Peer) {
	p.Lock()
	defer p.Unlock()

	_, ok := p.m[v.Key()]
	if ok {
		return
	}

	p.m[v.Key()] = v
}

// Delete removes a Peer from the PeerManager based on the provided key.
func (p *PeerManager) Delete(v string) {
	p.Lock()
	defer p.Unlock()

	delete(p.m, v)
}

// Len returns the number of elements in the PeerManager.
func (p *PeerManager) Len() int {
	p.RLock()
	defer p.RUnlock()

	return len(p.m)
}

// Iterate iterates over each element in the PeerManager and applies the provided function.
// If the function returns true, the iteration stops.
// If the function returns an error, the iteration stops and the error is returned.
func (p *PeerManager) Iterate(fn func(key string, value *Peer) (bool, error)) error {
	p.RLock()
	defer p.RUnlock()

	for key, value := range p.m {
		stop, err := fn(key, value)
		if err != nil {
			return err
		}

		if stop {
			return nil
		}
	}

	return nil
}
