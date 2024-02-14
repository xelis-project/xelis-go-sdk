package daemon

import (
	"context"
	"testing"

	"github.com/xelis-project/xelis-go-sdk/config"
)

func setupRPC(t *testing.T) (daemon *RPC, ctx context.Context) {
	ctx = context.Background()
	daemon, err := NewRPC(ctx, config.TESTNET_NODE_RPC)
	if err != nil {
		t.Fatal(err)
	}

	return
}

func TestRPCGetInfo(t *testing.T) {
	daemon, _ := setupRPC(t)

	version, err := daemon.GetVersion()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", version)

	info, err := daemon.GetInfo()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", info)

	topBlock, err := daemon.GetTopBlock(GetTopBlockParams{})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", topBlock)

	height, err := daemon.GetHeight()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", height)

	topoheight, err := daemon.GetTopoheight()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", topoheight)

	stableheight, err := daemon.GetStableheight()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", stableheight)

	countAssets, err := daemon.CountAssets()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", countAssets)

	countTransactions, err := daemon.CountTransactions()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", countTransactions)

	nonce, err := daemon.GetNonce("xet1qqqyvh9vgkcurtj2la0e4jspnfsq7vkaqm863zcfdnej92xg4mpzz3suf96k4")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", nonce)

	/*
		assets, err := daemon.GetAssets(GetAssetsParams{})
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%+v", assets)
	*/

	status, err := daemon.P2PStatus()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", status)

	mempool, err := daemon.GetMempool()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", mempool)

	tips, err := daemon.GetTips()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", tips)

	fees, err := daemon.GetDevFeeThresholds()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", fees)

	size, err := daemon.GetSizeOnDisk()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", size)
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
	daemon, _ := setupRPC(t)
	res, err := daemon.HasNonce("xet1qqq8ar5gagvjhznhj59l3r4lqhe7edutendy6vd4y7jd59exl6u7xschfuhym")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(res)

	res, err = daemon.HasBalance(GetBalanceParams{
		Address: "xet1qqq8ar5gagvjhznhj59l3r4lqhe7edutendy6vd4y7jd59exl6u7xschfuhym",
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(res)
}

func TestRPCGetBlocksRange(t *testing.T) {
	daemon, _ := setupRPC(t)

	topoheight, err := daemon.GetTopoheight()
	if err != nil {
		t.Fatal(err)
	}

	blocks, err := daemon.GetBlocksRangeByTopoheight(GetTopoHeightRangeParams{
		StartTopoheight: topoheight - 10,
		EndTopoheight:   topoheight,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(blocks)
}

func TestRPCGetTransactions(t *testing.T) {
	daemon, _ := setupRPC(t)
	txs, err := daemon.GetTransactions(GetTransactionsParams{
		TxHashes: []string{},
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(txs)
}

func TestRPCExecutedInBlock(t *testing.T) {
	// https://testnet-explorer.xelis.io/blocks/109
	daemon, _ := setupRPC(t)
	executed, err := daemon.IsTxExecutedInBlock(IsTxExecutedInBlockParams{
		TxHash:    "d4992c0c439ebdba9d8f0086cdefc21e95adcedae04985f21d309c208108765d",
		BlockHash: "000000a95b4b4ea13ba99d58c7a9bf68a502a2923049a54f8ac9e3826496cd9b",
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(executed)
}
