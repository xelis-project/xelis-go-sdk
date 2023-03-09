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

func TestRPCInvalidMethod(t *testing.T) {
	daemon, err := NewRPC(config.DEV_NODE_RPC)
	if err != nil {
		t.Error(err)
		return
	}

	res, err := daemon.Client.Call(context.Background(), "InvalidMethod", nil)
	if err == nil {
		t.Error("Expected an error")
		return
	}

	t.Log(res)
}

func TestRPCGetBlocks(t *testing.T) {
	daemon, err := NewRPC(config.DEV_NODE_RPC)
	if err != nil {
		t.Error(err)
		return
	}

	blocks, err := daemon.GetBlocks(context.Background(), &GetRangeParams{
		StartTopoheight: 0,
		EndTopoheight:   5,
	})
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(blocks)
}

func TestRPCGetTransactions(t *testing.T) {
	daemon, err := NewRPC(config.DEV_NODE_RPC)
	if err != nil {
		t.Error(err)
		return
	}

	txs, err := daemon.GetTransactions(context.Background(), &GetTransactionsParams{
		TxHashes: []string{},
	})
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(txs)
}
