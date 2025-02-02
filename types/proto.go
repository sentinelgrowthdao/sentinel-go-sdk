package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdkstd "github.com/cosmos/cosmos-sdk/std"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	authvestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	vpntypes "github.com/sentinel-official/hub/v12/x/vpn/types/v1"
)

// NewInterfaceRegistry initializes and returns a new InterfaceRegistry with registered interfaces.
func NewInterfaceRegistry() codectypes.InterfaceRegistry {
	// Create a new InterfaceRegistry instance.
	registry := codectypes.NewInterfaceRegistry()

	// Register Cosmos SDK module interfaces.
	sdkstd.RegisterInterfaces(registry)
	authtypes.RegisterInterfaces(registry)
	authvestingtypes.RegisterInterfaces(registry)
	authz.RegisterInterfaces(registry)
	banktypes.RegisterInterfaces(registry)
	feegrant.RegisterInterfaces(registry)

	// Register Sentinel Hub module interfaces.
	vpntypes.RegisterInterfaces(registry)

	// Return the populated InterfaceRegistry.
	return registry
}

// NewProtoCodec creates and returns a new ProtoCodecMarshaler with a populated InterfaceRegistry.
func NewProtoCodec() codec.ProtoCodecMarshaler {
	// Initialize the InterfaceRegistry.
	registry := NewInterfaceRegistry()

	// Create and return a new ProtoCodecMarshaler.
	return codec.NewProtoCodec(registry)
}
