package client

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"

	"github.com/sentinel-official/sentinel-go-sdk/v1/types"
)

// Client contains necessary components for transaction handling, encoding and decoding.
type Client struct {
	codec.ProtoCodecMarshaler // Marshaler for protobuf types
	client.TxConfig           // Configuration for transactions
}

// New creates a new instance of Client.
// It takes a ProtoCodecMarshaler as input parameter and returns a pointer to Client.
func New(protoCodec codec.ProtoCodecMarshaler) *Client {
	return &Client{
		ProtoCodecMarshaler: protoCodec,
		TxConfig:            authtx.NewTxConfig(protoCodec, authtx.DefaultSignModes),
	}
}

// NewDefault creates a new instance of Client with a default ProtoCodecMarshaler.
// It returns a pointer to Client.
func NewDefault() *Client {
	return New(codec.NewProtoCodec(types.NewInterfaceRegistry()))
}
