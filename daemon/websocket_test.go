package daemon

import (
	"testing"
	"time"

	"github.com/xelis-project/xelis-go-sdk/config"
)

func TestWSGetInfo(t *testing.T) {
	daemon, err := NewWebSocket(config.DEV_NODE_WS)
	if err != nil {
		t.Error(err)
		return
	}

	info, err := daemon.GetInfo()
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v", info)
}

func TestWSNewBlock(t *testing.T) {
	daemon, err := NewWebSocket(config.DEV_NODE_WS)
	if err != nil {
		t.Error(err)
		return
	}

	done := make(chan struct{})
	err = daemon.OnNewBlock(func(newBlock *NewBlockResult, err error) {
		if err != nil {
			t.Error(err)
			return
		}

		t.Logf("%+v", newBlock)
		close(done)
	}, done)

	if err != nil {
		t.Error(err)
		return
	}

	select {
	case <-done:
	case <-time.After(30 * time.Second):
		t.Error("timeout")
	}
}
