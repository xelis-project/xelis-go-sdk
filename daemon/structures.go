package daemon

import "encoding/json"

type RPCRequest struct {
	ID      int64                  `json:"id"`
	JSONRPC string                 `json:"jsonrpc"`
	Method  RPCMethod              `json:"method"`
	Params  map[string]interface{} `json:"params,omitempty"`
}

type RPCResponse struct {
	ID     int64           `json:"id"`
	Result json.RawMessage `json:"result,omitempty"`
	Error  *RPCError       `json:"error,omitempty"`
}

type RPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type GetRangeParams struct {
	StartTopoheight uint64 `json:"start_topoheight"`
	EndTopoheight   uint64 `json:"end_topoheight"`
}

type GetTransactionsParams struct {
	TxHashes []string `json:"tx_hashes"`
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
}

type Transfer struct {
	Amount    int64  `json:"amount"`
	Asset     string `json:"asset"`
	ExtraData string `json:"extra_data"`
	To        string `json:"to"`
}

type Data struct {
	Transfer []Transfer `json:"Transfer"`
}

type Transaction struct {
	Blocks    []string `json:"blocks"`
	Hash      string   `json:"hash"`
	Data      Data     `json:"data"`
	Fee       int64    `json:"fee"`
	Nonce     int64    `json:"nonce"`
	Owner     string   `json:"owner"`
	Signature string   `json:"signature"`
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
	BlockTimeTarget uint64 `json:"block_time_target"`
	Difficulty      uint64 `json:"difficulty"`
	Height          uint64 `json:"height"`
	MempoolSize     uint64 `json:"mempool_size"`
	NativeSupply    uint64 `json:"native_supply"`
	StableHeight    uint64 `json:"stableheight"`
	TopHash         string `json:"top_hash"`
	Version         string `json:"version"`
	Network         string `json:"network"`
	TopoHeight      uint64 `json:"topoheight"`
}

type RPCEvent string

const (
	NewBlock                  RPCEvent = `NewBlock`
	TransactionAddedInMempool RPCEvent = `TransactionAddedInMempool`
	TransactionExecuted       RPCEvent = `TransactionExecuted`
	BlockOrdered              RPCEvent = `BlockOrdered`
)

type RPCMethod string

const (
	GetInfo                RPCMethod = "get_info"
	GetHeight              RPCMethod = "get_height"
	GetTopoHeight          RPCMethod = "get_topoheight"
	GetStableHeight        RPCMethod = "get_stableheight"
	GetBlockTemplate       RPCMethod = "get_block_template"
	GetBlockAtTopoHeight   RPCMethod = "get_block_at_topoheight"
	GetBlocksAtHeight      RPCMethod = "get_blocks_at_height"
	GetBlockByHash         RPCMethod = "get_block_by_hash"
	GetTopBlock            RPCMethod = "get_top_block"
	GetNonce               RPCMethod = "get_nonce"
	GetLastBalance         RPCMethod = "get_last_balance"
	GetBalanceAtTopoHeight RPCMethod = "get_balance_at_topoheight"
	GetAssets              RPCMethod = "get_assets"
	CountTransactions      RPCMethod = "count_transactions"
	GetTips                RPCMethod = "get_tips"
	P2PStatus              RPCMethod = "p2p_status"
	GetDAGOrder            RPCMethod = "get_dag_order"
	GetMempool             RPCMethod = "get_mempool"
	GetTransaction         RPCMethod = "get_transaction"
	GetTransactions        RPCMethod = "get_transactions"
	GetBlocks              RPCMethod = "get_blocks"
)
