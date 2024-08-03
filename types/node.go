package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	base "github.com/sentinel-official/hub/v12/types"

	"github.com/sentinel-official/sentinel-go-sdk/libs/geoip"
)

type (
	Bandwidth struct {
		Download int64 `json:"download"`
		Upload   int64 `json:"upload"`
	}
	Handshake struct {
		Enable bool   `json:"enable"`
		Peers  uint64 `json:"peers"`
	}
	QOS struct {
		MaxPeers int `json:"max_peers"`
	}
)

type NodeInfo struct {
	Address                base.NodeAddress `json:"address"`
	Bandwidth              *Bandwidth       `json:"bandwidth"`
	Handshake              *Handshake       `json:"handshake"`
	IntervalSetSessions    time.Duration    `json:"interval_set_sessions"`
	IntervalUpdateSessions time.Duration    `json:"interval_update_sessions"`
	IntervalUpdateStatus   time.Duration    `json:"interval_update_status"`
	Location               *geoip.Location  `json:"location"`
	Moniker                string           `json:"moniker"`
	Operator               sdk.AccAddress   `json:"operator"`
	Peers                  int              `json:"peers"`
	GigabytePrices         sdk.Coins        `json:"gigabyte_prices"`
	HourlyPrices           sdk.Coins        `json:"hourly_prices"`
	QOS                    *QOS             `json:"qos"`
	Type                   ServiceType      `json:"type"`
	Version                string           `json:"version"`
}
