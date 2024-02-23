package daemon

type GetTopoheightRangeParams struct {
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
	ExtraData []byte `json:"extra_data"`
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

type TransactionExecutedResult struct {
	BlockHash  string `json:"block_hash"`
	Topoheight uint64 `json:"topoheight"`
	TxHash     string `json:"tx_hash"`
}

type PeerDirection string

const (
	PeerIn   PeerDirection = "In"
	PeerOut  PeerDirection = "Out"
	PeerBoth PeerDirection = "Both"
)

type Peer struct {
	Id                   uint64                   `json:"id"`
	CumulativeDifficulty uint64                   `json:"cumulative_difficulty"`
	PrunedTopoheight     uint64                   `json:"pruned_topoheight"`
	ConnectedOn          uint64                   `json:"connected_on"`
	Height               uint64                   `json:"height"`
	LocalPort            int                      `json:"local_port"`
	TopBlockHash         string                   `json:"top_block_hash"`
	Addr                 string                   `json:"addr"`
	LastPing             uint64                   `json:"last_ping"`
	Tag                  string                   `json:"tag"`
	Topoheight           uint64                   `json:"topoheight"`
	Peers                map[string]PeerDirection `json:"peers"`
	Version              string                   `json:"version"`
}

const (
	NewBlock                  string = `new_block`
	TransactionAddedInMempool string = `transaction_added_in_mempool`
	TransactionExecuted       string = `transaction_executed`
	BlockOrdered              string = `block_ordered`
	PeerConnected             string = `peer_connected`
	PeerDisconnected          string = `peer_disconnect`
	PeerStateUpdated          string = `peer_state_updated`
)

const (
	GetVersion                 string = "get_version"
	GetInfo                    string = "get_info"
	GetHeight                  string = "get_height"
	GetTopoHeight              string = "get_topoheight"
	GetStableHeight            string = "get_stableheight"
	GetBlockTemplate           string = "get_block_template"
	GetBlockAtTopoheight       string = "get_block_at_topoheight"
	GetBlocksAtHeight          string = "get_blocks_at_height"
	GetBlockByHash             string = "get_block_by_hash"
	GetTopBlock                string = "get_top_block"
	GetNonce                   string = "get_nonce"
	HasNonce                   string = "has_nonce"
	GetBalance                 string = "get_balance"
	HasBalance                 string = "has_balance"
	GetBalanceAtTopoheight     string = "get_balance_at_topoheight"
	GetAsset                   string = "get_asset"
	GetAssets                  string = "get_assets"
	CountAssets                string = "count_assets"
	CountTransactions          string = "count_transactions"
	CountAccounts              string = "count_accounts"
	GetTips                    string = "get_tips"
	P2PStatus                  string = "p2p_status"
	GetDAGOrder                string = "get_dag_order"
	SubmitBlock                string = "submit_block"
	SubmitTransaction          string = "submit_transaction"
	GetMempool                 string = "get_mempool"
	GetTransaction             string = "get_transaction"
	GetTransactions            string = "get_transactions"
	GetBlocksRangeByHeight     string = "get_blocks_range_by_height"
	GetBlocksRangeByTopoheight string = "get_blocks_range_by_topoheight"
	GetAccounts                string = "get_accounts"
	GetAccountHistory          string = "get_account_history"
	GetAccountAssets           string = "get_account_assets"
	GetPeers                   string = "get_peers"
	GetDevFeeThresholds        string = "get_dev_fee_thresholds"
	GetSizeOnDisk              string = "get_size_on_disk"
	IsTxExecutedInBlock        string = "is_tx_executed_in_block"
)
