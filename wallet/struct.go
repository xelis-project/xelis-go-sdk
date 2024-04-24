package wallet

import (
	"github.com/xelis-project/xelis-go-sdk/daemon"
)

type GetAddressParams struct {
	IntegratedData *interface{} `json:"integrated_data,omitempty"`
}

type SplitAddressParams struct {
	Address string `json:"address"`
}

type SplitAddressResult struct {
	Address        string      `json:"address"`
	IntegratedData interface{} `json:"integrated_data"`
}

type GetBalanceParams struct {
	Asset string `json:"asset"`
}

type GetTransactionParams struct {
	Hash string `json:"hash"`
}

type RescanParams struct {
	UntilTopoheight uint64 `json:"until_topoheight"`
}

type GetAssetPrecisionParams struct {
	Asset string `json:"asset"`
}

type TransferIn struct {
	Amount    uint64       `json:"amount"`
	Asset     string       `json:"asset"`
	ExtraData *interface{} `json:"extra_data"`
}

/*
// we use *interface{} instead of DataElement so user can serialize/deserialize how he wants
type DataElement struct {
	Value  interface{}     `json:"value,omitempty"`
	Array  []DataElement   `json:"array,omitempty"`
	Fields json.RawMessage `json:"fields,omitempty"` // can't do map[interface{}]DataElement json unsupported parsing
}
*/

type TransferOut struct {
	Amount      uint64       `json:"amount"`
	Asset       string       `json:"asset"`
	Destination string       `json:"destination"`
	ExtraData   *interface{} `json:"extra_data,omitempty"`
}

type FeeBuilder struct {
	Multiplier *float64 `json:"multiplier,omitempty"`
	Value      *uint64  `json:"value,omitempty"`
}

type BuildTransactionParams struct {
	Transfers []TransferOut `json:"transfers"`
	Burn      *daemon.Burn  `json:"burn,omitempty"`
	Broadcast bool          `json:"broadcast"`
	TxAsHex   bool          `json:"tx_as_hex"`
	Fee       *FeeBuilder   `json:"fee,omitempty"`
}

// !!! not the same as daemon.Transfer
// the destination is []byte and the other it's string
type Transfer struct {
	Asset           string       `json:"asset"`
	ExtraData       *[]byte      `json:"extra_data"`
	Destination     []byte       `json:"destination"`
	Commitment      []byte       `json:"commitment"`
	SenderHandle    []byte       `json:"sender_handle"`
	ReceiverHandle  []byte       `json:"receiver_handle"`
	CTValidityProof daemon.Proof `json:"ct_validity_proof"`
}

type TransactionData struct {
	Transfers []Transfer   `json:"transfers"`
	Burn      *daemon.Burn `json:"burn"`
}

type BuildTransactionResult struct {
	Data              TransactionData           `json:"data"`
	Fee               uint64                    `json:"fee"`
	Hash              string                    `json:"hash"`
	Nonce             uint64                    `json:"nonce"`
	RangeProof        []byte                    `json:"range_proof"`
	Reference         daemon.Reference          `json:"reference"`
	Signature         string                    `json:"signature"`
	Source            []byte                    `json:"source"`
	SourceCommitments []daemon.SourceCommitment `json:"source_commitments"`
	TxAsHex           string                    `json:"tx_as_hex"`
	Version           uint64                    `json:"version"`
}

type Outgoing struct {
	Fee       uint64        `json:"fee"`
	Nonce     uint64        `json:"nonce"`
	Transfers []TransferOut `json:"transfers"`
}

type Incoming struct {
	From      string       `json:"from"`
	Transfers []TransferIn `json:"transfers"`
}

type Coinbase struct {
	Reward uint64 `json:"reward"`
}

type TransactionEntry struct {
	Hash       string       `json:"hash"`
	Topoheight uint64       `json:"topoheight"`
	Outgoing   *Outgoing    `json:"outgoing"`
	Burn       *daemon.Burn `json:"burn"`
	Incoming   *Incoming    `json:"incoming"`
	Coinbase   *Coinbase    `json:"coinbase"`
}

type ListTransactionsParams struct {
	MinTopoheight  *uint64 `json:"min_topoheight"`
	MaxTopoheight  *uint64 `json:"max_topoheight"`
	Address        *string `json:"address"`
	AcceptIncoming bool    `json:"accept_incoming"`
	AcceptOutgoing bool    `json:"accept_outgoing"`
	AcceptCoinbase bool    `json:"accept_coinbase"`
	AcceptBurn     bool    `json:"accept_burn"`
}

type EstimateFeesParams struct {
	Transfers *[]TransferOut `json:"transfers"`
	Burn      *daemon.Burn   `json:"burn"`
}

type BalanceChangedResult struct {
	Asset   string `json:"asset"`
	Balance uint64 `json:"balance"`
}

// Methods
const (
	GetVersion        string = "get_version"
	GetNetwork        string = "get_network"
	GetNonce          string = "get_nonce"
	GetTopoheight     string = "get_topoheight"
	GetAddress        string = "get_address"
	SplitAddress      string = "split_address"
	Rescan            string = "rescan"
	GetBalance        string = "get_balance"
	HasBalance        string = "has_balance"
	GetTrackedAssets  string = "get_tracked_assets"
	GetAssetPrecision string = "get_asset_precision"
	GetTransaction    string = "get_transaction"
	BuildTransaction  string = "build_transaction"
	ListTransactions  string = "list_transactions"
	IsOnline          string = "is_online"
	SetOnlineMode     string = "set_online_mode"
	SetOfflineMode    string = "set_offline_mode"
	SignData          string = "sign_data"
	EstimateFees      string = "estimate_fees"

	// Interact with wallet encrypted database
	GetMatchingKeys string = "get_matching_keys"
	GetValueFromKey string = "get_value_from_key"
	Store           string = "store"
	Delete          string = "delete"
	HasKey          string = "has_key"
	QueryDB         string = "query_db"
)

// Events
const (
	NewTopoheight  string = `new_topo_height`
	NewAsset       string = `new_asset`
	NewTransaction string = `new_transaction`
	BalanceChanged string = `balance_changed`
	//Rescan         string = `rescan`
	Online  string = `online`
	Offline string = `offline`
)
