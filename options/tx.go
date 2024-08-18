package options

import (
	"errors"

	cosmossdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"

	"github.com/sentinel-official/sentinel-go-sdk/flags"
)

// Tx represents options for transactions.
type Tx struct {
	ChainID            string  `json:"chain_id" toml:"chain_id"`                         // ChainID is the identifier of the blockchain network.
	FeeGranterAddr     string  `json:"fee_granter_addr" toml:"fee_granter_addr"`         // FeeGranterAddr is the address of the entity granting fees.
	Fees               string  `json:"fees" toml:"fees"`                                 // Fees is the transaction fees.
	FromName           string  `json:"from_name" toml:"from_name"`                       // FromName is the name of the sender.
	Gas                uint64  `json:"gas" toml:"gas"`                                   // Gas is the gas limit for the transaction.
	GasAdjustment      float64 `json:"gas_adjustment" toml:"gas_adjustment"`             // GasAdjustment is the adjustment factor for gas estimation.
	GasPrices          string  `json:"gas_prices" toml:"gas_prices"`                     // GasPrices is the gas prices for transaction execution.
	Memo               string  `json:"memo" toml:"memo"`                                 // Memo is a memo attached to the transaction.
	SimulateAndExecute bool    `json:"simulate_and_execute" toml:"simulate_and_execute"` // SimulateAndExecute indicates whether to simulate and execute the transaction.
	TimeoutHeight      uint64  `json:"timeout_height" toml:"timeout_height"`             // TimeoutHeight is the block height at which the transaction times out.
}

// NewTx creates a new Tx instance with default values.
func NewTx() *Tx {
	return &Tx{
		ChainID:            flags.DefaultTxChainID,
		FeeGranterAddr:     flags.DefaultTxFeeGranterAddr,
		Fees:               flags.DefaultTxFees,
		FromName:           flags.DefaultTxFromName,
		Gas:                flags.DefaultTxGas,
		GasAdjustment:      flags.DefaultTxGasAdjustment,
		GasPrices:          flags.DefaultTxGasPrices,
		Memo:               flags.DefaultTxMemo,
		SimulateAndExecute: flags.DefaultTxSimulateAndExecute,
		TimeoutHeight:      flags.DefaultTxTimeoutHeight,
	}
}

// WithChainID sets the ChainID field and returns the modified Tx instance.
func (t *Tx) WithChainID(v string) *Tx {
	t.ChainID = v
	return t
}

// WithFeeGranterAddr sets the FeeGranterAddr field and returns the modified Tx instance.
func (t *Tx) WithFeeGranterAddr(v cosmossdk.AccAddress) *Tx {
	t.FeeGranterAddr = v.String()
	return t
}

// WithFees sets the Fees field and returns the modified Tx instance.
func (t *Tx) WithFees(v cosmossdk.Coins) *Tx {
	t.Fees = v.String()
	return t
}

// WithFromName sets the FromName field and returns the modified Tx instance.
func (t *Tx) WithFromName(v string) *Tx {
	t.FromName = v
	return t
}

// WithGas sets the Gas field and returns the modified Tx instance.
func (t *Tx) WithGas(v uint64) *Tx {
	t.Gas = v
	return t
}

// WithGasAdjustment sets the GasAdjustment field and returns the modified Tx instance.
func (t *Tx) WithGasAdjustment(v float64) *Tx {
	t.GasAdjustment = v
	return t
}

// WithGasPrices sets the GasPrices field and returns the modified Tx instance.
func (t *Tx) WithGasPrices(v cosmossdk.DecCoins) *Tx {
	t.GasPrices = v.String()
	return t
}

// WithMemo sets the Memo field and returns the modified Tx instance.
func (t *Tx) WithMemo(v string) *Tx {
	t.Memo = v
	return t
}

// WithSimulateAndExecute sets the SimulateAndExecute field and returns the modified Tx instance.
func (t *Tx) WithSimulateAndExecute(v bool) *Tx {
	t.SimulateAndExecute = v
	return t
}

// WithTimeoutHeight sets the TimeoutHeight field and returns the modified Tx instance.
func (t *Tx) WithTimeoutHeight(v uint64) *Tx {
	t.TimeoutHeight = v
	return t
}

// GetChainID returns the ChainID field.
func (t *Tx) GetChainID() string {
	return t.ChainID
}

// GetFeeGranterAddr returns the FeeGranterAddr field.
func (t *Tx) GetFeeGranterAddr() cosmossdk.AccAddress {
	if t.FeeGranterAddr == "" {
		return nil
	}

	v, err := cosmossdk.AccAddressFromBech32(t.FeeGranterAddr)
	if err != nil {
		panic(err)
	}

	return v
}

// GetFees returns the Fees field.
func (t *Tx) GetFees() cosmossdk.Coins {
	v, err := cosmossdk.ParseCoinsNormalized(t.Fees)
	if err != nil {
		panic(err)
	}

	return v
}

