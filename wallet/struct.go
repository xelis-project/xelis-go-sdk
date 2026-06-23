package wallet

import (
	"encoding/json"

	"github.com/xelis-project/xelis-go-sdk/data"
	"github.com/xelis-project/xelis-go-sdk/transaction"
	"github.com/xelis-project/xelis-go-sdk/xvm"
)

// use []uint instead of []byte or []uint8 to avoid json.Marshal base64 encoding on params

type GetAddressParams struct {
	IntegratedData *interface{} `json:"integrated_data,omitempty"`
}

type SplitAddressParams struct {
	Address string `json:"address"`
}

type SplitAddressResult struct {
	Address        string      `json:"address"`
	IntegratedData interface{} `json:"integrated_data"`
	Size           uint64      `json:"size"`
}

type GetBalanceParams struct {
	Asset string `json:"asset,omitempty"`
}

type GetTransactionParams struct {
	Hash string `json:"hash"`
}

type RescanParams struct {
	UntilTopoheight *uint64 `json:"until_topoheight,omitempty"`
	AutoReconnect   bool    `json:"auto_reconnect"`
}

type GetAssetPrecisionParams struct {
	Asset string `json:"asset"`
}

type GetAssetsParams struct {
	Skip    *uint64 `json:"skip,omitempty"`
	Maximum *uint64 `json:"maximum,omitempty"`
}

type GetAssetsEntry struct {
	Asset string `json:"asset"`
	Data  Asset  `json:"data"`
}

type TrackAssetParams = GetAssetPrecisionParams
type IsAssetTrackedParams = GetAssetPrecisionParams

type TransferIn struct {
	Amount    uint64              `json:"amount"`
	Asset     string              `json:"asset"`
	ExtraData *PlaintextExtraData `json:"extra_data"`
}

type PlaintextExtraData struct {
	SharedKey string      `json:"shared_key"`
	Data      interface{} `json:"data"`
}

type TransferOut struct {
	Amount      uint64              `json:"amount"`
	Asset       string              `json:"asset"`
	Destination string              `json:"destination"`
	ExtraData   *PlaintextExtraData `json:"extra_data,omitempty"`
}

type TransferBuilder struct {
	Amount           uint64       `json:"amount"`
	Asset            string       `json:"asset"`
	Destination      string       `json:"destination"`
	ExtraData        *interface{} `json:"extra_data,omitempty"`
	EncryptExtraData *bool        `json:"encrypt_extra_data,omitempty"`
}

type ExtraFeeMode struct {
	Tip        *uint64  `json:"tip,omitempty"`
	Multiplier *float64 `json:"multiplier,omitempty"`
}

func (e ExtraFeeMode) MarshalJSON() ([]byte, error) {
	switch {
	case e.Tip != nil:
		return json.Marshal(map[string]uint64{"tip": *e.Tip})
	case e.Multiplier != nil:
		return json.Marshal(map[string]float64{"multiplier": *e.Multiplier})
	default:
		return json.Marshal("none")
	}
}

type FeeBuilder struct {
	Fixed *uint64       `json:"-"`
	Extra *ExtraFeeMode `json:"-"`

	// Deprecated: use Extra.Multiplier. Kept as a compatibility bridge for older SDK callers.
	Multiplier *float64 `json:"multiplier,omitempty"`
	// Deprecated: use Fixed. Kept as a compatibility bridge for older SDK callers.
	Value *uint64 `json:"value,omitempty"`
}

func (f FeeBuilder) MarshalJSON() ([]byte, error) {
	switch {
	case f.Fixed != nil:
		return json.Marshal(map[string]uint64{"fixed": *f.Fixed})
	case f.Value != nil:
		return json.Marshal(map[string]uint64{"fixed": *f.Value})
	case f.Extra != nil:
		return json.Marshal(map[string]ExtraFeeMode{"extra": *f.Extra})
	case f.Multiplier != nil:
		return json.Marshal(map[string]map[string]float64{
			"extra": {"multiplier": *f.Multiplier},
		})
	default:
		return json.Marshal(map[string]string{"extra": "none"})
	}
}

type BaseFeeMode struct {
	Fixed *uint64
	Cap   *uint64
}

