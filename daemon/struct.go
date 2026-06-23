package daemon

import (
	"encoding/json"

	"github.com/xelis-project/xelis-go-sdk/transaction"
	"github.com/xelis-project/xelis-go-sdk/xvm"
)

type Transfer = transaction.Transfer
type Burn = transaction.Burn
type CallContract = transaction.CallContract
type MultiSigPayload = transaction.MultiSigPayload
type ValueCell = transaction.ValueCell
type ContractDeposit = transaction.ContractDeposit
type InvokeContractPayload = transaction.InvokeContractPayload
type InvokeConstructorPayload = transaction.InvokeConstructorPayload
type DeployContractPayload = transaction.DeployContractPayload
type TransactionType = transaction.TransactionType
type Reference = transaction.Reference
type Proof = transaction.Proof
type EqProof = transaction.EqProof
type SourceCommitment = transaction.SourceCommitment
type SignatureId = transaction.SignatureId
type MultiSig = transaction.MultiSig
type Transaction = transaction.Transaction
type TransactionResponse = transaction.TransactionResponse
type ChunkAccessType = transaction.ChunkAccessType
type ModuleChunk = transaction.ModuleChunk
type Module = transaction.Module

const (
	ChunkAccessAll      = transaction.ChunkAccessAll
	ChunkAccessInternal = transaction.ChunkAccessInternal
	ChunkAccessEntry    = transaction.ChunkAccessEntry
	ChunkAccessHook     = transaction.ChunkAccessHook
)

type RPCMethodInfo struct {
	Name   string    `json:"name"`
	Schema RPCSchema `json:"schema"`
}

type RPCSchemaResponse struct {
	Schema  string                     `json:"$schema"`
	Defs    map[string]json.RawMessage `json:"$defs"`
	Methods []RPCMethodInfo            `json:"methods"`
}

type RPCSchema struct {
	ParamsSchema  *json.RawMessage `json:"params_schema"`
	ReturnsSchema json.RawMessage  `json:"returns_schema"`
}

type SubscribeParams struct {
	Notify interface{} `json:"notify"`
}

type GetTopoheightRangeParams struct {
	StartTopoheight *uint64 `json:"start_topoheight,omitempty"`
	EndTopoheight   *uint64 `json:"end_topoheight,omitempty"`
}

type GetBlockAtTopoheightParams struct {
	Topoheight uint64 `json:"topoheight"`
	IncludeTxs bool   `json:"include_txs"`
}

type GetBlocksAtHeightParams struct {
	Height     uint64 `json:"height"`
	IncludeTxs bool   `json:"include_txs"`
}

type GetBlockSummaryAtTopoheightParams struct {
	Topoheight uint64 `json:"topoheight"`
}

type GetBlockByHashParams struct {
	Hash       string `json:"hash"`
	IncludeTxs bool   `json:"include_txs"`
}

type GetBlockSummaryByHashParams struct {
	Hash string `json:"hash"`
}

type GetBlockDifficultyByHashParams struct {
	BlockHash string `json:"block_hash"`
}

type GetBlockBaseFeeByHashParams = GetBlockDifficultyByHashParams

type GetBlockBaseFeeByHashResult struct {
	FeePerKB     uint64 `json:"fee_per_kb"`
	BlockSizeEMA uint64 `json:"block_size_ema"`
}

type BlockSummary struct {
	BlockHash            string               `json:"block_hash"`
	Topoheight           *uint64              `json:"topoheight,omitempty"`
	BlockType            BlockType            `json:"block_type"`
	Difficulty           string               `json:"difficulty"`
	Supply               *uint64              `json:"supply,omitempty"`
	Reward               *uint64              `json:"reward,omitempty"`
	MinerReward          *uint64              `json:"miner_reward,omitempty"`
	DevReward            *uint64              `json:"dev_reward,omitempty"`
	CumulativeDifficulty string               `json:"cumulative_difficulty"`
	TotalFees            *uint64              `json:"total_fees,omitempty"`
	TotalFeesBurned      *uint64              `json:"total_fees_burned,omitempty"`
	Timestamp            uint64               `json:"timestamp"`
	Height               uint64               `json:"height"`
	Miner                string               `json:"miner"`
	Transactions         []TransactionSummary `json:"transactions"`
}

