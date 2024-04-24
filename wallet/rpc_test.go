package wallet

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/xelis-project/xelis-go-sdk/config"
	"github.com/xelis-project/xelis-go-sdk/daemon"
)

const TESTING_ADDR = "xet:qf5u2p46jpgqmypqc2xwtq25yek2t7qhnqtdhw5kpfwcrlavs5asq0r83r7"
const MAINNET_ADDR = "xel:as3mgjlevw5ve6k70evzz8lwmsa5p0lgws2d60fulxylnmeqrp9qqukwdfg"

// cargo run --bin xelis_wallet -- --wallet-path ./wallets/test --password test --network dev --rpc-password test --rpc-bind-address 127.0.0.1:8081 --rpc-username test

func useRPCLocal(t *testing.T) (wallet *RPC, ctx context.Context) {
	ctx = context.Background()
	wallet, err := NewRPC(ctx, config.LOCAL_WALLET_RPC, "test", "test")
	if err != nil {
		t.Fatal(err)
	}

	return
}

func TestRPCGetInfo(t *testing.T) {
	wallet, _ := useRPCLocal(t)

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

	isOnline, err := wallet.IsOnline()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", isOnline)
}

func TestRPCIntegratedAddress(t *testing.T) {
	wallet := useWSLocal(t)

	integratedData := DataElement{
		Value:  "hello world",
		Array:  []DataElement{DataElement{Value: "test"}},
		Fields: json.RawMessage(`"more data"`),
	}

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
}

func TestRPCRescan(t *testing.T) {
	wallet, _ := useRPCLocal(t)
	_, err := wallet.Rescan(RescanParams{})
	if err != nil {
		t.Fatal(err)
	}
}

func TestRPCSignData(t *testing.T) {
	wallet, _ := useRPCLocal(t)
	data, err := wallet.SignData(map[string]interface{}{"hello": "world"})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", data)
}

func TestRPCBalanceAndAsset(t *testing.T) {
	wallet, _ := useRPCLocal(t)

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
	wallet, _ := useRPCLocal(t)

	// Send: 50ebdb059e5c9ad0f9fc7ac5d970b17ec2fc81bf197c9e737cf2d3ca14c5ae84
	tx, err := wallet.GetTransaction(GetTransactionParams{
		Hash: "50ebdb059e5c9ad0f9fc7ac5d970b17ec2fc81bf197c9e737cf2d3ca14c5ae84",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", tx)

	// Burn: 8383b7027694615e790bea812a3385af8f140b55734b8eb89bf8a42d0671aec7
	tx, err = wallet.GetTransaction(GetTransactionParams{
		Hash: "8383b7027694615e790bea812a3385af8f140b55734b8eb89bf8a42d0671aec7",
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", tx)

	// Receive: 37ecec82d39ea38d94240335d9fc1de01d039d52c764709f37766d36e3f5c336
	tx, err = wallet.GetTransaction(GetTransactionParams{
		Hash: "37ecec82d39ea38d94240335d9fc1de01d039d52c764709f37766d36e3f5c336",
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", tx)

	// Coinbase: 000000001fd2bf51d9c895bc200bd3e17597edd9827ac616d66884b75b55ddab
	tx, err = wallet.GetTransaction(GetTransactionParams{
		Hash: "000000001fd2bf51d9c895bc200bd3e17597edd9827ac616d66884b75b55ddab",
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", tx)
}

func TestRPCGetTransactionWithExtraData(t *testing.T) {
	wallet, _ := useRPCLocal(t)

	tx, err := wallet.GetTransaction(GetTransactionParams{
		Hash: "5459a2567c7666d902fa5042db601d50b8353cd73927d6b5c3ad4f99a1368206",
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", tx)
}

func TestRPCListTransactions(t *testing.T) {
	wallet, _ := useRPCLocal(t)

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

func TestRPCBurn(t *testing.T) {
	wallet, _ := useRPCLocal(t)

	result, err := wallet.BuildTransaction(BuildTransactionParams{
		Burn: &daemon.Burn{
			Asset:  config.XELIS_ASSET,
			Amount: 1,
		},
		Broadcast: false,
		TxAsHex:   true,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", result)
}

func TestRPCTransfer(t *testing.T) {
	wallet, _ := useRPCLocal(t)

	result, err := wallet.BuildTransaction(BuildTransactionParams{
		Transfers: []TransferOut{
			{Amount: 1, Asset: config.XELIS_ASSET, Destination: TESTING_ADDR},
		},
		Broadcast: false,
		TxAsHex:   true,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", result)
}

func TestRPCSendExtraData(t *testing.T) {
	wallet, _ := useRPCLocal(t)

	extraData := DataElement{
		Value:  10,
		Array:  []DataElement{DataElement{Value: "test"}},
		Fields: json.RawMessage(`10`),
	}

	result, err := wallet.BuildTransaction(BuildTransactionParams{
		Transfers: []TransferOut{
			{Amount: 0, Asset: config.XELIS_ASSET, Destination: MAINNET_ADDR, ExtraData: &extraData},
		},
		Broadcast: false,
		TxAsHex:   true,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", result)
}
