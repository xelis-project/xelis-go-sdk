package wallet

import "github.com/xelis-project/xelis-go-sdk/lib"

type GetAddressParams struct {
	IntegratedData string `json:"integrated_data"`
}

type SplitAddressParams struct {
	Address string `json:"address"`
}

type SplitAddressResult struct {
	Address        string `json:"address"`
	IntegratedData string `json:"integrated_data"`
}

type GetBalanceParams struct {
	Asset string `json:"asset"`
}

type GetTransactionParams = GetBalanceParams

type RescanParams struct {
	UntilTopoheight uint64 `json:"until_topoheight"`
}

type GetAssetPrecisionParams struct {
	Asset string `json:"asset"`
}

type Transfer struct {
	Amount uint64 `json:"amount"`
	Asset  string `json:"asset"`
	To     string `json:"to"`
}

type Burn struct {
	Asset  string `json:"asset"`
	Amount uint64 `json:"amount"`
}

type SmartContractCall struct {
	Contract string `json:"contract"`
	Assets   map[string]uint64
	Params   map[string]string
}

type TransactionType struct {
	Transfers      []Transfer        `json:"transfers"`
	Burn           Burn              `json:"burn"`
	CallContract   SmartContractCall `json:"call_contract"`
	DeployContract string            `json:"deploy_contract"`
}

type BuildTransactionParams struct {
	TxType    TransactionType `json:"tx_type"`
	Broadcast bool            `json:"broadcast"`
}

type BuildTransactionInner struct {
	Hash string      `json:"hash"`
	Data interface{} `json:"data"`
}

type BuildTransactionResult struct {
	TxAsHex string                `json:"tx_as_hex"`
	Inner   BuildTransactionInner `json:"inner"`
}

type ListTransactionsParams struct {
	MinTopoheight  uint64 `json:"min_topoheight"`
	MaxTopoheight  uint64 `json:"max_topoheight"`
	Address        string `json:"address"`
	AcceptIncoming bool   `json:"accept_incoming"`
	AcceptOutgoing bool   `json:"accept_outgoing"`
	AcceptCoinbase bool   `json:"accept_coinbase"`
	AcceptBurn     bool   `json:"accept_burn"`
}

type EstimateFeesParams struct {
	TxType TransactionType `json:"tx_type"`
}

const (
	GetVersion        lib.RPCMethod = "get_version"
	GetNetwork        lib.RPCMethod = "get_network"
	GetNonce          lib.RPCMethod = "get_nonce"
	GetTopoheight     lib.RPCMethod = "get_topoheight"
	GetAddress        lib.RPCMethod = "get_address"
	SplitAddress      lib.RPCMethod = "split_address"
	Rescan            lib.RPCMethod = "rescan"
	GetBalance        lib.RPCMethod = "get_balance"
	GetTrackedAssets  lib.RPCMethod = "get_tracked_assets"
	GetAssetPrecision lib.RPCMethod = "get_asset_precision"
	GetTransaction    lib.RPCMethod = "get_transaction"
	BuildTransaction  lib.RPCMethod = "build_transaction"
	ListTransactions  lib.RPCMethod = "list_transactions"
	IsOnline          lib.RPCMethod = "is_online"
	SignData          lib.RPCMethod = "sign_data"
	EstimateFees      lib.RPCMethod = "estimate_fees"
)
