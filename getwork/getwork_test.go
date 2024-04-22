package getwork

import (
	"log"
	"testing"

	"github.com/xelis-project/xelis-go-sdk/config"
)

const TESTNET_WALLET = "xet:6eadzwf5xdacts6fs4y3csmnsmy4mcxewqt3xyygwfx0hm0tm32sqxdy9zk"
const MAINNET_WALLET = "xel:vs3mfyywt0fjys0rgslue7mm4wr23xdgejsjk0ld7f2kxng4d4nqqnkdufz"

func TestGetworkAccepted(t *testing.T) {
	// using local daemon --network dev to test accepted block
	getwork, err := NewGetwork(config.LOCAL_NODE_GETWORK, TESTNET_WALLET, "xelis-go-sdk")
	if err != nil {
		t.Fatal(err)
	}

	for {
		select {
		case job := <-getwork.Job:
			t.Logf("%+v", job)

			err := getwork.SubmitBlock(job.Template)
			if err != nil {
				t.Fatal(err)
				return
			}
		case accepted := <-getwork.AcceptedBlock:
			t.Logf("%+v", accepted)
			return
		case err := <-getwork.Err:
			log.Fatal(err)
			return
		}
	}
}

func TestGetworkRejected(t *testing.T) {
	getwork, err := NewGetwork(config.TESTNET_NODE_GETWORK, TESTNET_WALLET, "xelis-go-sdk")
	if err != nil {
		t.Fatal(err)
	}

	for {
		select {
		case job := <-getwork.Job:
			t.Logf("%+v", job)

			err := getwork.SubmitBlock(job.Template)
			if err != nil {
				t.Fatal(err)
				return
			}
		case reason := <-getwork.RejectedBlock:
			t.Logf("%+v", reason)
			return
		case err := <-getwork.Err:
			log.Fatal(err)
			return
		}
	}
}
