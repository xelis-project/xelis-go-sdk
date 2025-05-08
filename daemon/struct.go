package daemon

import "github.com/xelis-project/xelis-go-sdk/sc_constant"

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
	Commitment []uint `json:"commitment"`
	Handle     []uint `json:"handle"`
}

type VersionedBalance struct {
	BalanceType        BalanceType       `json:"balance_type"`
	FinalBalance       EncryptedBalance  `json:"final_balance"`
	OutputBalance      *EncryptedBalance `json:"output_balance"`
	PreviousTopoheight *uint64           `json:"previous_topoheight"`
}

type GetBalanceResult struct {
	Version    VersionedBalance `json:"version"`
	Topoheight uint64           `json:"topoheight"`
}

type GetStableBalanceResult struct {
	StableTopoheight uint64           `json:"stable_topoheight"`
	StableBlockHash  string           `json:"stable_block_hash"`
	Version          VersionedBalance `json:"version"`
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
	Skip              uint64 `json:"skip,omitempty"`
	Maximum           uint64 `json:"maximum,omitempty"`
	MinimumTopoheight uint64 `json:"minimum_topoheight,omitempty"`
	MaximumTopoheight uint64 `json:"maximum_topoheight,omitempty"`
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
	TotalSizeInBytes     uint64        `json:"total_size_in_bytes"`
	Version              uint64        `json:"version"`
	Tips                 []string      `json:"tips"`
	Timestamp            uint64        `json:"timestamp"`
	Height               uint64        `json:"height"`
	Nonce                uint64        `json:"nonce"`
	ExtraNonce           string        `json:"extra_nonce"`
	Miner                string        `json:"miner"`
	TxsHashes            []string      `json:"txs_hashes"`
	Transactions         []Transaction `json:"transactions"` // if include_txs is true in params
}

type Transfer struct {
	Asset           string      `json:"asset"`
	ExtraData       *[]uint     `json:"extra_data"`
	Destination     interface{} `json:"destination"` // can be []uint or string - string from daemon - []uint from wallet
	Commitment      []uint      `json:"commitment"`
	SenderHandle    []uint      `json:"sender_handle"`
	ReceiverHandle  []uint      `json:"receiver_handle"`
	CTValidityProof Proof       `json:"ct_validity_proof"`
}

type Burn struct {
	Asset  string `json:"asset"`
	Amount uint64 `json:"amount"`
}

type CallContract struct {
	Contract string `json:"contract"`
}

type MutliSigPayload struct {
	Threshold    uint8    `json:"threshold"`
	Participants []string `json:"participants"`
}

type ContractDeposit struct {
	Public uint64 `json:"public"`
	// TODO: Private
}

type InvokeContractPayload struct {
	Contract   string                     `json:"contract"`
	Deposits   map[string]ContractDeposit `json:"deposits"`
	ChunkId    uint16                     `json:"chunk_id"`
	MaxGas     uint64                     `json:"max_gas"`
	Parameters [][]uint                   `json:"parameters"`
}

type TransactionType struct {
	Transfers      *[]Transfer            `json:"transfers"`
	Burn           *Burn                  `json:"burn"`
	MultiSig       *MutliSigPayload       `json:"multi_sig"`
	InvokeContract *InvokeContractPayload `json:"invoke_contract"`
	DeployContract *Module                `json:"deploy_contract"`
}

type Reference struct {
	Hash       string `json:"hash"`
	Topoheight uint64 `json:"topoheight"`
}

type Proof struct {
	Y_0 []uint `json:"Y_0"`
	Y_1 []uint `json:"Y_1"`
	Z_R []uint `json:"z_r"`
	Z_X []uint `json:"z_x"`
}

type EqProof struct {
	Y_0 []uint `json:"Y_0"`
	Y_1 []uint `json:"Y_1"`
	Y_2 []uint `json:"Y_2"`
	Z_R []uint `json:"z_r"`
	Z_S []uint `json:"z_s"`
	Z_X []uint `json:"z_x"`
}

type SourceCommitment struct {
	Commitment []uint  `json:"commitment"`
	Proof      EqProof `json:"proof"`
	Asset      string  `json:"asset"`
}

type SignatureId struct {
	Id        uint8  `json:"id"`
	Signature string `json:"signature"`
}

type MultiSig struct {
	Signatures map[uint8]SignatureId `json:"signatures"`
}

type Transaction struct {
	Hash              string             `json:"hash"`
	Version           uint64             `json:"version"`
	Source            string             `json:"source"`
	Data              TransactionType    `json:"data"`
	Fee               uint64             `json:"fee"`
	Nonce             uint64             `json:"nonce"`
	SourceCommitments []SourceCommitment `json:"source_commitments"`
	RangeProof        []uint             `json:"range_proof"`
	Reference         Reference          `json:"reference"`
	MultiSig          *MultiSig          `json:"multisig"`
	Signature         string             `json:"signature"`
	Size              uint64             `json:"size"`
}

