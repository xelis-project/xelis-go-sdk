package daemon

import "github.com/xelis-project/xelis-go-sdk/lib"

type GetTopoHeightRangeParams struct {
	StartTopoheight uint64 `json:"start_topoheight"`
	EndTopoheight   uint64 `json:"end_topoheight"`
}

type GetBlockAtTopoheightParams struct {
	Topoheight uint64 `json:"topoheight"`
	IncludeTxs bool   `json:"include_txs"`
}

type GetBlocksAtHeightParams struct {
	Height     uint64 `json:"height"`
	IncludeTxs bool   `json:"include_txs"`
}

type GetBlockByHashParams struct {
	Hash       string `json:"hash"`
	IncludeTxs bool   `json:"include_txs"`
}

type GetTopBlockParams struct {
	IncludeTxs bool `json:"include_txs"`
}

type GetBalanceParams struct {
	Address string `json:"address"`
	Asset   string `json:"asset"`
}

type Balance struct {
	Balance            uint64 `json:"balance"`
	PreviousTopoheight uint64 `json:"previous_topoheight"`
}

type GetBalanceResult struct {
	Version    Balance `json:"version"`
	Topoheight uint64  `json:"topoheight"`
}

type GetBalanceAtTopoheightParams struct {
	Address    string `json:"address"`
	Asset      string `json:"asset"`
	Topoheight uint64 `json:"topoheight"`
}

type GetHeightRangeParams struct {
	StartHeight uint64 `json:"start_height"`
	EndHeight   uint64 `json:"end_height"`
}

type GetTransactionsParams struct {
	TxHashes []string `json:"tx_hashes"`
}

type P2PStatusResult struct {
	BestTopoheight uint64 `json:"best_topoheight"`
	MaxPeers       uint64 `json:"max_peers"`
	OurTopoheight  uint64 `json:"our_topoheight"`
	PeerCount      uint64 `json:"peer_count"`
	PeerId         uint64 `json:"peer_id"`
	Tag            string `json:"tag"`
}

type GetAssetsParams = GetAccountsParams

type GetAccountsParams struct {
	Skip              uint64 `json:"skip"`
	Maximum           uint64 `json:"maximum"`
	MinimumTopoheight uint64 `json:"minimum_topoheight"`
	MaximumTopoheight uint64 `json:"maximum_topoheight"`
}

type Block struct {
	BlockType            string   `json:"block_type"`
	CumulativeDifficulty uint64   `json:"cumulative_difficulty"`
	Difficulty           uint64   `json:"difficulty"`
	ExtraNonce           string   `json:"extra_nonce"`
	Hash                 string   `json:"hash"`
	Height               uint64   `json:"height"`
	Miner                string   `json:"miner"`
	Nonce                uint64   `json:"nonce"`
	Reward               uint64   `json:"reward"`
	Supply               uint64   `json:"supply"`
	Timestamp            uint64   `json:"timestamp"`
	Tips                 []string `json:"tips"`
	Topoheight           uint64   `json:"topoheight"`
	TotalFees            uint64   `json:"total_fees"`
	TotalSizeInBytes     uint64   `json:"total_size_in_bytes"`
	TxsHashes            []string `json:"txs_hashes"`
	Version              uint64   `json:"version"`
}

type Transfer struct {
	Amount    uint64 `json:"amount"`
	Asset     string `json:"asset"`
	ExtraData string `json:"extra_data"`
	To        string `json:"to"`
}

type Burn struct {
	Asset  string `json:"asset"`
	Amount uint64 `json:"amount"`
}

type CallContract struct {
	Contract string `json:"contract"`
}

type TransactionData struct {
	Transfers      []Transfer `json:"transfers"`
	Burn           Burn       `json:"burn"`
	CallContract   string     `json:"call_contract"`
	DeployContract string     `json:"deploy_contract"`
}

type Transaction struct {
	Blocks          []string        `json:"blocks"`
	Hash            string          `json:"hash"`
	Data            TransactionData `json:"data"`
	Fee             uint64          `json:"fee"`
	Nonce           uint64          `json:"nonce"`
	Owner           string          `json:"owner"`
	Signature       string          `json:"signature"`
	ExecutedInBlock string          `json:"executed_in_block"`
	Version         uint64          `json:"version"`
	FirstSeen       uint64          `json:"first_seen"`
}

type NewBlockResult struct {
	BlockType            string   `json:"block_type"`
	CumulativeDifficulty uint64   `json:"cumulative_difficulty"`
	Difficulty           uint64   `json:"difficulty"`
	Event                string   `json:"event"`
	ExtraNonce           string   `json:"extra_nonce"`
	Hash                 string   `json:"hash"`
	Height               uint64   `json:"height"`
	Miner                string   `json:"miner"`
	Nonce                uint64   `json:"nonce"`
	Reward               uint64   `json:"reward"`
	Supply               uint64   `json:"supply"`
	Timestamp            uint64   `json:"timestamp"`
	Tips                 []string `json:"tips"`
	Topoheight           uint64   `json:"topoheight"`
	TotalFees            uint64   `json:"total_fees"`
	TotalSizeInBytes     uint64   `json:"total_size_in_bytes"`
	TxsHashes            []string `json:"txs_hashes"`
}