type GetTopBlockParams struct {
	IncludeTxs bool `json:"include_txs"`
}

type GetBlockTemplateParams struct {
	Address string `json:"address"`
}

type GetBalanceParams struct {
	Address string `json:"address"`
	Asset   string `json:"asset"`
}

type HasBalanceParams struct {
	Address    string  `json:"address"`
	Asset      string  `json:"asset"`
	Topoheight *uint64 `json:"topoheight,omitempty"`
}

type BalanceType string

var (
	BalanceInput  BalanceType = `input`
	BalanceOutput BalanceType = `output`
	BalanceBoth   BalanceType = `both`
)

type EncryptedBalance struct {
	Compressed []uint `json:"Compressed"`
}

type VersionedBalance struct {
	BalanceType        BalanceType       `json:"balance_type"`
	FinalBalance       EncryptedBalance  `json:"final_balance"`
	OutputBalance      *EncryptedBalance `json:"output_balance"`
	PreviousTopoheight *uint64           `json:"previous_topoheight"`
}

type RPCVersionedBalance struct {
	Topoheight uint64 `json:"topoheight"`
	VersionedBalance
}

type GetBalanceResult struct {
	Version    VersionedBalance `json:"version"`
	Topoheight uint64           `json:"topoheight"`
}

type GetStableBalanceResult struct {
	StableTopoheight uint64           `json:"stable_topoheight"`
	StableBlockHash  string           `json:"stable_block_hash"`
	Topoheight       uint64           `json:"topoheight"`
	Version          VersionedBalance `json:"version"`
}

type ExistResult struct {
	Exist bool `json:"exist"`
}

type GetNonceAtTopoheightParams struct {
	Address    string `json:"address"`
	Topoheight uint64 `json:"topoheight"`
}

type GetNonceParams struct {
	Address string `json:"address"`
}

type HasNonceParams struct {
	Address    string  `json:"address"`
	Topoheight *uint64 `json:"topoheight,omitempty"`
}

type GetBalanceAtTopoheightParams struct {
	Address    string `json:"address"`
	Asset      string `json:"asset"`
	Topoheight uint64 `json:"topoheight"`
}

type GetBalancesAtMaximumTopoheightParams struct {
	Address           string   `json:"address"`
	Assets            []string `json:"assets"`
	MaximumTopoheight uint64   `json:"maximum_topoheight"`
}

type GetHeightRangeParams struct {
	StartHeight *uint64 `json:"start_height,omitempty"`
	EndHeight   *uint64 `json:"end_height,omitempty"`
}

type SubmitTransactionParams struct {
	Data string `json:"data"`
}

type GetTransactionParams struct {
	Hash string `json:"hash"`
}

type GetTransactionsParams struct {
	TxHashes []string `json:"tx_hashes"`
}

type TransactionSummary struct {
	Hash   string `json:"hash"`
	Source string `json:"source"`
	Fee    uint64 `json:"fee"`
	Size   uint64 `json:"size"`
}

type P2PStatusResult struct {
	BestTopoheight   uint64  `json:"best_topoheight"`
	MaxPeers         uint64  `json:"max_peers"`
	MedianTopoheight uint64  `json:"median_topoheight"`
	OurTopoheight    uint64  `json:"our_topoheight"`
	PeerCount        uint64  `json:"peer_count"`
	PeerId           uint64  `json:"peer_id"`
	Tag              *string `json:"tag"`
}

type GetP2PBlockPropagationParams struct {
	Hash     string `json:"hash"`
	Incoming *bool  `json:"incoming,omitempty"`
	Outgoing *bool  `json:"outgoing,omitempty"`
}

