package wallet

import (
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
	Amount      uint64       `json:"amount"`
	Asset       string       `json:"asset"`
	Destination string       `json:"destination"`
	ExtraData   *interface{} `json:"extra_data,omitempty"`
}

type FeeBuilder struct {
	Multiplier *float64 `json:"multiplier,omitempty"`
	Value      *uint64  `json:"value,omitempty"`
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
	DeployContract *string                `json:"deploy_contract,omitempty"`
	Fee            *FeeBuilder            `json:"fee,omitempty"`
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

type TransactionEntry struct {
	Hash       string            `json:"hash"`
	Topoheight uint64            `json:"topoheight"`
	Outgoing   *Outgoing         `json:"outgoing"`
	Burn       *transaction.Burn `json:"burn"`
	Incoming   *Incoming         `json:"incoming"`
	Coinbase   *Coinbase         `json:"coinbase"`
}

type Transaction = transaction.Transaction

type TransactionResponse struct {
	transaction.Transaction
	TxAsHex *string `json:"tx_as_hex,omitempty"`
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
	Transfers      []TransferBuilder      `json:"transfers"`
	Burn           *transaction.Burn      `json:"burn,omitempty"`
	MultiSig       *MutliSigBuilder       `json:"multi_sig,omitempty"`
	InvokeContract *InvokeContractBuilder `json:"invoke_contract,omitempty"`
	DeployContract *string                `json:"deploy_contract,omitempty"`
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
	DeployContract *string                `json:"deploy_contract,omitempty"`
	Fee            *FeeBuilder            `json:"fee,omitempty"`
	TxVersion      *uint8                 `json:"tx_version,omitempty"`
	TxAsHex        bool                   `json:"tx_as_hex"`
	Reference      transaction.Reference  `json:"reference"`
	Nonce          uint64                 `json:"nonce"`
	Signers        *[]SignerId            `json:"signers,omitempty"`
}

type BuildUnsignedTransactionParams struct {
	Transfers      []TransferOut          `json:"transfers"`
	Burn           *transaction.Burn      `json:"burn,omitempty"`
	MultiSig       *MutliSigBuilder       `json:"multi_sig,omitempty"`
	InvokeContract *InvokeContractBuilder `json:"invoke_contract,omitempty"`
	DeployContract *string                `json:"deploy_contract,omitempty"`
	Nonce          *uint64                `json:"nonce,omitempty"`
	Fee            *FeeBuilder            `json:"fee,omitempty"`
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
	Nonce             uint64                         `json:"nonce"`
	SourceCommitments []transaction.SourceCommitment `json:"source_commitments"`
	Reference         transaction.Reference          `json:"reference"`
	RangeProof        RangeProof                     `json:"range_proof"`
	MultiSig          []string                       `json:"multisig"`
	Hash              string                         `json:"hash"`
	Threshold         uint8                          `json:"threshold"`
	TxAsHex           bool                           `json:"tx_as_hex"`
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
	Signatures []SignatureId `json:"signatures"`
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
	ConnectedTo       string `json:"connected_to"`
}

type CompressedCiphertext struct {
	Commitment []uint `json:"commitment"`
	Handle     []uint `json:"handle"`
}

type DecryptCiphertextParams struct {
	Ciphertext CompressedCiphertext `json:"ciphertext"`
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
	Tree  string `json:"tree"`
	Query *Query `json:"query,omitempty"`
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
	Tree          string `json:"tree"`
	Key           *Query `json:"key,omitempty"`
	Value         *Query `json:"value,omitempty"`
	ReturnOnFirst bool   `json:"return_on_first"`
}

type QueryResult struct {
	Entries map[string]interface{} `json:"entries"`
	Next    *uint64                `json:"next"`
}

type Asset struct {
	Decimals  int     `json:"decimals"`
	Name      string  `json:"name"`
	MaxSupply *uint64 `json:"max_supply"`
}

type GetAssetParams = GetAssetPrecisionParams