type GetInfoResult struct {
	AverageBlocktime uint64 `json:"average_block_time"`
	BlockReward      uint64 `json:"block_reward"`
	BlockTimeTarget  uint64 `json:"block_time_target"`
	Difficulty       uint64 `json:"difficulty"`
	Height           uint64 `json:"height"`
	MempoolSize      uint64 `json:"mempool_size"`
	NativeSupply     uint64 `json:"native_supply"`
	Network          string `json:"network"`
	PrunedTopoheight uint64 `json:"pruned_topoheight"`
	Stableheight     uint64 `json:"stableheight"`
	TopHash          string `json:"top_hash"`
	Topoheight       uint64 `json:"topoheight"`
	Version          string `json:"version"`
}

type GetBlockTemplateResult struct {
	Difficulty uint64 `json:"difficulty"`
	Height     uint64 `json:"height"`
	Template   string `json:"template"`
}

type GetNonceResult struct {
	Nonce              uint64 `json:"nonce"`
	PreviousTopoheight uint64 `json:"previous_topoheight"`
	Topoheight         uint64 `json:"topoheight"`
}

type Peer struct {
	Addr                 string `json:"addr"`
	CumulativeDifficulty uint64 `json:"cumulative_difficulty"`
	Height               uint64 `json:"height"`
	Id                   uint64 `json:"id"`
	LastPing             uint64 `json:"last_ping"`
	PrunedTopoheight     uint64 `json:"pruned_topoheight"`
	Tag                  string
	TopBlockHash         string `json:"top_block_hash"`
	Topoheight           uint64 `json:"topoheight"`
	Version              string `json:"version"`
}

type MiningHistory struct {
	Reward uint64 `json:"reward"`
}

type AmountHistory struct {
	Amount uint64 `json:"amount"`
}

type Asset struct {
	Topoheight uint64 `json:"topoheight"`
	Decimals   int    `json:"decimals"`
}

type Fee struct {
	FeePercentage int    `json:"fee_percentage"`
	Height        uint64 `json:"height"`
}

type SizeOnDisk struct {
	SizeBytes     uint64 `json:"size_bytes"`
	SizeFormatted string `json:"size_formatted"`
}

type IsTxExecutedInBlockParams struct {
	TxHash    string `json:"tx_hash"`
	BlockHash string `json:"block_hash"`
}

type AccountHistory struct {
	Topoheight     uint64        `json:"topoheight"`
	BlockTimestamp uint64        `json:"block_timestamp"`
	Hash           string        `json:"hash"`
	Mining         MiningHistory `json:"mining"`
	Burn           AmountHistory `json:"burn"`
	Outgoing       AmountHistory `json:"outgoing"`
	Incoming       AmountHistory `json:"incoming"`
}

const (
	NewBlock                  lib.RPCEvent = `new_block`
	TransactionAddedInMempool lib.RPCEvent = `transaction_added_in_mempool`
	TransactionExecuted       lib.RPCEvent = `transaction_executed`
	BlockOrdered              lib.RPCEvent = `block_ordered`
)

const (
	GetVersion                 lib.RPCMethod = "get_version"
	GetInfo                    lib.RPCMethod = "get_info"
	GetHeight                  lib.RPCMethod = "get_height"
	GetTopoHeight              lib.RPCMethod = "get_topoheight"
	GetStableHeight            lib.RPCMethod = "get_stableheight"
	GetBlockTemplate           lib.RPCMethod = "get_block_template"
	GetBlockAtTopoheight       lib.RPCMethod = "get_block_at_topoheight"
	GetBlocksAtHeight          lib.RPCMethod = "get_blocks_at_height"
	GetBlockByHash             lib.RPCMethod = "get_block_by_hash"
	GetTopBlock                lib.RPCMethod = "get_top_block"
	GetNonce                   lib.RPCMethod = "get_nonce"
	HasNonce                   lib.RPCMethod = "has_nonce"
	GetBalance                 lib.RPCMethod = "get_balance"
	HasBalance                 lib.RPCMethod = "has_balance"
	GetBalanceAtTopoheight     lib.RPCMethod = "get_balance_at_topoheight"
	GetAsset                   lib.RPCMethod = "get_asset"
	GetAssets                  lib.RPCMethod = "get_assets"
	CountAssets                lib.RPCMethod = "count_assets"
	CountTransactions          lib.RPCMethod = "count_transactions"
	CountAccounts              lib.RPCMethod = "count_accounts"
	GetTips                    lib.RPCMethod = "get_tips"
	P2PStatus                  lib.RPCMethod = "p2p_status"
	GetDAGOrder                lib.RPCMethod = "get_dag_order"
	SubmitBlock                lib.RPCMethod = "submit_block"
	SubmitTransaction          lib.RPCMethod = "submit_transaction"
	GetMempool                 lib.RPCMethod = "get_mempool"
	GetTransaction             lib.RPCMethod = "get_transaction"
	GetTransactions            lib.RPCMethod = "get_transactions"
	GetBlocksRangeByHeight     lib.RPCMethod = "get_blocks_range_by_height"
	GetBlocksRangeByTopoheight lib.RPCMethod = "get_blocks_range_by_topoheight"
	GetAccounts                lib.RPCMethod = "get_accounts"
	GetAccountHistory          lib.RPCMethod = "get_account_history"
	GetAccountAssets           lib.RPCMethod = "get_account_assets"
	GetPeers                   lib.RPCMethod = "get_peers"
	GetDevFeeThresholds        lib.RPCMethod = "get_dev_fee_thresholds"
	GetSizeOnDisk              lib.RPCMethod = "get_size_on_disk"
	IsTxExecutedInBlock        lib.RPCMethod = "is_tx_executed_in_block"
)