type P2PBlockPropagationResult struct {
	FirstSeen    *uint64                       `json:"first_seen"`
	ProcessingAt *uint64                       `json:"processing_at"`
	Peers        map[string]TimedPeerDirection `json:"peers"`
}

type GetAssetsParams = GetAccountsParams

type GetAssetParams struct {
	Asset string `json:"asset"`
}

type GetAssetSupplyAtTopoheightParams struct {
	Asset      string `json:"asset"`
	Topoheight uint64 `json:"topoheight"`
}

type VersionedUint64 struct {
	Topoheight         uint64  `json:"topoheight"`
	Data               uint64  `json:"data"`
	PreviousTopoheight *uint64 `json:"previous_topoheight"`
}

type VersionedUint64AtTopoheight struct {
	Data               uint64  `json:"data"`
	PreviousTopoheight *uint64 `json:"previous_topoheight"`
}

type GetAccountsParams struct {
	Skip              *uint64 `json:"skip,omitempty"`
	Maximum           *uint64 `json:"maximum,omitempty"`
	MinimumTopoheight *uint64 `json:"minimum_topoheight,omitempty"`
	MaximumTopoheight *uint64 `json:"maximum_topoheight,omitempty"`
}

type BlockType string

const (
	BlockSync     BlockType = "Sync"
	BlockSide     BlockType = "Side"
	BlockOrphaned BlockType = "Orphaned"
	BlockNormal   BlockType = "Normal"
)

type Block struct {
	Hash                 string        `json:"hash"`
	Topoheight           *uint64       `json:"topoheight"`
	BlockType            BlockType     `json:"block_type"`
	Difficulty           string        `json:"difficulty"`
	Supply               *uint64       `json:"supply"`
	Reward               *uint64       `json:"reward"` // full reward miner_reward + dev_reward
	MinerReward          *uint64       `json:"miner_reward"`
	DevReward            *uint64       `json:"dev_reward"`
	CumulativeDifficulty string        `json:"cumulative_difficulty"`
	TotalFees            *uint64       `json:"total_fees"`
	TotalFeesBurned      *uint64       `json:"total_fees_burned"`
	TotalSizeInBytes     uint64        `json:"total_size_in_bytes"`
	Version              BlockVersion  `json:"version"`
	Tips                 []string      `json:"tips"`
	Timestamp            uint64        `json:"timestamp"`
	Height               uint64        `json:"height"`
	Nonce                uint64        `json:"nonce"`
	ExtraNonce           string        `json:"extra_nonce"`
	Miner                string        `json:"miner"`
	TxsHashes            []string      `json:"txs_hashes"`
	Transactions         []Transaction `json:"transactions"` // if include_txs is true in params
}

type GetMempoolResult struct {
	Total        uint64                `json:"total"`
	Transactions []TransactionResponse `json:"transactions"`
}

type GetMempoolParams struct {
	Maximum *uint64 `json:"maximum,omitempty"`
	Skip    *uint64 `json:"skip,omitempty"`
}

type MempoolTransactionSummary struct {
	Hash      string `json:"hash"`
	Source    string `json:"source"`
	Fee       uint64 `json:"fee"`
	FirstSeen uint64 `json:"first_seen"`
	Size      uint64 `json:"size"`
	FeePerKB  uint64 `json:"fee_per_kb"`
}

type GetMempoolSummaryResult struct {
	Total        uint64                      `json:"total"`
	Transactions []MempoolTransactionSummary `json:"transactions"`
}

type PredicatedBaseFeeResult struct {
	FeePerKB           uint64 `json:"fee_per_kb"`
	PredicatedFeePerKB uint64 `json:"predicated_fee_per_kb"`
}

type BlockVersion string

// hardforks
const BlockV0 BlockVersion = "V0" // genesis
const BlockV1 BlockVersion = "V1" // pow algo, diff adjust
const BlockV2 BlockVersion = "V2" // multisig, p2p upgrade
const BlockV3 BlockVersion = "V3" // smart contracts
const BlockV4 BlockVersion = "V4"
const BlockV5 BlockVersion = "V5"
const BlockV6 BlockVersion = "V6"

