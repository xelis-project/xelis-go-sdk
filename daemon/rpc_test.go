package daemon

import (
	"context"
	"testing"

	"github.com/xelis-project/xelis-go-sdk/config"
)

func TestRPCGetInfo(t *testing.T) {
	daemon, err := NewRPC(config.DEV_NODE_RPC)
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
