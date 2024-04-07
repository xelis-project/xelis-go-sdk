package daemon

import (
	"context"
	"testing"

	"github.com/xelis-project/xelis-go-sdk/config"
)

var TESTING_ADDR = "xet:qf5u2p46jpgqmypqc2xwtq25yek2t7qhnqtdhw5kpfwcrlavs5asq0r83r7"

func setupRPC(t *testing.T) (daemon *RPC, ctx context.Context) {
	ctx = context.Background()
	daemon, err := NewRPC(ctx, config.TESTNET_NODE_RPC)
	if err != nil {
		t.Fatal(err)
	}

	return
}

func TestRPCMethods(t *testing.T) {
	daemon, _ := setupRPC(t)

	version, err := daemon.GetVersion()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", version)

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

	template, err := daemon.GetBlockTemplate(TESTING_ADDR)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", template)

	genesisBlock, err := daemon.GetBlockAtTopoheight(GetBlockAtTopoheightParams{Topoheight: 0})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", genesisBlock)

	blocks, err := daemon.GetBlocksAtHeight(GetBlocksAtHeightParams{Height: 0})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", blocks)

	block, err := daemon.GetBlockByHash(GetBlockByHashParams{Hash: `b715cb0229d13f5f540ae48adf03bc31b094b040b0756a2454631b2ddd899c3a`})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", block)

	topBlock, err := daemon.GetTopBlock(GetTopBlockParams{})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", topBlock)

	info, err := daemon.GetInfo()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", info)

	asset, err := daemon.GetAsset(config.XELIS_ASSET)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", asset)

	assets, err := daemon.GetAssets(GetAssetsParams{})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", assets)

	countAssets, err := daemon.CountAssets()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", countAssets)

	countAccounts, err := daemon.CountAccounts()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", countAccounts)

	countTransactions, err := daemon.CountTransactions()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", countTransactions)

	status, err := daemon.P2PStatus()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", status)

	result, err := daemon.GetPeers()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(result)

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

	dagOrder, err := daemon.GetDAGOrder(GetTopoheightRangeParams{})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", dagOrder)

	accounts, err := daemon.GetAccounts(GetAccountsParams{Maximum: 5})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", accounts)

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

func TestRPCNonceAndBalance(t *testing.T) {
	daemon, _ := setupRPC(t)
	has, err := daemon.HasNonce(TESTING_ADDR)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(has)

	has, err = daemon.HasBalance(GetBalanceParams{
		Address: TESTING_ADDR,
		Asset:   config.XELIS_ASSET,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(has)

	nonce, err := daemon.GetNonce(TESTING_ADDR)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(nonce)

	nonce, err = daemon.GetNonceAtTopoheight(GetNonceAtTopoheightParams{
		Address:    TESTING_ADDR,
		Topoheight: nonce.Topoheight,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(nonce)

	balance, err := daemon.GetBalance(GetBalanceParams{
		Address: TESTING_ADDR,
		Asset:   config.XELIS_ASSET,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(balance)

	versionedBalance, err := daemon.GetBalanceAtTopoheight(GetBalanceAtTopoheightParams{
		Address:    TESTING_ADDR,
		Asset:      config.XELIS_ASSET,
		Topoheight: nonce.Topoheight, // the testing addr does not have a balance before 322
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(versionedBalance)
}

func TestRPCGetBlocksRange(t *testing.T) {
	daemon, _ := setupRPC(t)

	topoheight, err := daemon.GetTopoheight()
	if err != nil {
		t.Fatal(err)
	}

	blocks, err := daemon.GetBlocksRangeByTopoheight(GetTopoheightRangeParams{
		StartTopoheight: topoheight - 10,
		EndTopoheight:   topoheight,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(blocks)

	height, err := daemon.GetHeight()
	if err != nil {
		t.Fatal(err)
	}

	blocks, err = daemon.GetBlocksRangeByHeight(GetHeightRangeParams{
		StartHeight: height - 10,
		EndHeight:   height,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(blocks)
}

func TestRPCGetTransactions(t *testing.T) {
	daemon, _ := setupRPC(t)
	txHash := "1de03df36b75916c2a440e428a854109c9628ed2c2bd628f9b9408baa78c6f52"

	txs, err := daemon.GetTransactions(GetTransactionsParams{
		TxHashes: []string{txHash},
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(txs)

	tx, err := daemon.GetTransaction(txHash)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(tx)
}

func TestRPCExecutedInBlock(t *testing.T) {
	// https://testnet-explorer.xelis.io/blocks/000000001849d07bbb4165c8ba1d1fc472a0629f56895efb8689e06ce62b3ca8
	daemon, _ := setupRPC(t)
	executed, err := daemon.IsTxExecutedInBlock(IsTxExecutedInBlockParams{
		TxHash:    "6e4bbd77b305fb68e2cc7576b4846d2db3617e3cbc2eb851cb2ae69b879e9d0f",
		BlockHash: "000000001849d07bbb4165c8ba1d1fc472a0629f56895efb8689e06ce62b3ca8",
	})
	if err != nil {
		t.Fatal(err)
	}

	if !executed {
		t.Errorf("tx should be executed in block")
	}

	t.Log(executed)
}

func TestRPCAccount(t *testing.T) {
	daemon, _ := setupRPC(t)
	history, err := daemon.GetAccountHistory(TESTING_ADDR)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(history)

	assets, err := daemon.GetAccountAssets(TESTING_ADDR)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(assets)
}
