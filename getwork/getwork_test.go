package getwork

import (
	"testing"

	"github.com/xelis-project/xelis-go-sdk/config"
)

func setupGetwork(t *testing.T) (getwork *Getwork) {
	getwork, err := NewGetwork(config.LOCAL_NODE_GETWORK, config.MAINNET_WALLET, "xelis-go-sdk")
	if err != nil {
		t.Fatal(err)
	}

	return
}

func TestGetwork(t *testing.T) {
	getwork := setupGetwork(t)

	job := <-getwork.Jobs

	t.Log("job:", job)

	err := getwork.SubmitBlock(job.Template)
	if err != nil {
		t.Fatal(err)
	}

	rejbl := <-getwork.RejectedBlocks

	t.Log("rejected block", rejbl)
}