func (b BaseFeeMode) MarshalJSON() ([]byte, error) {
	switch {
	case b.Fixed != nil:
		return json.Marshal(map[string]uint64{"fixed": *b.Fixed})
	case b.Cap != nil:
		return json.Marshal(map[string]uint64{"cap": *b.Cap})
	default:
		return json.Marshal("none")
	}
}

type MutliSigBuilder struct {
	Participants []string `json:"participants"`
	Threshold    uint8    `json:"threshold"`
}

type ContractDepositBuilder struct {
	Amount  uint64 `json:"amount"`
	Private bool   `json:"private"`
}

type InvokeContractBuilder struct {
	Contract   string                            `json:"contract"`
	MaxGas     uint64                            `json:"max_gas"`
	EntryId    uint16                            `json:"entry_id"`
	Parameters []xvm.ValueCell                   `json:"parameters"`
	Deposits   map[string]ContractDepositBuilder `json:"deposits"`
	Permission interface{}                       `json:"permission,omitempty"`
}

type DeployContractInvokeBuilder struct {
	MaxGas   uint64                            `json:"max_gas"`
	Deposits map[string]ContractDepositBuilder `json:"deposits,omitempty"`
}

type DeployContractBuilder struct {
	ContractVersion interface{}                  `json:"contract_version,omitempty"`
	Module          string                       `json:"module"`
	Invoke          *DeployContractInvokeBuilder `json:"invoke,omitempty"`
}

type BlobPayloadBuilder struct {
	Data         interface{} `json:"data"`
	Encrypt      *bool       `json:"encrypt,omitempty"`
	Destinations []string    `json:"destinations"`
}

type SignerId struct {
	Id         uint8  `json:"id"`
	PrivateKey []uint `json:"private_key"`
}

type BuildTransactionParams struct {
	Transfers      []TransferBuilder      `json:"transfers,omitempty"`
	Burn           *transaction.Burn      `json:"burn,omitempty"`
	MultiSig       *MutliSigBuilder       `json:"multi_sig,omitempty"`
	InvokeContract *InvokeContractBuilder `json:"invoke_contract,omitempty"`
	DeployContract interface{}            `json:"deploy_contract,omitempty"`
	Blob           *BlobPayloadBuilder    `json:"blob,omitempty"`
	Fee            *FeeBuilder            `json:"fee,omitempty"`
	BaseFee        *BaseFeeMode           `json:"base_fee,omitempty"`
	FeeLimit       *uint64                `json:"fee_limit,omitempty"`
	Nonce          *uint64                `json:"nonce,omitempty"`
	TxVersion      *uint8                 `json:"tx_version,omitempty"`
	Broadcast      bool                   `json:"broadcast"`
	TxAsHex        bool                   `json:"tx_as_hex"`
	Signers        *[]SignerId            `json:"signers,omitempty"`
}

