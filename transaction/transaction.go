package transaction

import "github.com/xelis-project/xelis-go-sdk/xvm"

type Transfer struct {
	Asset           string  `json:"asset"`
	ExtraData       *[]uint `json:"extra_data"`
	Destination     string  `json:"destination"`
	Commitment      []uint  `json:"commitment"`
	SenderHandle    []uint  `json:"sender_handle"`
	ReceiverHandle  []uint  `json:"receiver_handle"`
	CTValidityProof Proof   `json:"ct_validity_proof"`
}

type Burn struct {
	Asset  string `json:"asset"`
	Amount uint64 `json:"amount"`
}

type CallContract struct {
	Contract string `json:"contract"`
}

type MultiSigPayload struct {
	Threshold    uint8    `json:"threshold"`
	Participants []string `json:"participants"`
}

type ValueCell struct {
	Type xvm.ValueCell `json:"type"`
}

type ContractDeposit struct {
	Public uint64 `json:"public"`
	// TODO: Private
}

type InvokeContractPayload struct {
	Contract   string                     `json:"contract"`
	Deposits   map[string]ContractDeposit `json:"deposits"`
	EntryId    uint16                     `json:"entry_id"`
	MaxGas     uint64                     `json:"max_gas"`
	Parameters []xvm.ValueCell            `json:"parameters"`
}

type InvokeConstructorPayload struct {
	MaxGas   uint64                     `json:"max_gas"`
	Deposits map[string]ContractDeposit `json:"deposits"`
}

type DeployContractPayload struct {
	Version string                    `json:"version"`
	Module  Module                    `json:"module"`
	Invoke  *InvokeConstructorPayload `json:"invoke,omitempty"`
}

type BlobPayload struct {
	Data         []uint   `json:"data"`
	Destinations []string `json:"destinations"`
}

type TransactionType struct {
	Transfers      *[]Transfer            `json:"transfers"`
	Burn           *Burn                  `json:"burn"`
	MultiSig       *MultiSigPayload       `json:"multi_sig"`
	InvokeContract *InvokeContractPayload `json:"invoke_contract"`
	DeployContract *DeployContractPayload `json:"deploy_contract"`
	Blob           *BlobPayload           `json:"blob"`
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

type FeeUsage struct {
	FeePaid   uint64 `json:"fee_paid"`
	FeeRefund uint64 `json:"fee_refund"`
}

type Transaction struct {
	Hash              string             `json:"hash"`
	Version           uint64             `json:"version"`
	Source            string             `json:"source"`
	Data              TransactionType    `json:"data"`
	Fee               uint64             `json:"fee"`
	FeeLimit          uint64             `json:"fee_limit"`
	FeePaid           *uint64            `json:"fee_paid,omitempty"`
	FeeRefund         *uint64            `json:"fee_refund,omitempty"`
	Nonce             uint64             `json:"nonce"`
	SourceCommitments []SourceCommitment `json:"source_commitments"`
	RangeProof        []uint             `json:"range_proof"`
	Reference         Reference          `json:"reference"`
	MultiSig          *MultiSig          `json:"multisig"`
	Signature         string             `json:"signature"`
	Size              uint64             `json:"size"`
}

type TransactionResponse struct {
	Hash              string             `json:"hash"`
	Version           uint64             `json:"version"`
	Source            string             `json:"source"`
	Data              TransactionType    `json:"data"`
	Fee               uint64             `json:"fee"`
	FeeLimit          uint64             `json:"fee_limit"`
	FeePaid           *uint64            `json:"fee_paid,omitempty"`
	FeeRefund         *uint64            `json:"fee_refund,omitempty"`
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

type ChunkAccessType string

const (
	ChunkAccessAll      ChunkAccessType = "all"
	ChunkAccessInternal ChunkAccessType = "internal"
	ChunkAccessEntry    ChunkAccessType = "entry"
	ChunkAccessHook     ChunkAccessType = "hook"
)

type ModuleChunk struct {
	// Instructions is hex string of bytecode instructions
	Instructions string          `json:"instructions"`
	Type         ChunkAccessType `json:"type"`
	Id           *uint8          `json:"id,omitempty"`
}

type Module struct {
	Constants []xvm.ValueCell `json:"constants"`
	Chunks    []ModuleChunk   `json:"chunks"`
}
