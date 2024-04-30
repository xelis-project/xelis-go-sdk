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

type BalanceType string

var (
	BalanceInput  BalanceType = `input`
	BalanceOutput BalanceType = `output`
	BalanceBoth   BalanceType = `both`
)

type EncryptedBalance struct {
	Commitment []byte `json:"commitment"`
	Handle     []byte `json:"handle"`
}

type VersionedBalance struct {
	BalanceType        BalanceType      `json:"balance_type"`
	FinalBalance       EncryptedBalance `json:"final_balance"`
	OutputBalance      EncryptedBalance `json:"output_balance"`
	PreviousTopoheight uint64           `json:"previous_topoheight"`
}

type GetBalanceResult struct {
	Version    VersionedBalance `json:"version"`
	Topoheight uint64           `json:"topoheight"`
}

type GetNonceAtTopoheightParams struct {
	Address    string `json:"address"`
	Topoheight uint64 `json:"topoheight"`
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
	CumulativeDifficulty string   `json:"cumulative_difficulty"`
	Difficulty           string   `json:"difficulty"`
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
	Asset           string  `json:"asset"`
	ExtraData       *[]byte `json:"extra_data"`
	Destination     string  `json:"destination"`
	Commitment      []byte  `json:"commitment"`
	SenderHandle    []byte  `json:"sender_handle"`
	ReceiverHandle  []byte  `json:"receiver_handle"`
	CTValidityProof Proof   `json:"ct_validity_proof"`
}

type Burn struct {
	Asset  string `json:"asset"`
	Amount uint64 `json:"amount"`
}

type CallContract struct {
	Contract string `json:"contract"`
}

type TransactionData struct {
	Transfers []Transfer `json:"transfers"`
	Burn      *Burn      `json:"burn"`
	// CallContract   string     `json:"call_contract"`
	// DeployContract string     `json:"deploy_contract"`
}

type Reference struct {
	Hash       string `json:"hash"`
	Topoheight uint64 `json:"topoheight"`
}

type Proof struct {
	Y_0 []byte `json:"Y_0"`
	Y_1 []byte `json:"Y_1"`
	Z_R []byte `json:"z_r"`
	Z_X []byte `json:"z_x"`
}

type EqProof struct {
	Y_0 []byte `json:"Y_0"`
	Y_1 []byte `json:"Y_1"`
	Y_2 []byte `json:"Y_2"`
	Z_R []byte `json:"z_r"`
	Z_S []byte `json:"z_s"`
	Z_X []byte `json:"z_x"`
}

type SourceCommitment struct {
	Commitment []byte  `json:"commitment"`
	Proof      EqProof `json:"proof"`
	Asset      string  `json:"asset"`
}

type Transaction struct {
	Blocks            []string           `json:"blocks"`
	Hash              string             `json:"hash"`
	Data              TransactionData    `json:"data"`
	Fee               uint64             `json:"fee"`
	Nonce             uint64             `json:"nonce"`
	Source            string             `json:"source"`
	Reference         Reference          `json:"reference"`
	SourceCommitments []SourceCommitment `json:"source_commitments"`
	RangeProof        []byte             `json:"range_proof"`
	Signature         string             `json:"signature"`
	ExecutedInBlock   string             `json:"executed_in_block"`
	Version           uint64             `json:"version"`
	FirstSeen         uint64             `json:"first_seen"`
	InMempool         bool               `json:"in_mempool"`
}

type GetInfoResult struct {
	AverageBlocktime uint64 `json:"average_block_time"`
	BlockReward      uint64 `json:"block_reward"`
	BlockTimeTarget  uint64 `json:"block_time_target"`
	Difficulty       string `json:"difficulty"`
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
	Template   string `json:"template"`
	Height     uint64 `json:"height"`
	Topoheight uint64 `json:"topoheight"`
	Difficulty string `json:"difficulty"`
}

type SubmitBlockParams struct {
	BlockTemplate string  `json:"block_template"`
	MinerWork     *string `json:"miner_work,omitempty"`
}

