package v2ray

import (
	"fmt"

	"github.com/v2fly/v2ray-core/v5/common/uuid"
	"google.golang.org/protobuf/types/known/anypb"
)

// Tag represents a composite data structure combining Protocol, Network, and Security.
type Tag struct {
	p Protocol
	n Network
	s Security
}

// String returns a string representation of the Tag in the format "protocol_network_security".
func (t *Tag) String() string {
	return fmt.Sprintf("%s_%s_%s", t.p, t.n, t.s)
}

// Account generates an account message based on the Protocol stored in the Tag.
func (t *Tag) Account(uid uuid.UUID) *anypb.Any {
	return t.p.Account(uid)
}
