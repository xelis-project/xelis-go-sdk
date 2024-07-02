package daemon

import (
	"context"
	"testing"

	"github.com/xelis-project/xelis-go-sdk/config"
)

const TESTING_ADDR = "xet:62wnkswt0rmrdd9d2lawgpzuh87fkpmp4gx9j3g4u24yrdkdxgksqnuuucf"
const MAINNET_ADDR = "xel:as3mgjlevw5ve6k70evzz8lwmsa5p0lgws2d60fulxylnmeqrp9qqukwdfg"

func useRPCTestnet(t *testing.T) (daemon *RPC, ctx context.Context) {
	ctx = context.Background()
	daemon, err := NewRPC(ctx, config.TESTNET_NODE_RPC)
	if err != nil {
		t.Fatal(err)
	}

	return
}

func useRPCMainnet(t *testing.T) (daemon *RPC, ctx context.Context) {
	ctx = context.Background()
	daemon, err := NewRPC(ctx, config.MAINNET_NODE_RPC)
	if err != nil {
		t.Fatal(err)
	}

	return
}

func TestRPCMethods(t *testing.T) {
	daemon, _ := useRPCTestnet(t)

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

	stableheight, err := daemon.GetStableHeight()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", stableheight)

	stableTopoheight, err := daemon.GetStableTopoheight()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", stableTopoheight)

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

	block, err := daemon.GetBlockByHash(GetBlockByHashParams{Hash: `23827b240a9e6aeb0e7164a4e402838ffc383efdc92789d705921fccfed516b5`})
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

	diff, err := daemon.GetDifficulty()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", diff)
}

func TestValidateAddress(t *testing.T) {
	daemon, _ := useRPCTestnet(t)

	validAddr, err := daemon.ValidateAddress(ValidateAddressParams{
		Address:         TESTING_ADDR,
		AllowIntegrated: false,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", validAddr)

	bytePublicKey, err := daemon.ExtractKeyFromAddress(ExtractKeyFromAddressParams{
		Address: TESTING_ADDR,
		AsHex:   false,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", bytePublicKey)

	hexPublicKey, err := daemon.ExtractKeyFromAddress(ExtractKeyFromAddressParams{
		Address: TESTING_ADDR,
		AsHex:   true,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", hexPublicKey)
}

func TestRPCUnknownMethod(t *testing.T) {
	daemon, ctx := useRPCTestnet(t)
	res, err := daemon.Client.Call(ctx, "UnknownMethod", nil)
	if err == nil {
		t.Fatal("Expected an error")
	}

	t.Log(res)
}

func TestRPCNonceAndBalance(t *testing.T) {
	daemon, _ := useRPCTestnet(t)
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

	versionedNonce, err := daemon.GetNonceAtTopoheight(GetNonceAtTopoheightParams{
		Address:    TESTING_ADDR,
		Topoheight: nonce.Topoheight,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(versionedNonce)

	balance, err := daemon.GetBalance(GetBalanceParams{
		Address: TESTING_ADDR,
		Asset:   config.XELIS_ASSET,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(balance)

	stableBalance, err := daemon.GetStableBalance(GetBalanceParams{
		Address: TESTING_ADDR,
		Asset:   config.XELIS_ASSET,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(stableBalance)

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
	daemon, _ := useRPCTestnet(t)

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
	daemon, _ := useRPCTestnet(t)
	txHash := "d9a6810d667c212e499ceb2acf60a8fbc0096da66b1e7a59fb3ae5d412ad58f2"

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

func TestRPCGetTransaction(t *testing.T) {
	daemon, _ := useRPCTestnet(t)
	txHash := "5f5e2ff1677860ee1f3e3c58ba188f427fbcb2f344dfb15dd0f7ca60b03f624c"

	tx, err := daemon.GetTransaction(txHash)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(tx)
}

func TestRPCExecutedInBlock(t *testing.T) {
	// https://testnet-explorer.xelis.io/blocks/000000001849d07bbb4165c8ba1d1fc472a0629f56895efb8689e06ce62b3ca8
	daemon, _ := useRPCTestnet(t)
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
	daemon, _ := useRPCTestnet(t)
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

	topoheight, err := daemon.GetAccountRegistrationTopoheight(TESTING_ADDR)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(topoheight)
}

func TestRPCRegistration(t *testing.T) {
	// using mainnet for this test
	// we need to resync the blockchain to work on testnet
	daemon, _ := useRPCMainnet(t)

	topoheight, err := daemon.GetAccountRegistrationTopoheight(MAINNET_ADDR)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(topoheight)

	exists, err := daemon.IsAccountRegistered(IsAccountRegisteredParams{
		Address:        MAINNET_ADDR,
		InStableHeight: true,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(exists)
}

func TestGetMinerWork(t *testing.T) {
	daemon, _ := useRPCTestnet(t)

	var addr = "xet:w64wu066sq7jq4v9f37a5gy8hgyvc2gvt237u2457mme2m2r7avqqtmufz3"

	blockTemplate, err := daemon.GetBlockTemplate(addr)
	if err != nil {
		t.Fatal(err)
	}

	result, err := daemon.GetMinerWork(GetMinerWorkParams{
		Template: blockTemplate.Template,
		Address:  &addr,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(result)
}
