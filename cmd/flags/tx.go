package flags

import (
	cosmossdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
)

// Default values for transaction options.
const (
	DefaultTxChainID            = "sentinelhub-2"
	DefaultTxGas                = 200_000
	DefaultTxGasAdjustment      = 1.0 + (1.0 / 6)
	DefaultTxSimulateAndExecute = true
)

// GetTxChainIDFromCmd retrieves the value of the tx.chain-id flag from the given command.
func GetTxChainIDFromCmd(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("tx.chain-id")
}

// GetTxFeeGranterAddrFromCmd retrieves the value of the tx.fee-granter-addr flag from the given command.
func GetTxFeeGranterAddrFromCmd(cmd *cobra.Command) (cosmossdk.AccAddress, error) {
	s, err := cmd.Flags().GetString("tx.fee-granter-addr")
	if err != nil {
		return nil, err
	}

	return cosmossdk.AccAddressFromBech32(s)
}

// GetTxFeesFromCmd retrieves the value of the tx.fees flag from the given command.
func GetTxFeesFromCmd(cmd *cobra.Command) (cosmossdk.Coins, error) {
	s, err := cmd.Flags().GetString("tx.fees")
	if err != nil {
		return nil, err
	}

	return cosmossdk.ParseCoinsNormalized(s)
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
func GetTxGasPricesFromCmd(cmd *cobra.Command) (cosmossdk.DecCoins, error) {
	s, err := cmd.Flags().GetString("tx.gas-prices")
	if err != nil {
		return nil, err
	}

	return cosmossdk.ParseDecCoins(s)
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
	cmd.Flags().String("tx.gas-prices", "", "Gas prices to be applied for transaction execution.")
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
