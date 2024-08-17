package flags

import (
	"github.com/spf13/cobra"
)

// Default values for transaction options.
const (
	DefaultTxChainID            = "sentinelhub-2"
	DefaultTxFeeGranterAddr     = ""
	DefaultTxFees               = ""
	DefaultTxFromName           = ""
	DefaultTxGas                = 200_000
	DefaultTxGasAdjustment      = 1.0 + (1.0 / 6)
	DefaultTxGasPrices          = ""
	DefaultTxMemo               = ""
	DefaultTxSimulateAndExecute = true
	DefaultTxTimeoutHeight      = 0
)

// GetTxChainID retrieves the value of the tx.chain-id flag from the given command.
func GetTxChainID(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("tx.chain-id")
}

// GetTxFeeGranterAddr retrieves the value of the tx.fee-granter-addr flag from the given command.
func GetTxFeeGranterAddr(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("tx.fee-granter-addr")
}

// GetTxFees retrieves the value of the tx.fees flag from the given command.
func GetTxFees(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("tx.fees")
}

// GetTxFromName retrieves the value of the tx.from-name flag from the given command.
func GetTxFromName(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("tx.from-name")
}

// GetTxGas retrieves the value of the tx.gas flag from the given command.
func GetTxGas(cmd *cobra.Command) (uint64, error) {
	return cmd.Flags().GetUint64("tx.gas")
}

// GetTxGasAdjustment retrieves the value of the tx.gas-adjustment flag from the given command.
func GetTxGasAdjustment(cmd *cobra.Command) (float64, error) {
	return cmd.Flags().GetFloat64("tx.gas-adjustment")
}

// GetTxGasPrices retrieves the value of the tx.gas-prices flag from the given command.
func GetTxGasPrices(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("tx.gas-prices")
}

// GetTxMemo retrieves the value of the tx.memo flag from the given command.
func GetTxMemo(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("tx.memo")
}

// GetTxSimulateAndExecute retrieves the value of the tx.simulate-and-execute flag from the given command.
func GetTxSimulateAndExecute(cmd *cobra.Command) (bool, error) {
	return cmd.Flags().GetBool("tx.simulate-and-execute")
}

// GetTxTimeoutHeight retrieves the value of the tx.timeout-height flag from the given command.
func GetTxTimeoutHeight(cmd *cobra.Command) (uint64, error) {
	return cmd.Flags().GetUint64("tx.timeout-height")
}

// SetFlagTxChainID adds the tx.chain-id flag to the given command.
func SetFlagTxChainID(cmd *cobra.Command) {
	cmd.Flags().String("tx.chain-id", DefaultTxChainID, "Blockchain network identifier.")
}

// SetFlagTxFeeGranterAddr adds the tx.fee-granter-addr flag to the given command.
func SetFlagTxFeeGranterAddr(cmd *cobra.Command) {
	cmd.Flags().String("tx.fee-granter-addr", DefaultTxFeeGranterAddr, "Address of the entity granting fees for the transaction.")
}

// SetFlagTxFees adds the tx.fees flag to the given command.
func SetFlagTxFees(cmd *cobra.Command) {
	cmd.Flags().String("tx.fees", DefaultTxFees, "Transaction fees to be paid.")
}

// SetFlagTxFromName adds the tx.from-name flag to the given command.
func SetFlagTxFromName(cmd *cobra.Command) {
	cmd.Flags().String("tx.from-name", DefaultTxFromName, "Name of the sender's account in the keyring.")
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
	cmd.Flags().String("tx.memo", DefaultTxMemo, "Memo text attached to the transaction.")
}

// SetFlagTxSimulateAndExecute adds the tx.simulate-and-execute flag to the given command.
func SetFlagTxSimulateAndExecute(cmd *cobra.Command) {
	cmd.Flags().Bool("tx.simulate-and-execute", DefaultTxSimulateAndExecute, "Flag to simulate the transaction before execution.")
}

// SetFlagTxTimeoutHeight adds the tx.timeout-height flag to the given command.
func SetFlagTxTimeoutHeight(cmd *cobra.Command) {
	cmd.Flags().Uint64("tx.timeout-height", DefaultTxTimeoutHeight, "Block height after which the transaction will not be processed.")
}

// AddTxFlags configures all transaction-related flags for the given command.
func AddTxFlags(cmd *cobra.Command) {
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
