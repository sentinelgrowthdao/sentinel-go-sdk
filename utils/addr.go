package utils

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	base "github.com/sentinel-official/hub/v12/types"
)

// MustAccAddrFromBech32 converts a Bech32-encoded string to a sdk.AccAddress,
// panicking if there is an error during the conversion.
func MustAccAddrFromBech32(v string) sdk.AccAddress {
	// If the input string is empty, return nil
	if v == "" {
		return nil
	}

	// Attempt to convert the Bech32 string to a sdk.AccAddress
	addr, err := sdk.AccAddressFromBech32(v)

	// If there is an error during the conversion, panic
	if err != nil {
		panic(err)
	}

	// Return the converted address
	return addr
}

// MustNodeAddrFromBech32 converts a Bech32-encoded string to a base.NodeAddress,
// panicking if there is an error during the conversion.
func MustNodeAddrFromBech32(v string) base.NodeAddress {
	// If the input string is empty, return nil
	if v == "" {
		return nil
	}

	// Attempt to convert the Bech32 string to a base.NodeAddress
	addr, err := base.NodeAddressFromBech32(v)

	// If there is an error during the conversion, panic
	if err != nil {
		panic(err)
	}

	// Return the converted address
	return addr
}
