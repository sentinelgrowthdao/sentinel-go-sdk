package options

import (
	"github.com/spf13/cobra"
)

// Default values for transaction options.
const (
	DefaultTxChainID            = "sentinelhub-2"
	DefaultGas                  = 200_000
	DefaultTxGasAdjustment      = 1.0 + (1.0 / 6)
	DefaultTxGasPrices          = "0.1udvpn"
	DefaultTxSimulateAndExecute = true
)

// TxOptions represents options for transactions.
type TxOptions struct {
	*KeyOptions                // Embedding KeyOptions for key-related options.
	*QueryOptions              // Embedding QueryOptions for query-related options.
	ChainID            string  `json:"chain_id,omitempty"`             // ChainID is the identifier of the blockchain network.
	FeeGranterAddr     string  `json:"fee_granter_addr,omitempty"`     // FeeGranterAddr is the address of the entity granting fees.
	Fees               string  `json:"fees,omitempty"`                 // Fees is the transaction fees.
	FromName           string  `json:"from_name,omitempty"`            // FromName is the name of the sender.
	GasAdjustment      float64 `json:"gas_adjustment,omitempty"`       // GasAdjustment is the adjustment factor for gas estimation.
	Gas                uint64  `json:"gas,omitempty"`                  // Gas is the gas limit for the transaction.
	GasPrices          string  `json:"gas_prices,omitempty"`           // GasPrices is the gas prices for transaction execution.
	Memo               string  `json:"memo,omitempty"`                 // Memo is a memo attached to the transaction.
	SimulateAndExecute bool    `json:"simulate_and_execute,omitempty"` // SimulateAndExecute indicates whether to simulate and execute the transaction.
	TimeoutHeight      uint64  `json:"timeout_height,omitempty"`       // TimeoutHeight is the block height at which the transaction times out.
}

// Tx creates a new TxOptions instance with default values.
func Tx() *TxOptions {
	return &TxOptions{
		KeyOptions:         Key(),   // Initialize embedded KeyOptions.
		QueryOptions:       Query(), // Initialize embedded QueryOptions.
		ChainID:            DefaultTxChainID,
		Gas:                DefaultGas,
		GasAdjustment:      DefaultTxGasAdjustment,
		GasPrices:          DefaultTxGasPrices,
		SimulateAndExecute: DefaultTxSimulateAndExecute,
	}
}

// WithKeyOptions sets the KeyOptions field and returns the modified TxOptions instance.
func (t *TxOptions) WithKeyOptions(v *KeyOptions) *TxOptions {
	t.KeyOptions = v
	return t
}

// WithQueryOptions sets the QueryOptions field and returns the modified TxOptions instance.
func (t *TxOptions) WithQueryOptions(v *QueryOptions) *TxOptions {
	t.QueryOptions = v
	return t
}

// WithChainID sets the ChainID field and returns the modified TxOptions instance.
func (t *TxOptions) WithChainID(v string) *TxOptions {
	t.ChainID = v
	return t
}

// WithFeeGranterAddr sets the FeeGranterAddr field and returns the modified TxOptions instance.
func (t *TxOptions) WithFeeGranterAddr(v string) *TxOptions {
	t.FeeGranterAddr = v
	return t
}

// WithFees sets the Fees field and returns the modified TxOptions instance.
func (t *TxOptions) WithFees(v string) *TxOptions {
	t.Fees = v
	return t
}

// WithFromName sets the FromName field and returns the modified TxOptions instance.
func (t *TxOptions) WithFromName(v string) *TxOptions {
	t.FromName = v
	return t
}

// WithGasAdjustment sets the GasAdjustment field and returns the modified TxOptions instance.
func (t *TxOptions) WithGasAdjustment(v float64) *TxOptions {
	t.GasAdjustment = v
	return t
}

// WithGas sets the Gas field and returns the modified TxOptions instance.
func (t *TxOptions) WithGas(v uint64) *TxOptions {
	t.Gas = v
	return t
}

// WithGasPrices sets the GasPrices field and returns the modified TxOptions instance.
func (t *TxOptions) WithGasPrices(v string) *TxOptions {
	t.GasPrices = v
	return t
}

// WithSimulateAndExecute sets the SimulateAndExecute field and returns the modified TxOptions instance.
func (t *TxOptions) WithSimulateAndExecute(v bool) *TxOptions {
	t.SimulateAndExecute = v
	return t
}

