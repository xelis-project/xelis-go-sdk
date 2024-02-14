package wallet

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
	GetVersion        string = "get_version"
	GetNetwork        string = "get_network"
	GetNonce          string = "get_nonce"
	GetTopoheight     string = "get_topoheight"
	GetAddress        string = "get_address"
	SplitAddress      string = "split_address"
	Rescan            string = "rescan"
	GetBalance        string = "get_balance"
	GetTrackedAssets  string = "get_tracked_assets"
	GetAssetPrecision string = "get_asset_precision"
	GetTransaction    string = "get_transaction"
	BuildTransaction  string = "build_transaction"
	ListTransactions  string = "list_transactions"
	IsOnline          string = "is_online"
	SignData          string = "sign_data"
	EstimateFees      string = "estimate_fees"
)
