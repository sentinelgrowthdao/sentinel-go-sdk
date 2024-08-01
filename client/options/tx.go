package options

import (
	"github.com/spf13/cobra"
)

// Default values for transaction options.
const (
	DefaultTxChainID            = "sentinelhub-2"
	DefaultTxGas                = 200_000
	DefaultTxGasAdjustment      = 1.0 + (1.0 / 6)
	DefaultTxGasPrices          = "0.1udvpn"
	DefaultTxSimulateAndExecute = true
)

// TxOptions represents options for transactions.
type TxOptions struct {
	ChainID            string  `json:"chain_id,omitempty"`             // ChainID is the identifier of the blockchain network.
	FeeGranterAddr     string  `json:"fee_granter_addr,omitempty"`     // FeeGranterAddr is the address of the entity granting fees.
	Fees               string  `json:"fees,omitempty"`                 // Fees is the transaction fees.
	FromName           string  `json:"from_name,omitempty"`            // FromName is the name of the sender.
	Gas                uint64  `json:"gas,omitempty"`                  // Gas is the gas limit for the transaction.
	GasAdjustment      float64 `json:"gas_adjustment,omitempty"`       // GasAdjustment is the adjustment factor for gas estimation.
	GasPrices          string  `json:"gas_prices,omitempty"`           // GasPrices is the gas prices for transaction execution.
	Memo               string  `json:"memo,omitempty"`                 // Memo is a memo attached to the transaction.
	SimulateAndExecute bool    `json:"simulate_and_execute,omitempty"` // SimulateAndExecute indicates whether to simulate and execute the transaction.
	TimeoutHeight      uint64  `json:"timeout_height,omitempty"`       // TimeoutHeight is the block height at which the transaction times out.
}