type Network string

const NetworkDev Network = "Dev"
const NetworkTestnet Network = "Testnet"
const NetworkStagenet Network = "Stagenet"
const NetworkMainnet Network = "Mainnet"

type GetInfoResult struct {
	Height            uint64       `json:"height"`
	Topoheight        uint64       `json:"topoheight"`
	Stableheight      uint64       `json:"stableheight"`
	StableTopoheight  uint64       `json:"stable_topoheight"`
	PrunedTopoheight  *uint64      `json:"pruned_topoheight"`
	TopBlockHash      string       `json:"top_block_hash"`
	CirculatingSupply uint64       `json:"circulating_supply"`
	BurnedSupply      uint64       `json:"burned_supply"`
	EmittedSupply     uint64       `json:"emitted_supply"`
	MaximumSupply     uint64       `json:"maximum_supply"`
	Difficulty        string       `json:"difficulty"`
	BlockTimeTarget   uint64       `json:"block_time_target"`
	AverageBlockTime  uint64       `json:"average_block_time"`
	BlockReward       uint64       `json:"block_reward"`
	DevReward         uint64       `json:"dev_reward"`
	MinerReward       uint64       `json:"miner_reward"`
	MempoolSize       uint64       `json:"mempool_size"`
	Version           string       `json:"version"`
	Network           Network      `json:"network"`
	BlockVersion      BlockVersion `json:"block_version"`
}

type AlgorithmVersion = string

const AlgorithmV1 AlgorithmVersion = "xel/v1"
const AlgorithmV2 AlgorithmVersion = "xel/v2"
const AlgorithmV3 AlgorithmVersion = "xel/v3"

type GetBlockTemplateResult struct {
	Template   string           `json:"template"`
	Algorithm  AlgorithmVersion `json:"algorithm"`
	Height     uint64           `json:"height"`
	Topoheight uint64           `json:"topoheight"`
	Difficulty string           `json:"difficulty"`
}

type SubmitBlockParams struct {
	BlockTemplate string  `json:"block_template"`
	MinerWork     *string `json:"miner_work,omitempty"`
}

type GetMinerWorkParams struct {
	Template string  `json:"template"`
	Address  *string `json:"address,omitempty"`
}

type GetMinerWorkResult struct {
	MinerWork  string `json:"miner_work"`
	Algorithm  string `json:"algorithm"`
	Height     uint64 `json:"height"`
	Difficulty string `json:"difficulty"`
	Topoheight uint64 `json:"topoheight"`
}

type GetNonceResult struct {
	Nonce              uint64  `json:"nonce"`
	PreviousTopoheight *uint64 `json:"previous_topoheight"`
	Topoheight         uint64  `json:"topoheight"`
}

type VersionedNonce struct {
	Nonce              uint64  `json:"nonce"`
	PreviousTopoheight *uint64 `json:"previous_topoheight"`
}

type MiningHistory struct {
	Reward uint64 `json:"reward"`
}

type BurnHistory struct {
	Asset  string `json:"asset"`
	Amount uint64 `json:"amount"`
}

type OutgoingHistory struct {
	Asset string `json:"asset"`
	To    string `json:"to"`
}

type IncomingHistory struct {
	Asset string `json:"asset"`
	From  string `json:"from"`
}

type MultiSigHistory struct {
	Participants []string `json:"participants"`
	Threshold    uint8    `json:"threshold"`
}

type InvokeContractHistory struct {
	Contract string   `json:"contract"`
	EntryID  uint16   `json:"entry_id"`
	Deposits []string `json:"deposits"`
}

type DeployContractHistory struct {
	Deposits []string `json:"deposits"`
}

type FromContractHistory struct {
	Contract string `json:"contract"`
	Asset    string `json:"asset"`
	Amount   uint64 `json:"amount"`
}

