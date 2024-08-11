package options

import (
	"github.com/spf13/cobra"

	"github.com/sentinel-official/sentinel-go-sdk/cmd/flags"
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
		Gas:                flags.DefaultTxGas,
		GasAdjustment:      flags.DefaultTxGasAdjustment,
		GasPrices:          flags.DefaultTxGasPrices,
		SimulateAndExecute: flags.DefaultTxSimulateAndExecute,
	}
}

// WithChainID sets the ChainID field and returns the modified Tx instance.
func (t *Tx) WithChainID(v string) *Tx {
	t.ChainID = v
	return t
}

// WithFeeGranterAddr sets the FeeGranterAddr field and returns the modified Tx instance.
func (t *Tx) WithFeeGranterAddr(v string) *Tx {
	t.FeeGranterAddr = v
	return t
}

// WithFees sets the Fees field and returns the modified Tx instance.
func (t *Tx) WithFees(v string) *Tx {
	t.Fees = v
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
func (t *Tx) WithGasPrices(v string) *Tx {
	t.GasPrices = v
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

// AddTxFlagsToCmd configures all transaction-related flags for the given command.
func AddTxFlagsToCmd(cmd *cobra.Command) {
	flags.SetFlagTxChainID(cmd)
	flags.SetFlagTxFeeGranterAddr(cmd)
	flags.SetFlagTxFees(cmd)
	flags.SetFlagTxFromName(cmd)
	flags.SetFlagTxGas(cmd)
	flags.SetFlagTxGasAdjustment(cmd)
	flags.SetFlagTxGasPrices(cmd)
	flags.SetFlagTxMemo(cmd)
	flags.SetFlagTxSimulateAndExecute(cmd)
	flags.SetFlagTxTimeoutHeight(cmd)
}

// NewTxFromCmd creates and returns Tx from the given cobra command's flags.
func NewTxFromCmd(cmd *cobra.Command) (*Tx, error) {
	// Retrieve the chain ID flag value from the command.
	chainID, err := flags.GetTxChainIDFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the fee granter address flag value from the command.
	feeGranterAddr, err := flags.GetTxFeeGranterAddrFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the fees flag value from the command.
	fees, err := flags.GetTxFeesFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the from name flag value from the command.
	fromName, err := flags.GetTxFromNameFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the gas flag value from the command.
	gas, err := flags.GetTxGasFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the gas adjustment flag value from the command.
	gasAdjustment, err := flags.GetTxGasAdjustmentFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the gas prices flag value from the command.
	gasPrices, err := flags.GetTxGasPricesFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the memo flag value from the command.
	memo, err := flags.GetTxMemoFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the simulate and execute flag value from the command.
	simulateAndExecute, err := flags.GetTxSimulateAndExecuteFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the timeout height flag value from the command.
	timeoutHeight, err := flags.GetTxTimeoutHeightFromCmd(cmd)
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
