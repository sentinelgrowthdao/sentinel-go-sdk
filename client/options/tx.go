package options

// Default values for transaction options.
const (
	DefaultTxBroadcastMode      = "sync"
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
	BroadcastMode      string  `json:"broadcast_mode,omitempty"`       // BroadcastMode is the mode of broadcasting transactions.
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
		BroadcastMode:      DefaultTxBroadcastMode,
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

// WithBroadcastMode sets the BroadcastMode field and returns the modified TxOptions instance.
func (t *TxOptions) WithBroadcastMode(v string) *TxOptions {
	t.BroadcastMode = v
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
