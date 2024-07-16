package client

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"

	"github.com/sentinel-official/sentinel-go-sdk/v1/types"
)

// Context contains necessary components for transaction handling, encoding and decoding.
type Context struct {
	codec.ProtoCodecMarshaler // Marshaler for protobuf types
	client.TxConfig           // Configuration for transactions
}

// NewContext creates a new instance of Context.
// It takes a ProtoCodecMarshaler as input parameter and returns a pointer to Context.
func NewContext(protoCodec codec.ProtoCodecMarshaler) *Context {
	return &Context{
		ProtoCodecMarshaler: protoCodec,
		TxConfig:            authtx.NewTxConfig(protoCodec, authtx.DefaultSignModes),
	}
}

// NewDefaultContext creates a new instance of Context with a default ProtoCodecMarshaler.
// It returns a pointer to Context.
func NewDefaultContext() *Context {
	return NewContext(codec.NewProtoCodec(types.NewInterfaceRegistry()))
}
