package xelis

import (
	"context"
	"testing"
)

func TestRPCGetInfo(t *testing.T) {
	daemon, err := NewDaemonRPC(DEV_NODE_RPC)
	if err != nil {
		t.Error(err)
		return
	}

	info, err := daemon.GetInfo(context.Background())
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v", info)
}
