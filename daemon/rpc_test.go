package daemon

import (
	"context"
	"testing"

	"github.com/xelis-project/xelis-go-sdk/config"
)

func setupRPC(t *testing.T) (daemon *RPC, ctx context.Context) {
	ctx = context.Background()
	daemon, err := NewRPC(config.TESTNET_NODE_RPC)
	if err != nil {
		t.Fatal(err)
	}

	return
}

func TestRPCGetInfo(t *testing.T) {
	daemon, ctx := setupRPC(t)

	version, err := daemon.GetVersion(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", version)

	info, err := daemon.GetInfo(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", info)

	topBlock, err := daemon.GetTopBlock(ctx, GetTopBlockParams{})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", topBlock)

	height, err := daemon.GetHeight(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", height)

	topoheight, err := daemon.GetTopoheight(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", topoheight)

	stableheight, err := daemon.GetStableheight(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", stableheight)

	countAssets, err := daemon.CountAssets(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", countAssets)

	countTransactions, err := daemon.CountTransactions(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", countTransactions)

	nonce, err := daemon.GetNonce(ctx, "xet1qqqyvh9vgkcurtj2la0e4jspnfsq7vkaqm863zcfdnej92xg4mpzz3suf96k4")
	if err != nil {
		//t.Fatal(err)
	}
	t.Logf("%+v", nonce)

	assets, err := daemon.GetAssets(ctx, GetAssetsParams{})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", assets)

	status, err := daemon.P2PStatus(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", status)

	mempool, err := daemon.GetMempool(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", mempool)

	tips, err := daemon.GetTips(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", tips)
}

func TestRPCUnknownMethod(t *testing.T) {
	daemon, ctx := setupRPC(t)
	res, err := daemon.Client.Call(ctx, "UnknownMethod", nil)
	if err == nil {
		t.Fatal("Expected an error")
	}

	t.Log(res)
}

func TestRPCHasNonce(t *testing.T) {
	daemon, ctx := setupRPC(t)
	res, err := daemon.HasNonce(ctx, "xet1qqq8ar5gagvjhznhj59l3r4lqhe7edutendy6vd4y7jd59exl6u7xschfuhym")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(res)
}

func TestRPCGetBlocksRange(t *testing.T) {
	daemon, ctx := setupRPC(t)

	topoheight, err := daemon.GetTopoheight(ctx)
	if err != nil {
		t.Fatal(err)
	}

	blocks, err := daemon.GetBlocksRangeByTopoheight(ctx, GetTopoHeightRangeParams{
		StartTopoheight: topoheight - 10,
		EndTopoheight:   topoheight,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(blocks)
}

func TestRPCGetTransactions(t *testing.T) {
	daemon, ctx := setupRPC(t)
	txs, err := daemon.GetTransactions(ctx, GetTransactionsParams{
		TxHashes: []string{},
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(txs)
}