type AssetData struct {
	Asset      string      `json:"asset"`
	Topoheight uint64      `json:"topoheight"`
	Decimals   int         `json:"decimals"`
	Name       string      `json:"name"`
	Ticker     string      `json:"ticker"`
	MaxSupply  interface{} `json:"max_supply"`
	Owner      interface{} `json:"owner"`
}

type Fee struct {
	FeePercentage uint64 `json:"fee_percentage"`
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
	Topoheight     uint64                 `json:"topoheight"`
	BlockTimestamp uint64                 `json:"block_timestamp"`
	Hash           string                 `json:"hash"`
	Mining         *MiningHistory         `json:"mining"`
	Burn           *BurnHistory           `json:"burn"`
	Outgoing       *OutgoingHistory       `json:"outgoing"`
	Incoming       *IncomingHistory       `json:"incoming"`
	DevFee         *MiningHistory         `json:"dev_fee"`
	MultiSig       *MultiSigHistory       `json:"multi_sig"`
	InvokeContract *InvokeContractHistory `json:"invoke_contract"`
	DeployContract *DeployContractHistory `json:"deploy_contract"`
	FromContract   *FromContractHistory   `json:"from_contract"`
	Blob           *json.RawMessage       `json:"blob"`
}

type GetAccountHistoryParams struct {
	Address           string  `json:"address"`
	Asset             *string `json:"asset,omitempty"`
	IncomingFlow      *bool   `json:"incoming_flow,omitempty"`
	OutgoingFlow      *bool   `json:"outgoing_flow,omitempty"`
	MinimumTopoheight *uint64 `json:"minimum_topoheight,omitempty"`
	MaximumTopoheight *uint64 `json:"maximum_topoheight,omitempty"`
}

type GetAccountAssetsParams struct {
	Address string  `json:"address"`
	Skip    *uint64 `json:"skip,omitempty"`
	Maximum *uint64 `json:"maximum,omitempty"`
}

type GetAccountRegistrationParams struct {
	Address string `json:"address"`
}

type TransactionExecutedEvent struct {
	BlockHash  string `json:"block_hash"`
	TxHash     string `json:"tx_hash"`
	Topoheight uint64 `json:"topoheight"`
}

type PeerDirection string

const (
	PeerDirectionIn   PeerDirection = "In"
	PeerDirectionOut  PeerDirection = "Out"
	PeerDirectionBoth PeerDirection = "Both"
)

type TimestampMillis int64

type TimedPeerDirection struct {
	In   *In   `json:"in,omitempty"`
	Out  *Out  `json:"out,omitempty"`
	Both *Both `json:"both,omitempty"`
}

func (td *TimedPeerDirection) GetDirection() PeerDirection {
	if td.Both != nil {
		return PeerDirectionBoth
	} else if td.In != nil {
		return PeerDirectionIn
	} else if td.Out != nil {
		return PeerDirectionOut
	}
	return ""
}

type In struct {
	ReceivedAt TimestampMillis `json:"received_at"`
}

type Out struct {
	SentAt TimestampMillis `json:"sent_at"`
}

type Both struct {
	ReceivedAt TimestampMillis `json:"received_at"`
	SentAt     TimestampMillis `json:"sent_at"`
}

type Peer struct {
	Id                   uint64                        `json:"id"`
	Addr                 string                        `json:"addr"`
	LocalPort            int                           `json:"local_port"`
	Tag                  *string                       `json:"tag"`
	Version              string                        `json:"version"`
	TopBlockHash         string                        `json:"top_block_hash"`
	Topoheight           uint64                        `json:"topoheight"`
	Height               uint64                        `json:"height"`
	LastPing             uint64                        `json:"last_ping"`
	PrunedTopoheight     *uint64                       `json:"pruned_topoheight"`
	Peers                map[string]TimedPeerDirection `json:"peers"`
	CumulativeDifficulty string                        `json:"cumulative_difficulty"`
	ConnectedOn          uint64                        `json:"connected_on"`
	BytesSent            uint64                        `json:"bytes_sent"`
	BytesReceived        uint64                        `json:"bytes_recv"`
}

