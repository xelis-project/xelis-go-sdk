package getwork

const (
	NewJob        string = `new_job`
	BlockAccepted string = `block_accepted`
	BlockRejected string = `block_rejected`

	SubmitBlock string = `submit`
)

type BlockTemplate struct {
	Difficulty string `json:"difficulty"`
	Height     uint64 `json:"height"`
	Template   string `json:"template"`
}
