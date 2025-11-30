package getwork

import "github.com/xelis-project/xelis-go-sdk/daemon"

const (
	NewJob        string = `new_job`
	BlockAccepted string = `block_accepted`
	BlockRejected string = `block_rejected`

	SubmitBlock string = `submit`
)

type BlockTemplate struct {
	Algorithm  daemon.AlgorithmVersion `json:"algorithm"`
	Difficulty string                  `json:"difficulty"`
	Height     uint64                  `json:"height"`
	TopoHeight uint64                  `json:"topoheight"`
	Template   string                  `json:"template"`
}

type MinerWork struct {
	Algorithm  daemon.AlgorithmVersion `json:"algorithm"`
	MinerWork  string                  `json:"miner_work"`
	Height     uint64                  `json:"height"`
	Difficulty string                  `json:"difficulty"`
	TopoHeight uint64                  `json:"topoheight"`
}