type GetPeersResult struct {
	Peers       []Peer `json:"peers"`
	TotalPeers  int    `json:"total_peers"`
	HiddenPeers int    `json:"hidden_peers"`
}

func (r *GetPeersResult) UnmarshalJSON(data []byte) error {
	var peers []Peer
	if err := json.Unmarshal(data, &peers); err == nil {
		r.Peers = peers
		r.TotalPeers = len(peers)
		r.HiddenPeers = 0
		return nil
	}

	type getPeersResult GetPeersResult
	var result getPeersResult
	if err := json.Unmarshal(data, &result); err != nil {
		return err
	}

	*r = GetPeersResult(result)
	return nil
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
	Address               string  `json:"address"`
	AllowIntegrated       bool    `json:"allow_integrated"`
	MaxIntegratedDataSize *uint64 `json:"max_integrated_data_size,omitempty"`
}

type ValidateAddressResult struct {
	IsIntegrated bool `json:"is_integrated"`
	IsValid      bool `json:"is_valid"`
}

type ExtractKeyFromAddressParams struct {
	Address string `json:"address"`
	AsHex   bool   `json:"as_hex"`
}

type ExtractKeyFromAddressResult struct {
	Hex   *string   `json:"hex"`
	Bytes *[32]byte `json:"bytes"`
}

type KeyToAddressParams interface{}

type SplitAddressParams struct {
	Address string `json:"address"`
}

type SplitAddressResult struct {
	Address        string      `json:"address"`
	IntegratedData interface{} `json:"integrated_data"`
	Size           uint64      `json:"size"`
}

type GetTransactionExecutorParams struct {
	Hash string `json:"hash"`
}

type GetTransactionExecutorResult struct {
	BlockTopoheight uint64 `json:"block_topoheight"`
	BlockHash       string `json:"block_hash"`
	BlockTimestamp  uint64 `json:"block_timestamp"`
}

type HasMultisigAtTopoheightParams struct {
	Address    string `json:"address"`
	Topoheight uint64 `json:"topoheight"`
}

type GetMultisigAtTopoheightParams struct {
	Address    string `json:"address"`
	Topoheight uint64 `json:"topoheight"`
}

type GetMultisigAtTopoheightResult struct {
	State MultisigState `json:"state"`
}

type GetMultisigParams struct {
	Address string `json:"address"`
}

type GetMultisigResult struct {
	State      MultisigState `json:"state"`
	Topoheight uint64        `json:"topoheight"`
}

type HasMultisigParams struct {
	Address string `json:"address"`
}

type MultisigActiveState struct {
	Participants []string `json:"participants"`
	Threshold    uint8    `json:"threshold"`
}

type MultisigState struct {
	Deleted bool
	Active  *MultisigActiveState
}

func (s *MultisigState) UnmarshalJSON(data []byte) error {
	if string(data) == `"deleted"` {
		s.Deleted = true
		s.Active = nil
		return nil
	}

	var state struct {
		Active *MultisigActiveState `json:"active"`
	}
	if err := json.Unmarshal(data, &state); err != nil {
		return err
	}

	s.Deleted = false
	s.Active = state.Active
	return nil
}

func (s MultisigState) MarshalJSON() ([]byte, error) {
	if s.Deleted {
		return []byte(`"deleted"`), nil
	}

	return json.Marshal(struct {
		Active *MultisigActiveState `json:"active"`
	}{
		Active: s.Active,
	})
}

type GetContractOutputsParams struct {
	Address    string `json:"address"`
	Topoheight uint64 `json:"topoheight"`
}

type GetContractsOutputsResult struct {
	Executions []ContractTransfersKV `json:"executions"`
}

type ContractTransfersKV struct {
	Key   ContractTransfersEntryKey `json:"key"`
	Value ContractTransfersEntry    `json:"value"`
}

type ContractTransfersEntryKey struct {
	Contract string `json:"contract"`
	Caller   string `json:"caller"`
}

type ContractTransfersEntry struct {
	Transfers map[string]uint64 `json:"transfers"`
}