// !!! not the same as transaction.Transfer
// the destination is []uint and the other can be string
type Transfer struct {
	Asset           string            `json:"asset"`
	ExtraData       *[]uint           `json:"extra_data"`
	Destination     []uint            `json:"destination"`
	Commitment      []uint            `json:"commitment"`
	SenderHandle    []uint            `json:"sender_handle"`
	ReceiverHandle  []uint            `json:"receiver_handle"`
	CTValidityProof transaction.Proof `json:"ct_validity_proof"`
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

type BurnEntry struct {
	Asset  string `json:"asset"`
	Amount uint64 `json:"amount"`
	Fee    uint64 `json:"fee"`
	Nonce  uint64 `json:"nonce"`
}

type MultiSigEntry struct {
	Participants []string `json:"participants"`
	Threshold    uint8    `json:"threshold"`
	Fee          uint64   `json:"fee"`
	Nonce        uint64   `json:"nonce"`
}

type DeployInvoke struct {
	MaxGas   uint64            `json:"max_gas"`
	Deposits map[string]uint64 `json:"deposits"`
}

type InvokeContractEntry struct {
	Contract string            `json:"contract"`
	Deposits map[string]uint64 `json:"deposits"`
	Received map[string]uint64 `json:"received"`
	ChunkId  uint16            `json:"chunk_id"`
	Fee      uint64            `json:"fee"`
	MaxGas   uint64            `json:"max_gas"`
	Nonce    uint64            `json:"nonce"`
}

type DeployContractEntry struct {
	Fee    uint64        `json:"fee"`
	Nonce  uint64        `json:"nonce"`
	Invoke *DeployInvoke `json:"invoke"`
}

type IncomingContract struct {
	Transfers map[string]uint64 `json:"transfers"`
}

type OutgoingBlob struct {
	Destinations []string            `json:"destinations"`
	Fee          uint64              `json:"fee"`
	Nonce        uint64              `json:"nonce"`
	Data         *PlaintextExtraData `json:"data"`
}

type IncomingBlob struct {
	From         string              `json:"from"`
	Destinations []string            `json:"destinations"`
	Data         *PlaintextExtraData `json:"data"`
}

type TransactionEntry struct {
	Hash             string               `json:"hash"`
	Topoheight       uint64               `json:"topoheight"`
	Timestamp        uint64               `json:"timestamp"`
	Outgoing         *Outgoing            `json:"outgoing"`
	Burn             *BurnEntry           `json:"burn"`
	Incoming         *Incoming            `json:"incoming"`
	Coinbase         *Coinbase            `json:"coinbase"`
	MultiSig         *MultiSigEntry       `json:"multi_sig"`
	InvokeContract   *InvokeContractEntry `json:"invoke_contract"`
	DeployContract   *DeployContractEntry `json:"deploy_contract"`
	IncomingContract *IncomingContract    `json:"incoming_contract"`
	OutgoingBlob     *OutgoingBlob        `json:"outgoing_blob"`
	IncomingBlob     *IncomingBlob        `json:"incoming_blob"`
}

type SearchTransactionParams = GetTransactionParams

type SearchTransactionResult struct {
	Transaction *TransactionEntry `json:"transaction"`
	Index       *uint64           `json:"index"`
	IsRawSearch bool              `json:"is_raw_search"`
}

type TransactionPending struct {
	Hash             string               `json:"hash"`
	Timestamp        uint64               `json:"timestamp"`
	Outgoing         *Outgoing            `json:"outgoing"`
	Burn             *BurnEntry           `json:"burn"`
	Incoming         *Incoming            `json:"incoming"`
	Coinbase         *Coinbase            `json:"coinbase"`
	MultiSig         *MultiSigEntry       `json:"multi_sig"`
	InvokeContract   *InvokeContractEntry `json:"invoke_contract"`
	DeployContract   *DeployContractEntry `json:"deploy_contract"`
	IncomingContract *IncomingContract    `json:"incoming_contract"`
	OutgoingBlob     *OutgoingBlob        `json:"outgoing_blob"`
	IncomingBlob     *IncomingBlob        `json:"incoming_blob"`
}

type Transaction = transaction.Transaction

type TransactionResponse struct {
	transaction.Transaction
	TxAsHex *string `json:"tx_as_hex,omitempty"`
}

type ListTransactionsParams struct {
	Asset          *string `json:"asset,omitempty"`
	MinTopoheight  *uint64 `json:"min_topoheight,omitempty"`
	MaxTopoheight  *uint64 `json:"max_topoheight,omitempty"`
	MinTimestamp   *uint64 `json:"min_timestamp,omitempty"`
	MaxTimestamp   *uint64 `json:"max_timestamp,omitempty"`
	Address        *string `json:"address,omitempty"`
	AcceptIncoming bool    `json:"accept_incoming"`
	AcceptOutgoing bool    `json:"accept_outgoing"`
	AcceptCoinbase bool    `json:"accept_coinbase"`
	AcceptBurn     bool    `json:"accept_burn"`
	AcceptBlob     bool    `json:"accept_blob"`
	Query          *Query  `json:"query,omitempty"`
	Limit          *uint64 `json:"limit,omitempty"`
	Skip           *uint64 `json:"skip,omitempty"`
}

type EstimateFeesParams struct {
	Transfers      []TransferBuilder      `json:"transfers"`
	Burn           *transaction.Burn      `json:"burn,omitempty"`
	MultiSig       *MutliSigBuilder       `json:"multi_sig,omitempty"`
	InvokeContract *InvokeContractBuilder `json:"invoke_contract,omitempty"`
	DeployContract interface{}            `json:"deploy_contract,omitempty"`
	Blob           *BlobPayloadBuilder    `json:"blob,omitempty"`
	Fee            *FeeBuilder            `json:"fee,omitempty"`
	BaseFee        *BaseFeeMode           `json:"base_fee,omitempty"`
}

type BalanceChangedResult struct {
	Asset   string `json:"asset"`
	Balance uint64 `json:"balance"`
}

type BuildTransactionOfflineParams struct {
	Transfers      []TransferOut          `json:"transfers"`
	Burn           *transaction.Burn      `json:"burn,omitempty"`
	MultiSig       *MutliSigBuilder       `json:"multi_sig,omitempty"`
	InvokeContract *InvokeContractBuilder `json:"invoke_contract,omitempty"`
	DeployContract interface{}            `json:"deploy_contract,omitempty"`
	Blob           *BlobPayloadBuilder    `json:"blob,omitempty"`
	Fee            *FeeBuilder            `json:"fee,omitempty"`
	BaseFee        *uint64                `json:"base_fee,omitempty"`
	FeeLimit       *uint64                `json:"fee_limit,omitempty"`
	TxVersion      *uint8                 `json:"tx_version,omitempty"`
	TxAsHex        bool                   `json:"tx_as_hex"`
	Balances       map[string]interface{} `json:"balances,omitempty"`
	Reference      transaction.Reference  `json:"reference"`
	Nonce          uint64                 `json:"nonce"`
	Signers        *[]SignerId            `json:"signers,omitempty"`
}

type BuildUnsignedTransactionParams struct {
	Transfers      []TransferOut          `json:"transfers"`
	Burn           *transaction.Burn      `json:"burn,omitempty"`
	MultiSig       *MutliSigBuilder       `json:"multi_sig,omitempty"`
	InvokeContract *InvokeContractBuilder `json:"invoke_contract,omitempty"`
	DeployContract interface{}            `json:"deploy_contract,omitempty"`
	Blob           *BlobPayloadBuilder    `json:"blob,omitempty"`
	Nonce          *uint64                `json:"nonce,omitempty"`
	Fee            *FeeBuilder            `json:"fee,omitempty"`
	BaseFee        *BaseFeeMode           `json:"base_fee,omitempty"`
	FeeLimit       *uint64                `json:"fee_limit,omitempty"`
	TxVersion      *uint8                 `json:"tx_version,omitempty"`
	TxAsHex        bool                   `json:"tx_as_hex"`
}

type InnerProductProof struct {
}

type RangeProof struct {
	A            []uint            `json:"A"`
	S            []uint            `json:"S"`
	T_1          []uint            `json:"T_1"`
	T_2          []uint            `json:"T_2"`
	T_X          []uint            `json:"t_x"`
	T_X_Blinding []uint            `json:"t_x_blinding"`
	E_Blinding   []uint            `json:"e_blinding"`
	IppProof     InnerProductProof `json:"ipp_proof"`
}

type UnsignedTransactionResponse struct {
	Version           uint8                          `json:"version"`
	Source            []uint                         `json:"source"`
	Data              interface{}                    `json:"data"`
	Fee               uint64                         `json:"fee"`
	FeeLimit          uint64                         `json:"fee_limit"`
	Nonce             uint64                         `json:"nonce"`
	SourceCommitments []transaction.SourceCommitment `json:"source_commitments"`
	Reference         transaction.Reference          `json:"reference"`
	RangeProof        RangeProof                     `json:"range_proof"`
	MultiSig          []string                       `json:"multisig"`
	Hash              string                         `json:"hash"`
	Threshold         *uint8                         `json:"threshold"`
	TxAsHex           *string                        `json:"tx_as_hex"`
}

type SignUnsignedTransactionParams struct {
	Hash     string `json:"hash"`
	SignerId uint8  `json:"signer_id"`
}

type SignatureId struct {
	Id        uint8  `json:"id"`
	Signature string `json:"signature"`
}

type FinalizeUnsignedTransactionParams struct {
	Unsigned   string        `json:"unsigned"`
	Signatures []SignatureId `json:"signatures,omitempty"`
	Broadcast  bool          `json:"broadcast"`
	TxAsHex    bool          `json:"tx_as_hex"`
}

type EstimateExtraDataSizeParams struct {
	Destinations []string `json:"destinations"`
}

type EstimateExtraDataSizeResult struct {
	Size uint64 `json:"size"`
}

type NetworkInfoResult struct {
	Height            uint64  `json:"height"`
	Topoheight        uint64  `json:"topoheight"`
	Stableheight      uint64  `json:"stableheight"`
	StableTopoheight  uint64  `json:"stable_topoheight"`
	PrunedTopoheight  *uint64 `json:"pruned_topoheight"`
	TopBlockHash      string  `json:"top_block_hash"`
	CirculatingSupply uint64  `json:"circulating_supply"`
	BurnedSupply      uint64  `json:"burned_supply"`
	EmittedSupply     uint64  `json:"emitted_supply"`
	MaximumSupply     uint64  `json:"maximum_supply"`
	Difficulty        string  `json:"difficulty"`
	BlockTimeTarget   uint64  `json:"block_time_target"`
	AverageBlockTime  uint64  `json:"average_block_time"`
	BlockReward       uint64  `json:"block_reward"`
	DevReward         uint64  `json:"dev_reward"`
	MinerReward       uint64  `json:"miner_reward"`
	MempoolSize       uint64  `json:"mempool_size"`
	Version           string  `json:"version"`
	Network           string  `json:"network"`
	BlockVersion      uint8   `json:"block_version"`
	ConnectedTo       string  `json:"connected_to"`
}

type CompressedCiphertext struct {
	Commitment []uint `json:"commitment"`
	Handle     []uint `json:"handle"`
}

type DecryptCiphertextParams struct {
	Ciphertext CompressedCiphertext `json:"ciphertext"`
	MaxSupply  *uint64              `json:"max_supply,omitempty"`
}

type TxRole string

const (
	TxSenderRole   TxRole = "sender"
	TxReceiverRole TxRole = "receiver"
)

type DecryptExtraDataParams struct {
	ExtraData []uint `json:"extra_data"`
	Role      TxRole `json:"role"`
}

type GetMatchingKeysParams struct {
	Tree  string  `json:"tree"`
	Query *Query  `json:"query,omitempty"`
	Limit *uint64 `json:"limit,omitempty"`
	Skip  *uint64 `json:"skip,omitempty"`
}

type CountMatchingEntriesParams struct {
	Tree  string `json:"tree"`
	Key   *Query `json:"key,omitempty"`
	Value *Query `json:"value,omitempty"`
}

type GetValueFromKeyParams struct {
	Tree string      `json:"tree"`
	Key  interface{} `json:"key"`
}

type StoreParams struct {
	Tree  string      `json:"tree"`
	Key   interface{} `json:"key"`
	Value interface{} `json:"value"`
}

type DeleteParams struct {
	Tree string      `json:"tree"`
	Key  interface{} `json:"key"`
}

type DeleteTreeEntriesParams struct {
	Tree string `json:"tree"`
}

type HasKeyParams struct {
	Tree string      `json:"tree"`
	Key  interface{} `json:"key"`
}
type QueryDBParams struct {
	Tree  string  `json:"tree"`
	Key   *Query  `json:"key,omitempty"`
	Value *Query  `json:"value,omitempty"`
	Limit *uint64 `json:"limit,omitempty"`
	Skip  *uint64 `json:"skip,omitempty"`
}

type QueryResult struct {
	Entries map[string]interface{} `json:"entries"`
	Next    *uint64                `json:"next"`
}

type Asset struct {
	Decimals  int         `json:"decimals"`
	Name      string      `json:"name"`
	Ticker    string      `json:"ticker"`
	MaxSupply interface{} `json:"max_supply"`
	Owner     interface{} `json:"owner"`
}

type GetAssetParams = GetAssetPrecisionParams

type SetOnlineModeParams struct {
	DaemonAddress string `json:"daemon_address"`
	AutoReconnect bool   `json:"auto_reconnect"`
}

type VerifySignedDataParams struct {
	Data      data.Element `json:"data"`
	Signature string       `json:"signature"`
	Address   string       `json:"address"`
}

type CreateOwnershipProofParams struct {
	Asset      string  `json:"asset"`
	Topoheight *uint64 `json:"topoheight,omitempty"`
	Amount     uint64  `json:"amount"`
}

type CreateBalanceProofParams struct {
	Asset      string  `json:"asset"`
	Topoheight *uint64 `json:"topoheight,omitempty"`
}

type VerifyHumanReadableProofParams struct {
	Proof   interface{} `json:"proof"`
	Address string      `json:"address"`
}