// WithTimeoutHeight sets the TimeoutHeight field and returns the modified TxOptions instance.
func (t *TxOptions) WithTimeoutHeight(v uint64) *TxOptions {
	t.TimeoutHeight = v
	return t
}

// AddTxFlagsToCmd adds transaction-related flags to the given cobra command.
func AddTxFlagsToCmd(cmd *cobra.Command) {
	// Add key and query related flags to the command.
	AddKeyFlagsToCmd(cmd)
	AddQueryFlagsToCmd(cmd)

	cmd.Flags().String("tx.chain-id", DefaultTxChainID, "Blockchain network identifier.")
	cmd.Flags().String("tx.fee-granter-addr", "", "Address of the entity granting fees for the transaction.")
	cmd.Flags().String("tx.fees", "", "Transaction fees to be paid.")
	cmd.Flags().String("tx.from-name", "", "Name of the sender's account in the keyring.")
	cmd.Flags().Float64("tx.gas-adjustment", DefaultTxGasAdjustment, "Factor to adjust gas estimation (used in simulation).")
	cmd.Flags().Uint64("tx.gas", DefaultGas, "Gas limit set for the transaction.")
	cmd.Flags().String("tx.gas-prices", DefaultTxGasPrices, "Gas prices to be applied for transaction execution.")
	cmd.Flags().String("tx.memo", "", "Memo text attached to the transaction.")
	cmd.Flags().Bool("tx.simulate-and-execute", DefaultTxSimulateAndExecute, "Flag to simulate the transaction before execution.")
	cmd.Flags().Uint64("tx.timeout-height", 0, "Block height after which the transaction will not be processed.")
}

// NewTxOptionsFromCmd creates and returns TxOptions from the given cobra command's flags.
func NewTxOptionsFromCmd(cmd *cobra.Command) (*TxOptions, error) {
	// Retrieve and create KeyOptions from the command's flags.
	keyOpts, err := NewKeyOptionsFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve and create QueryOptions from the command's flags.
	queryOpts, err := NewQueryOptionsFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the value of the "tx.chain-id" flag.
	chainID, err := cmd.Flags().GetString("tx.chain-id")
	if err != nil {
		return nil, err
	}

	// Retrieve the value of the "tx.fee-granter-addr" flag.
	feeGranterAddr, err := cmd.Flags().GetString("tx.fee-granter-addr")
	if err != nil {
		return nil, err
	}

	// Retrieve the value of the "tx.fees" flag.
	fees, err := cmd.Flags().GetString("tx.fees")
	if err != nil {
		return nil, err
	}

	// Retrieve the value of the "tx.from-name" flag.
	fromName, err := cmd.Flags().GetString("tx.from-name")
	if err != nil {
		return nil, err
	}

	// Retrieve the value of the "tx.gas-adjustment" flag.
	gasAdjustment, err := cmd.Flags().GetFloat64("tx.gas-adjustment")
	if err != nil {
		return nil, err
	}

	// Retrieve the value of the "tx.gas" flag.
	gas, err := cmd.Flags().GetUint64("tx.gas")
	if err != nil {
		return nil, err
	}

	// Retrieve the value of the "tx.gas-prices" flag.
	gasPrices, err := cmd.Flags().GetString("tx.gas-prices")
	if err != nil {
		return nil, err
	}

	// Retrieve the value of the "tx.memo" flag.
	memo, err := cmd.Flags().GetString("tx.memo")
	if err != nil {
		return nil, err
	}

	// Retrieve the value of the "tx.simulate-and-execute" flag.
	simulateAndExecute, err := cmd.Flags().GetBool("tx.simulate-and-execute")
	if err != nil {
		return nil, err
	}

	// Retrieve the value of the "tx.timeout-height" flag.
	timeoutHeight, err := cmd.Flags().GetUint64("tx.timeout-height")
	if err != nil {
		return nil, err
	}

	// Return a new TxOptions instance populated with the retrieved flag values, KeyOptions, and QueryOptions.
	return &TxOptions{
		KeyOptions:         keyOpts,
		QueryOptions:       queryOpts,
		ChainID:            chainID,
		FeeGranterAddr:     feeGranterAddr,
		Fees:               fees,
		FromName:           fromName,
		GasAdjustment:      gasAdjustment,
		Gas:                gas,
		GasPrices:          gasPrices,
		Memo:               memo,
		SimulateAndExecute: simulateAndExecute,
		TimeoutHeight:      timeoutHeight,
	}, nil
}
