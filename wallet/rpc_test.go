package wallet

import (
	"context"
	"testing"

	"github.com/xelis-project/xelis-go-sdk/config"
)

// cargo run --bin xelis_wallet -- --wallet-path ./wallets/test --password test --network dev --rpc-password test --rpc-bind-address 127.0.0.1:8081 --rpc-username test
func setupRPC(t *testing.T) (wallet *RPC, ctx context.Context) {
	ctx = context.Background()
	wallet, err := NewRPC(ctx, config.LOCAL_WALLET_RPC, "test", "test")
	if err != nil {
		t.Fatal(err)
	}

	return
}

func TestRPCGetInfo(t *testing.T) {
	wallet, _ := setupRPC(t)

	version, err := wallet.GetVersion()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", version)

	network, err := wallet.GetNetwork()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", network)

	nonce, err := wallet.GetNonce()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", nonce)

	topo, err := wallet.GetTopoheight()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", topo)

	address, err := wallet.GetAddress(GetAddressParams{})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", address)

	integratedData := map[string]interface{}{"hello": "world"}
	integratedAddress, err := wallet.GetAddress(GetAddressParams{IntegratedData: &integratedData})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", integratedAddress)

	split, err := wallet.SplitAddress(SplitAddressParams{
		Address: integratedAddress,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", split)

	isOnline, err := wallet.IsOnline()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", isOnline)
}

func TestRPCRescan(t *testing.T) {
	wallet, _ := setupRPC(t)
	_, err := wallet.Rescan(RescanParams{})
	if err != nil {
		t.Fatal(err)
	}
}

func TestRPCSignData(t *testing.T) {
	wallet, _ := setupRPC(t)
	data, err := wallet.SignData(map[string]interface{}{"hello": "world"})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", data)
}

func TestRPCBalanceAndAsset(t *testing.T) {
	wallet, _ := setupRPC(t)

	balance, err := wallet.GetBalance(GetBalanceParams{
		Asset: config.XELIS_ASSET,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", balance)

	hasBalance, err := wallet.HasBalance(GetBalanceParams{
		Asset: config.XELIS_ASSET,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", hasBalance)

	precision, err := wallet.GetAssetPrecision(GetAssetPrecisionParams{
		Asset: config.XELIS_ASSET,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", precision)

	assets, err := wallet.GetTrackedAssets()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", assets)
}

func TestRPCGetTransaction(t *testing.T) {
	wallet, _ := setupRPC(t)

	// Send: 50ebdb059e5c9ad0f9fc7ac5d970b17ec2fc81bf197c9e737cf2d3ca14c5ae84
	// Burn: 8383b7027694615e790bea812a3385af8f140b55734b8eb89bf8a42d0671aec7
	// Receive: 37ecec82d39ea38d94240335d9fc1de01d039d52c764709f37766d36e3f5c336
	// Coingbase: 000000001fd2bf51d9c895bc200bd3e17597edd9827ac616d66884b75b55ddab
	txs, err := wallet.GetTransaction(GetTransactionParams{
		Hash: "50ebdb059e5c9ad0f9fc7ac5d970b17ec2fc81bf197c9e737cf2d3ca14c5ae84",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", txs)
}

func TestRPCListTransactions(t *testing.T) {
	wallet, _ := setupRPC(t)

	txs, err := wallet.ListTransactions(ListTransactionsParams{
		AcceptOutgoing: true,
		AcceptIncoming: true,
		AcceptCoinbase: true,
		AcceptBurn:     true,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", txs)
}