// NewDefaultTxOptions creates a new TxOptions instance with default values.
func NewDefaultTxOptions() *TxOptions {
	return &TxOptions{
		ChainID:            DefaultTxChainID,
		Gas:                DefaultTxGas,
		GasAdjustment:      DefaultTxGasAdjustment,
		GasPrices:          DefaultTxGasPrices,
		SimulateAndExecute: DefaultTxSimulateAndExecute,
	}
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

// WithGas sets the Gas field and returns the modified TxOptions instance.
func (t *TxOptions) WithGas(v uint64) *TxOptions {
	t.Gas = v
	return t
}

// WithGasAdjustment sets the GasAdjustment field and returns the modified TxOptions instance.
func (t *TxOptions) WithGasAdjustment(v float64) *TxOptions {
	t.GasAdjustment = v
	return t
}

// WithGasPrices sets the GasPrices field and returns the modified TxOptions instance.
func (t *TxOptions) WithGasPrices(v string) *TxOptions {
	t.GasPrices = v
	return t
}

// WithMemo sets the Memo field and returns the modified TxOptions instance.
func (t *TxOptions) WithMemo(v string) *TxOptions {
	t.Memo = v
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

// SetFlagTxChainID adds the tx.chain-id flag to the given command.
func SetFlagTxChainID(cmd *cobra.Command) {
	cmd.Flags().String("tx.chain-id", DefaultTxChainID, "Blockchain network identifier.")
}

// SetFlagTxFeeGranterAddr adds the tx.fee-granter-addr flag to the given command.
func SetFlagTxFeeGranterAddr(cmd *cobra.Command) {
	cmd.Flags().String("tx.fee-granter-addr", "", "Address of the entity granting fees for the transaction.")
}

// SetFlagTxFees adds the tx.fees flag to the given command.
func SetFlagTxFees(cmd *cobra.Command) {
	cmd.Flags().String("tx.fees", "", "Transaction fees to be paid.")
}

// SetFlagTxFromName adds the tx.from-name flag to the given command.
func SetFlagTxFromName(cmd *cobra.Command) {
	cmd.Flags().String("tx.from-name", "", "Name of the sender's account in the keyring.")
}

// SetFlagTxGas adds the tx.gas flag to the given command.
func SetFlagTxGas(cmd *cobra.Command) {
	cmd.Flags().Uint64("tx.gas", DefaultTxGas, "Gas limit set for the transaction.")
}

// SetFlagTxGasAdjustment adds the tx.gas-adjustment flag to the given command.
func SetFlagTxGasAdjustment(cmd *cobra.Command) {
	cmd.Flags().Float64("tx.gas-adjustment", DefaultTxGasAdjustment, "Factor to adjust gas estimation (used in simulation).")
}

// SetFlagTxGasPrices adds the tx.gas-prices flag to the given command.
func SetFlagTxGasPrices(cmd *cobra.Command) {
	cmd.Flags().String("tx.gas-prices", DefaultTxGasPrices, "Gas prices to be applied for transaction execution.")
}

// SetFlagTxMemo adds the tx.memo flag to the given command.
func SetFlagTxMemo(cmd *cobra.Command) {
	cmd.Flags().String("tx.memo", "", "Memo text attached to the transaction.")
}

// SetFlagTxSimulateAndExecute adds the tx.simulate-and-execute flag to the given command.
func SetFlagTxSimulateAndExecute(cmd *cobra.Command) {
	cmd.Flags().Bool("tx.simulate-and-execute", DefaultTxSimulateAndExecute, "Flag to simulate the transaction before execution.")
}

// SetFlagTxTimeoutHeight adds the tx.timeout-height flag to the given command.
func SetFlagTxTimeoutHeight(cmd *cobra.Command) {
	cmd.Flags().Uint64("tx.timeout-height", 0, "Block height after which the transaction will not be processed.")
}

// GetTxChainIDFromCmd retrieves the value of the tx.chain-id flag from the given command.
func GetTxChainIDFromCmd(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("tx.chain-id")
}

// GetTxFeeGranterAddrFromCmd retrieves the value of the tx.fee-granter-addr flag from the given command.
func GetTxFeeGranterAddrFromCmd(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("tx.fee-granter-addr")
}

// GetTxFeesFromCmd retrieves the value of the tx.fees flag from the given command.
func GetTxFeesFromCmd(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("tx.fees")
}

// GetTxFromNameFromCmd retrieves the value of the tx.from-name flag from the given command.
func GetTxFromNameFromCmd(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("tx.from-name")
}

// GetTxGasFromCmd retrieves the value of the tx.gas flag from the given command.
func GetTxGasFromCmd(cmd *cobra.Command) (uint64, error) {
	return cmd.Flags().GetUint64("tx.gas")
}

// GetTxGasAdjustmentFromCmd retrieves the value of the tx.gas-adjustment flag from the given command.
func GetTxGasAdjustmentFromCmd(cmd *cobra.Command) (float64, error) {
	return cmd.Flags().GetFloat64("tx.gas-adjustment")
}

// GetTxGasPricesFromCmd retrieves the value of the tx.gas-prices flag from the given command.
func GetTxGasPricesFromCmd(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("tx.gas-prices")
}

// GetTxMemoFromCmd retrieves the value of the tx.memo flag from the given command.
func GetTxMemoFromCmd(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("tx.memo")
}

// GetTxSimulateAndExecuteFromCmd retrieves the value of the tx.simulate-and-execute flag from the given command.
func GetTxSimulateAndExecuteFromCmd(cmd *cobra.Command) (bool, error) {
	return cmd.Flags().GetBool("tx.simulate-and-execute")
}

// GetTxTimeoutHeightFromCmd retrieves the value of the tx.timeout-height flag from the given command.
func GetTxTimeoutHeightFromCmd(cmd *cobra.Command) (uint64, error) {
	return cmd.Flags().GetUint64("tx.timeout-height")
}

// AddTxFlagsToCmd configures all transaction-related flags for the given command.
func AddTxFlagsToCmd(cmd *cobra.Command) {
	SetFlagTxChainID(cmd)
	SetFlagTxFeeGranterAddr(cmd)
	SetFlagTxFees(cmd)
	SetFlagTxFromName(cmd)
	SetFlagTxGas(cmd)
	SetFlagTxGasAdjustment(cmd)
	SetFlagTxGasPrices(cmd)
	SetFlagTxMemo(cmd)
	SetFlagTxSimulateAndExecute(cmd)
	SetFlagTxTimeoutHeight(cmd)
}

// NewTxOptionsFromCmd creates and returns TxOptions from the given cobra command's flags.
func NewTxOptionsFromCmd(cmd *cobra.Command) (*TxOptions, error) {
	// Retrieve the value of the "tx.chain-id" flag.
	chainID, err := GetTxChainIDFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the value of the "tx.fee-granter-addr" flag.
	feeGranterAddr, err := GetTxFeeGranterAddrFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the value of the "tx.fees" flag.
	fees, err := GetTxFeesFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the value of the "tx.from-name" flag.
	fromName, err := GetTxFromNameFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the value of the "tx.gas" flag.
	gas, err := GetTxGasFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the value of the "tx.gas-adjustment" flag.
	gasAdjustment, err := GetTxGasAdjustmentFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the value of the "tx.gas-prices" flag.
	gasPrices, err := GetTxGasPricesFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the value of the "tx.memo" flag.
	memo, err := GetTxMemoFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the value of the "tx.simulate-and-execute" flag.
	simulateAndExecute, err := GetTxSimulateAndExecuteFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the value of the "tx.timeout-height" flag.
	timeoutHeight, err := GetTxTimeoutHeightFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Return a new TxOptions instance populated with the retrieved flag values.
	return &TxOptions{
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