type GetNonceResult struct {
	Nonce              uint64 `json:"nonce"`
	PreviousTopoheight uint64 `json:"previous_topoheight"`
	Topoheight         uint64 `json:"topoheight"`
}

type MiningHistory struct {
	Reward uint64 `json:"reward"`
}

type BurnHistory struct {
	Amount uint64 `json:"amount"`
}

type OutgoingHistory struct {
	To string `json:"to"`
}

type IncomingHistory struct {
	From string `json:"from"`
}

type Asset struct {
	Topoheight uint64 `json:"topoheight"`
	Decimals   int    `json:"decimals"`
}

type AssetWithData struct {
	Asset      string `json:"asset"`
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
	Topoheight     uint64           `json:"topoheight"`
	BlockTimestamp uint64           `json:"block_timestamp"`
	Hash           string           `json:"hash"`
	Mining         *MiningHistory   `json:"mining"`
	Burn           *BurnHistory     `json:"burn"`
	Outgoing       *OutgoingHistory `json:"outgoing"`
	Incoming       *IncomingHistory `json:"incoming"`
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
	CumulativeDifficulty string                   `json:"cumulative_difficulty"`
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

type GetPeersResult struct {
	Peers       []Peer `json:"peers"`
	TotalPeers  int    `json:"total_peers"`
	HiddenPeers int    `json:"hidden_peers"`
}

type IsAccountRegisteredParams struct {
	Address        string `json:"address"`
	InStableHeight bool   `json:"in_stable_height"`
}

type GetDifficultyResult struct {
	Difficulty        string `json:"difficulty"`
	Hashrate          string `json:"hashrate"`
	HashrateFormatted string `json:"hashrate_formatted"`
}

type ValidateAddressParams struct {
	Address         string `json:"address"`
	AllowIntegrated bool   `json:"allow_integrated"`
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
	GetVersion                       string = "get_version"
	GetInfo                          string = "get_info"
	GetHeight                        string = "get_height"
	GetTopoHeight                    string = "get_topoheight"
	GetStableHeight                  string = "get_stableheight"
	GetBlockTemplate                 string = "get_block_template"
	GetBlockAtTopoheight             string = "get_block_at_topoheight"
	GetBlocksAtHeight                string = "get_blocks_at_height"
	GetBlockByHash                   string = "get_block_by_hash"
	GetTopBlock                      string = "get_top_block"
	GetNonce                         string = "get_nonce"
	HasNonce                         string = "has_nonce"
	GetNonceAtTopoheight             string = "get_nonce_at_topoheight"
	GetBalance                       string = "get_balance"
	HasBalance                       string = "has_balance"
	GetBalanceAtTopoheight           string = "get_balance_at_topoheight"
	GetAsset                         string = "get_asset"
	GetAssets                        string = "get_assets"
	CountAssets                      string = "count_assets"
	CountTransactions                string = "count_transactions"
	CountAccounts                    string = "count_accounts"
	GetTips                          string = "get_tips"
	P2PStatus                        string = "p2p_status"
	GetDAGOrder                      string = "get_dag_order"
	SubmitBlock                      string = "submit_block"
	SubmitTransaction                string = "submit_transaction"
	GetMempool                       string = "get_mempool"
	GetTransaction                   string = "get_transaction"
	GetTransactions                  string = "get_transactions"
	GetBlocksRangeByHeight           string = "get_blocks_range_by_height"
	GetBlocksRangeByTopoheight       string = "get_blocks_range_by_topoheight"
	GetAccounts                      string = "get_accounts"
	GetAccountHistory                string = "get_account_history"
	GetAccountAssets                 string = "get_account_assets"
	GetPeers                         string = "get_peers"
	GetDevFeeThresholds              string = "get_dev_fee_thresholds"
	GetSizeOnDisk                    string = "get_size_on_disk"
	IsTxExecutedInBlock              string = "is_tx_executed_in_block"
	GetAccountRegistrationTopoheight string = "get_account_registration_topoheight"
	IsAccountRegistered              string = "is_account_registered"
	GetDifficulty                    string = "get_difficulty"
	ValidateAddress                  string = "validate_address"
)