// GetFromName returns the FromName field.
func (t *Tx) GetFromName() string {
	return t.FromName
}

// GetGas returns the Gas field.
func (t *Tx) GetGas() uint64 {
	return t.Gas
}

// GetGasAdjustment returns the GasAdjustment field.
func (t *Tx) GetGasAdjustment() float64 {
	return t.GasAdjustment
}

// GetGasPrices returns the GasPrices field.
func (t *Tx) GetGasPrices() cosmossdk.DecCoins {
	v, err := cosmossdk.ParseDecCoins(t.GasPrices)
	if err != nil {
		panic(err)
	}

	return v
}

// GetMemo returns the Memo field.
func (t *Tx) GetMemo() string {
	return t.Memo
}

// GetSimulateAndExecute returns the SimulateAndExecute field.
func (t *Tx) GetSimulateAndExecute() bool {
	return t.SimulateAndExecute
}

// GetTimeoutHeight returns the TimeoutHeight field.
func (t *Tx) GetTimeoutHeight() uint64 {
	return t.TimeoutHeight
}

// ValidateTxChainID validates the ChainID field.
func ValidateTxChainID(v string) error {
	if v == "" {
		return errors.New("chain_id must not be empty")
	}

	return nil
}

// ValidateTxFeeGranterAddr validates the FeeGranterAddr field.
func ValidateTxFeeGranterAddr(v string) error {
	if v == "" {
		return nil
	}
	if _, err := cosmossdk.AccAddressFromBech32(v); err != nil {
		return errors.New("fee_granter_addr must be a valid address")
	}

	return nil
}

// ValidateTxFees validates the Fees field.
func ValidateTxFees(v string) error {
	if _, err := cosmossdk.ParseCoinsNormalized(v); err != nil {
		return errors.New("fees must be a valid coins format")
	}

	return nil
}

// ValidateTxFromName validates the FromName field.
func ValidateTxFromName(v string) error {
	if v == "" {
		return errors.New("from_name must not be empty")
	}

	return nil
}

// ValidateTxGas validates the Gas field.
func ValidateTxGas(v uint64) error {
	if v == 0 {
		return errors.New("gas must be greater than zero")
	}

	return nil
}

// ValidateTxGasAdjustment validates the GasAdjustment field.
func ValidateTxGasAdjustment(v float64) error {
	if v <= 0 {
		return errors.New("gas_adjustment must be greater than zero")
	}

	return nil
}

// ValidateTxGasPrices validates the GasPrices field.
func ValidateTxGasPrices(v string) error {
	if _, err := cosmossdk.ParseDecCoins(v); err != nil {
		return errors.New("gas_prices must be a valid decimal coins format")
	}

	return nil
}

// Validate validates all the fields of the Tx struct.
func (t *Tx) Validate() error {
	if err := ValidateTxChainID(t.ChainID); err != nil {
		return err
	}
	if err := ValidateTxFeeGranterAddr(t.FeeGranterAddr); err != nil {
		return err
	}
	if err := ValidateTxFees(t.Fees); err != nil {
		return err
	}
	if err := ValidateTxFromName(t.FromName); err != nil {
		return err
	}
	if err := ValidateTxGas(t.Gas); err != nil {
		return err
	}
	if err := ValidateTxGasAdjustment(t.GasAdjustment); err != nil {
		return err
	}
	if err := ValidateTxGasPrices(t.GasPrices); err != nil {
		return err
	}

	return nil
}

// NewTxFromCmd creates and returns Tx from the given cobra command's flags.
func NewTxFromCmd(cmd *cobra.Command) (*Tx, error) {
	// Retrieve the chain ID flag value from the command.
	chainID, err := flags.GetTxChainID(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the fee granter address flag value from the command.
	feeGranterAddr, err := flags.GetTxFeeGranterAddr(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the fees flag value from the command.
	fees, err := flags.GetTxFees(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the from name flag value from the command.
	fromName, err := flags.GetTxFromName(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the gas flag value from the command.
	gas, err := flags.GetTxGas(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the gas adjustment flag value from the command.
	gasAdjustment, err := flags.GetTxGasAdjustment(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the gas prices flag value from the command.
	gasPrices, err := flags.GetTxGasPrices(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the memo flag value from the command.
	memo, err := flags.GetTxMemo(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the simulate and execute flag value from the command.
	simulateAndExecute, err := flags.GetTxSimulateAndExecute(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the timeout height flag value from the command.
	timeoutHeight, err := flags.GetTxTimeoutHeight(cmd)
	if err != nil {
		return nil, err
	}

	// Return a new Tx instance populated with the retrieved flag values.
	return &Tx{
		ChainID:            chainID,
		FeeGranterAddr:     feeGranterAddr,
		Fees:               fees,
		FromName:           fromName,
		Gas:                gas,
		GasAdjustment:      gasAdjustment,
		GasPrices:          gasPrices,
		Memo:               memo,
		SimulateAndExecute: simulateAndExecute,
		TimeoutHeight:      timeoutHeight,
	}, nil
}