// copy of Transaction with more fields
type TransactionResponse struct {
	Hash              string             `json:"hash"`
	Version           uint64             `json:"version"`
	Source            string             `json:"source"`
	Data              TransactionType    `json:"data"`
	Fee               uint64             `json:"fee"`
	Nonce             uint64             `json:"nonce"`
	SourceCommitments []SourceCommitment `json:"source_commitments"`
	RangeProof        []uint             `json:"range_proof"`
	Reference         Reference          `json:"reference"`
	MultiSig          *MultiSig          `json:"multisig"`
	Signature         string             `json:"signature"`
	Size              uint64             `json:"size"`

	Blocks          []string `json:"blocks"`
	ExecutedInBlock *string  `json:"executed_in_block"`
	InMempool       bool     `json:"in_mempool"`
	FirstSeen       *uint64  `json:"first_seen"`
}

type GetInfoResult struct {
	Height            uint64 `json:"height"`
	Topoheight        uint64 `json:"topoheight"`
	Stableheight      uint64 `json:"stableheight"`
	PrunedTopoheight  uint64 `json:"pruned_topoheight"`
	TopBlockHash      string `json:"top_block_hash"`
	CirculatingSupply uint64 `json:"circulating_supply"`
	BurnedSupply      uint64 `json:"burned_supply"`
	EmittedSupply     uint64 `json:"emitted_supply"`
	MaximumSupply     uint64 `json:"maximum_supply"`
	Difficulty        string `json:"difficulty"`
	BlockTimeTarget   uint64 `json:"block_time_target"`
	AverageBlockTime  uint64 `json:"average_block_time"`
	BlockReward       uint64 `json:"block_reward"`
	DevReward         uint64 `json:"dev_reward"`
	MinerReward       uint64 `json:"miner_reward"`
	MempoolSize       uint64 `json:"mempool_size"`
	Version           string `json:"version"`
	Network           string `json:"network"`
}

type GetBlockTemplateResult struct {
	Template   string `json:"template"`
	Algorithm  string `json:"algorithm"`
	Height     uint64 `json:"height"`
	Topoheight uint64 `json:"topoheight"`
	Difficulty string `json:"difficulty"`
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
	Amount uint64 `json:"amount"`
}

type OutgoingHistory struct {
	To string `json:"to"`
}

type IncomingHistory struct {
	From string `json:"from"`
}

type AssetData struct {
	Asset      string  `json:"asset"`
	Topoheight uint64  `json:"topoheight"`
	Decimals   int     `json:"decimals"`
	Name       string  `json:"name"`
	MaxSupply  *uint64 `json:"max_supply"`
	Contract   *string `json:"contract"`
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
	DevFee         *MiningHistory   `json:"dev_fee"`
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

type SplitAddressParams struct {
	Address string `json:"address"`
}

type SplitAddressResult struct {
	Address        string      `json:"address"`
	IntegratedData interface{} `json:"integrated_data"`
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
	Topoheight uint32 `json:"topoheight"`
}

type GetMultisigAtTopoheightParams struct {
	Address    string `json:"address"`
	Topoheight uint32 `json:"topoheight"`
}

type GetMultisigAtTopoheightResult struct {
	State string `json:"state"`
}

type GetMultisigParams struct {
	Address string `json:"address"`
}

type GetMultisigResult struct {
	State      string `json:"state"`
	Topoheight uint64 `json:"topoheight"`
}

type HasMultisigParams struct {
	Address    string  `json:"address"`
	Topoheight *uint32 `json:"topoheight,omitempty"`
}

type GetContractOutputsParams struct {
	Transaction string `json:"transaction"`
}

type GetContractModuleParams struct {
	Contract string `json:"contract"`
}

type GetContractDataParams struct {
	Contract string               `json:"contract"`
	Key      sc_constant.Constant `json:"key"`
}

type GetContractDataAtTopoheightParams struct {
	Contract   string               `json:"contract"`
	Key        sc_constant.Constant `json:"key"`
	Topoheight uint64               `json:"topoheight"`
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

type Chunk struct {
	Instructions []uint `json:"instructions"`
}

type Module struct {
	Constants     []sc_constant.Constant `json:"constants"`
	Chunks        []Chunk                `json:"chunks"`
	EntryChunkIds []uint64               `json:"entry_chunk_ids"`
	Structs       []interface{}          `json:"structs"`
	Enums         []interface{}          `json:"enums"`
}

type GetContractModuleResult struct {
	PreviousTopoheight *uint64 `json:"previous_topoheight"`
	Data               *Module `json:"data"`
}

type GetContractDataResult struct {
	PreviousTopoheight *uint64      `json:"previous_topoheight"`
	Data               *interface{} `json:"data"`
}

type HardFork struct {
	Height             uint64  `json:"height"`
	Version            uint8   `json:"version"`
	Changelog          string  `json:"changelog"`
	VersionRequirement *string `json:"version_requirement"`
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
	SharedKey []uint `json:"shared_key"`
	ExtraData []uint `json:"extra_data"`
}

type GetMempoolCacheParams struct {
	Address string `json:"address"`
}

type GetMempoolCacheResult struct {
	Min      uint64                 `json:"min"`
	Max      uint64                 `json:"max"`
	Txs      []string               `json:"txs"`
	Balances map[string]interface{} `json:"balances"`
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
