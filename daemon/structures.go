package daemon

type RPCRequest struct {
	ID      int64                  `json:"id"`
	JSONRPC string                 `json:"jsonrpc"`
	Method  string                 `json:"method"`
	Params  map[string]interface{} `json:"params,omitempty"`
}

type RPCResponse struct {
	ID     int64       `json:"id"`
	Result interface{} `json:"result,omitempty"`
	Error  RPCError    `json:"error,omitempty"`
}

type RPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
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