type GetContractModuleParams struct {
	Contract string `json:"contract"`
}

type GetContractAssetsParams struct {
	Contract string  `json:"contract"`
	Skip     *uint64 `json:"skip,omitempty"`
	Maximum  *uint64 `json:"maximum,omitempty"`
}

type GetContractLogsParams struct {
	Caller            string  `json:"caller"`
	Contract          string  `json:"contract,omitempty"`
	MinimumTopoheight *uint64 `json:"minimum_topoheight,omitempty"`
	MaximumTopoheight *uint64 `json:"maximum_topoheight,omitempty"`
	Skip              *uint64 `json:"skip,omitempty"`
	Maximum           *uint64 `json:"maximum,omitempty"`
}

type GetContractExecutionsAtTopoheightParams struct {
	Topoheight uint64  `json:"topoheight"`
	Skip       *uint64 `json:"skip,omitempty"`
	Max        *uint64 `json:"max,omitempty"`
}

type RegisteredExecution struct {
	ExecutionHash       string `json:"execution_hash"`
	ExecutionTopoheight uint64 `json:"execution_topoheight"`
}

type ScheduledExecution struct {
	Hash       string            `json:"hash"`
	Contract   string            `json:"contract"`
	ChunkID    uint16            `json:"chunk_id"`
	Params     []xvm.ValueCell   `json:"params"`
	MaxGas     uint64            `json:"max_gas"`
	Kind       json.RawMessage   `json:"kind"`
	GasSources map[string]uint64 `json:"gas_sources"`
}

type ContractLog struct {
	Type  string          `json:"type"`
	Value json.RawMessage `json:"value,omitempty"`
}

type GetContractsParams struct {
	Skip              *uint64 `json:"skip,omitempty"`
	Maximum           *uint64 `json:"maximum,omitempty"`
	MinimumTopoheight *uint64 `json:"minimum_topoheight,omitempty"`
	MaximumTopoheight *uint64 `json:"maximum_topoheight,omitempty"`
}

type GetContractDataEntriesParams struct {
	Contract          string  `json:"contract"`
	MinimumTopoheight *uint64 `json:"minimum_topoheight,omitempty"`
	MaximumTopoheight *uint64 `json:"maximum_topoheight,omitempty"`
	Skip              *uint64 `json:"skip,omitempty"`
	Maximum           *uint64 `json:"maximum,omitempty"`
}

type GetContractTransactionsParams struct {
	Contract string  `json:"contract"`
	Skip     *uint64 `json:"skip,omitempty"`
	Maximum  *uint64 `json:"maximum,omitempty"`
}

type ContractDataEntry struct {
	Key   xvm.ValueCell `json:"key"`
	Value xvm.ValueCell `json:"value"`
}

type GetContractDataParams struct {
	Contract string        `json:"contract"`
	Key      xvm.ValueCell `json:"key"`
}

type GetContractDataAtTopoheightParams struct {
	Contract   string        `json:"contract"`
	Key        xvm.ValueCell `json:"key"`
	Topoheight uint64        `json:"topoheight"`
}

type GetContractBalanceParams struct {
	Contract string `json:"contract"`
	Asset    string `json:"asset"`
}

type GetContractBalanceAtTopoheightParams struct {
	Contract   string `json:"contract"`
	Asset      string `json:"asset"`
	Topoheight uint64 `json:"topoheight"`
}

type ContractOutputRefundGas struct {
	Amount uint64
}

type ContractOutputTransfer struct {
	Amount      uint64
	Asset       string
	Destination string
}

type ContractOutputExitCode struct {
	ExitCode uint64
}

type ContractOutputRefundDeposits struct{}

type ContractOutput interface{}

type GetContractModuleResult struct {
	Topoheight         uint64  `json:"topoheight"`
	PreviousTopoheight *uint64 `json:"previous_topoheight"`
	Data               *Module `json:"data"`
}

