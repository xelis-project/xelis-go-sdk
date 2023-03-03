package xelis

import (
	"testing"
	"time"
)

func TestWSGetInfo(t *testing.T) {
	daemon, err := NewDaemonWS(DEV_NODE_WS)
	if err != nil {
		t.Error(err)
		return
	}

	info, err := daemon.CallGetInfo()
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v", info)
}

func TestWSNewBlock(t *testing.T) {
	daemon, err := NewDaemonWS(DEV_NODE_WS)
	if err != nil {
		t.Error(err)
		return
	}

	done := make(chan struct{})
	err = daemon.NewBlock(func(newBlock NewBlockResult, err error) {
		t.Logf("%+v", newBlock)
		close(done)
	}, done)

	if err != nil {
		t.Log(err)
		return
	}

	select {
	case <-done:
	case <-time.After(30 * time.Second):
		t.Error("timeout")
	}
}