type GetContractDataResult struct {
	Topoheight         uint64         `json:"topoheight"`
	PreviousTopoheight *uint64        `json:"previous_topoheight"`
	Data               *xvm.ValueCell `json:"data"`
}

type GetContractDataAtTopoheightResult struct {
	PreviousTopoheight *uint64        `json:"previous_topoheight"`
	Data               *xvm.ValueCell `json:"data"`
}

type HardFork struct {
	Height             uint64       `json:"height"`
	Version            BlockVersion `json:"version"`
	Changelog          string       `json:"changelog"`
	VersionRequirement *string      `json:"version_requirement"`
}

type PruneChainParams struct {
	Topoheight uint64 `json:"topoheight"`
}

type PruneChainResult struct {
	PrunedTopoheight uint64 `json:"pruned_topoheight"`
}

type RewindChainParams struct {
	Count             uint64 `json:"count"`
	UntilStableHeight bool   `json:"until_stable_height"`
}

type RewindChainResult struct {
	Topoheight uint64   `json:"topoheight"`
	Txs        []string `json:"txs"`
}

type FeeRatesEstimated struct {
	Low     uint64 `json:"low"`
	Medium  uint64 `json:"medium"`
	High    uint64 `json:"high"`
	Default uint64 `json:"default"`
}

type MakeIntegratedAddressParams struct {
	Address        string      `json:"address"`
	IntegratedData interface{} `json:"integrated_data"`
}

type DecryptExtraDataParams struct {
	SharedKey string `json:"shared_key"`
	ExtraData []uint `json:"extra_data"`
}

type GetMempoolCacheParams struct {
	Address string `json:"address"`
}

type GetMempoolCacheResult struct {
	Min      uint64                       `json:"min"`
	Max      uint64                       `json:"max"`
	Txs      []string                     `json:"txs"`
	Balances map[string][]uint            `json:"balances"`
	Multisig *MempoolCacheMultiSigPayload `json:"multisig"`
}

type MempoolCacheMultiSigPayload struct {
	Threshold    uint8    `json:"threshold"`
	Participants [][]uint `json:"participants"`
}

type GetContractBalanceResult struct {
	Topoheight         uint64  `json:"topoheight"`
	Amount             uint64  `json:"data"`
	PreviousTopoheight *uint64 `json:"previous_topoheight"`
}

type GetContractBalanceAtTopoheightResult struct {
	Amount             uint64  `json:"data"`
	PreviousTopoheight *uint64 `json:"previous_topoheight"`
}

type BlockOrderedEvent struct {
	BlockHash  string    `json:"block_hash"`
	BlockType  BlockType `json:"block_type"`
	Topoheight uint64    `json:"topoheight"`
}

type BlockOrphanedEvent struct {
	BlockHash     string `json:"block_hash"`
	OldTopoheight uint64 `json:"old_topoheight"`
}

type StableHeightChangedEvent struct {
	PreviousStableHeight uint64 `json:"previous_stable_height"`
	NewStableHeight      uint64 `json:"new_stable_height"`
}

type StableTopoheightChangedEvent struct {
	PreviousStableTopoheight uint64 `json:"previous_stable_topoheight"`
	NewStableTopoheight      uint64 `json:"new_stable_topoheight"`
}

type PeerPeerListUpdatedEvent struct {
	PeerId   uint64   `json:"peer_id"`
	PeerList []string `json:"peerlist"`
}

type PeerPeerDisconnectedEvent struct {
	PeerId   uint64 `json:"peer_id"`
	PeerAddr string `json:"peer_addr"`
}

type NewContractEvent struct {
	Contract   string `json:"contract"`
	BlockHash  string `json:"block_hash"`
	Topoheight uint64 `json:"topoheight"`
}

type InvokeContractEvent struct {
	BlockHash       string           `json:"block_hash"`
	TxHash          string           `json:"tx_hash"`
	Topoheight      uint64           `json:"topoheight"`
	ContractOutputs []ContractOutput `json:"contract_outputs"`
}

type InvokeContractEventParams struct {
	Contract string `json:"contract"`
}
