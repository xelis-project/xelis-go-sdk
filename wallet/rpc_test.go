package wallet

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/xelis-project/xelis-go-sdk/config"
	da "github.com/xelis-project/xelis-go-sdk/daemon"
	d "github.com/xelis-project/xelis-go-sdk/data"
	"github.com/xelis-project/xelis-go-sdk/signature"
	"github.com/xelis-project/xelis-go-sdk/xvm"
)

const TESTING_ADDR = "xet:qf5u2p46jpgqmypqc2xwtq25yek2t7qhnqtdhw5kpfwcrlavs5asq0r83r7"
const MAINNET_ADDR = "xel:as3mgjlevw5ve6k70evzz8lwmsa5p0lgws2d60fulxylnmeqrp9qqukwdfg"

// cargo run --bin xelis_wallet -- --wallet-path ./wallets/test --password test --network dev --rpc-password test --rpc-bind-address 127.0.0.1:8081 --rpc-username test

func prepareRPC(t *testing.T) (wallet *RPC) {
	// wallet, err := NewRPC("http://192.168.1.53:8081/json_rpc", "test", "test")
	wallet, err := NewRPC(config.LOCAL_WALLET_RPC, "test", "test")
	if err != nil {
		t.Fatal(err)
	}

	return
}

func prepareDaemonRPC(t *testing.T) (daemon *da.RPC) {
	daemon, err := da.NewRPC(config.LOCAL_NODE_RPC)
	if err != nil {
		t.Fatal(err)
	}

	return
}

func TestRPCGetVersion(t *testing.T) {
	wallet := prepareRPC(t)

	version, err := wallet.GetVersion()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", version)
}

func TestRPCGetNetwork(t *testing.T) {
	wallet := prepareRPC(t)

	network, err := wallet.GetNetwork()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", network)
}

func TestRPCGetNonce(t *testing.T) {
	wallet := prepareRPC(t)

	nonce, err := wallet.GetNonce()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", nonce)
}

func TestRPCGetTopoheight(t *testing.T) {
	wallet := prepareRPC(t)

	topo, err := wallet.GetTopoheight()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", topo)
}

func TestRPCGetAddress(t *testing.T) {
	wallet := prepareRPC(t)

	address, err := wallet.GetAddress(GetAddressParams{})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", address)
}

func TestRPCIsOnline(t *testing.T) {
	wallet := prepareRPC(t)

	isOnline, err := wallet.IsOnline()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", isOnline)
}

func TestRPCIntegratedAddress(t *testing.T) {
	wallet := prepareRPC(t)

	var integratedData interface{} = map[string]interface{}{"hello": "world"}

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
	wallet := prepareRPC(t)
	_, err := wallet.Rescan(RescanParams{})
	if err != nil {
		t.Fatal(err)
	}
}

func TestRPCSignDataFields(t *testing.T) {
	wallet := prepareRPC(t)

	element := d.Element{
		Fields: map[d.Value]d.Element{
			"hello": d.Element{Value: "world"},
		},
	}

	data, err := wallet.SignData(element)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", data)
}

func TestRPCSignDataArray(t *testing.T) {
	wallet := prepareRPC(t)

	element := d.Element{
		Array: []d.Element{
			d.Element{Value: 23469234},
		},
	}

	data, err := wallet.SignData(element)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", data)
}

func TestRPCSignDataValue(t *testing.T) {
	wallet := prepareRPC(t)

	element := d.Element{
		Value: 3456349494,
	}

	data, err := wallet.SignData(element)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", data)
}

func TestSignature(t *testing.T) {
	wallet := prepareRPC(t)
	daemon := prepareDaemonRPC(t)

	addr, err := wallet.GetAddress(GetAddressParams{})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", addr)

	publicKey, err := daemon.ExtractKeyFromAddress(da.ExtractKeyFromAddressParams{
		Address: addr,
		AsHex:   false,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", publicKey)

	// data := map[string]interface{}{"hello": "world"}
	data := d.Element{Fields: map[d.Value]d.Element{
		"hello": d.Element{Value: "world"},
	}}

	dataSigned, err := wallet.SignData(data)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", dataSigned)

	// serialized bytes equivalent to map[string]interface{}{"hello": "world"}
	// b := []byte{2, 1, 1, 5, 104, 101, 108, 108, 111, 0, 1, 5, 119, 111, 114, 108, 100}
	b, err := data.ToBytes()
	if err != nil {
		return
	}

	valid, err := signature.Verify(*publicKey.Bytes, dataSigned, b)
	if err != nil {
		t.Fatal(err)
	}

	if !valid {
		t.Fatal("invalid verification")
	}
}

func TestRPCBalanceAndAsset(t *testing.T) {
	wallet := prepareRPC(t)

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

	assets, err := wallet.GetTrackedAssets(GetAssetsParams{})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", assets)
}

func TestRPCGetTransaction(t *testing.T) {
	wallet := prepareRPC(t)

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
	wallet := prepareRPC(t)

	tx, err := wallet.GetTransaction(GetTransactionParams{
		Hash: "5459a2567c7666d902fa5042db601d50b8353cd73927d6b5c3ad4f99a1368206",
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", tx)
}

func TestRPCListTransactions(t *testing.T) {
	wallet := prepareRPC(t)

	txs, err := wallet.ListTransactions(ListTransactionsParams{
		AcceptOutgoing: true,
		AcceptIncoming: true,
		AcceptCoinbase: false,
		AcceptBurn:     false,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", txs)
}

func TestRPCBurn(t *testing.T) {
	wallet := prepareRPC(t)

	result, err := wallet.BuildTransaction(BuildTransactionParams{
		Burn: &da.Burn{
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
	wallet := prepareRPC(t)

	result, err := wallet.BuildTransaction(BuildTransactionParams{
		Transfers: []TransferBuilder{
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
	wallet := prepareRPC(t)

	var extraData interface{} = map[string]interface{}{"hello": "world"}

	result, err := wallet.BuildTransaction(BuildTransactionParams{
		Transfers: []TransferBuilder{
			{Amount: 0, Asset: config.XELIS_ASSET, Destination: TESTING_ADDR, ExtraData: &extraData},
		},
		Broadcast: false,
		TxAsHex:   true,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", result)

	first_transfer := (*result.Data.Transfers)[0]

	result2, err := wallet.DecryptExtraData(DecryptExtraDataParams{
		ExtraData: *first_transfer.ExtraData,
		Role:      TxSenderRole,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", result2)
}

func TestRPCSendWithFeeBuilder(t *testing.T) {
	wallet := prepareRPC(t)

	// you can only have one of both
	// either use Multiplier or Value
	//feeMultiplier := float64(1)
	feeValue := uint64(1)

	result, err := wallet.BuildTransaction(BuildTransactionParams{
		Transfers: []TransferBuilder{
			{Amount: 0, Asset: config.XELIS_ASSET, Destination: MAINNET_ADDR},
		},
		Fee: &FeeBuilder{
			//Multiplier: &feeMultiplier,
			Value: &feeValue,
		},
		Broadcast: false,
		TxAsHex:   true,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", result)
}

func TestRPCDeploySC(t *testing.T) {
	wallet := prepareRPC(t)

	/*
		entry main(a: u64, b: u64) {
			println("Hello, World!");
			println(a + b);
			storage().store("test", "hello");
			return 0;
		}
	*/

	hex_program := "00000000000400090d48656c6c6f2c20576f726c64210009047465737400090568656c6c6f0004000000000000000000010000002e02000002010000000014550000010100000101001a14550000011468000000000100000200146b0001020003001000010000"
	//hex_program := "00000000000200090d48656c6c6f2c20576f726c64210004000000000000000000010000000c00000014550000010001001000010000"

	fee := uint64(50000)
	result, err := wallet.BuildTransaction(BuildTransactionParams{
		DeployContract: &hex_program,
		Broadcast:      true,
		Fee:            &FeeBuilder{Value: &fee},
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", result.Hash)
}

func TestRPCInvokeSC(t *testing.T) {
	wallet := prepareRPC(t)

	fee := float64(3)

	result, err := wallet.BuildTransaction(BuildTransactionParams{
		InvokeContract: &InvokeContractBuilder{
			Contract: "f8ffd882e1907c501c23a86c3947b8222cc544a55d448cadcb28798e5f554be0",
			MaxGas:   300,
			EntryId:  0,
			Parameters: []xvm.ValueCell{
				xvm.NewPrimitive(xvm.U64, uint64(1)),
				xvm.NewPrimitive(xvm.U64, uint64(2)),
			},
			Deposits: map[string]ContractDepositBuilder{
				config.XELIS_ASSET: {Amount: 100, Private: false},
			},
		},
		Fee:       &FeeBuilder{Multiplier: &fee},
		Broadcast: true,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", result.Hash)
}

func TestRPCEstimateFees(t *testing.T) {
	wallet := prepareRPC(t)

	result, err := wallet.EstimateFees(EstimateFeesParams{
		Transfers: []TransferBuilder{
			{Amount: 100, Asset: config.XELIS_ASSET, Destination: TESTING_ADDR},
			{Amount: 200, Asset: config.XELIS_ASSET, Destination: TESTING_ADDR},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", result)
}

func TestRPCNetworkInfo(t *testing.T) {
	wallet := prepareRPC(t)

	result, err := wallet.NetworkInfo()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", result)
}

func TestRPCClearTxCache(t *testing.T) {
	wallet := prepareRPC(t)

	result, err := wallet.ClearTxCache()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", result)
}

func TestRPCEstimateExtraDataSize(t *testing.T) {
	wallet := prepareRPC(t)

	var integratedData interface{} = map[string]interface{}{"hello": "world"}

	addr, err := wallet.GetAddress(GetAddressParams{IntegratedData: &integratedData})
	if err != nil {
		t.Fatal(err)
	}

	result, err := wallet.EstimateExtraDataSize(EstimateExtraDataSizeParams{
		Destinations: []string{addr},
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", result)
}

func TestRPCDecryptCiphertext(t *testing.T) {
	wallet := prepareRPC(t)

	tx, err := wallet.BuildTransaction(BuildTransactionParams{
		Transfers: []TransferBuilder{
			{Amount: 123, Asset: config.XELIS_ASSET, Destination: "xet:5e8entraya4v3264rdyrd33dhtj5f89mxt95s423ndyc66sr742sqaagsnv"},
		},
		Broadcast: false,
		TxAsHex:   true,
	})
	if err != nil {
		t.Fatal(err)
	}

	if tx.Data.Transfers != nil {
		t1 := (*tx.Data.Transfers)[0]
		t.Log(t1)

		result, err := wallet.DecryptCiphertext(DecryptCiphertextParams{
			Ciphertext: CompressedCiphertext{
				Commitment: t1.Commitment,
				Handle:     t1.SenderHandle,
			},
		})
		if err != nil {
			t.Fatal(err)
		}

		t.Logf("%+v", result)
	} else {
		t.Fail()
	}
}

func TestRPCGetAsset(t *testing.T) {
	wallet := prepareRPC(t)

	asset, err := wallet.GetAsset(GetAssetParams{
		Asset: config.XELIS_ASSET,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", asset)
}

func TestRPCGetAssets(t *testing.T) {
	wallet := prepareRPC(t)

	assets, err := wallet.GetAssets(GetAssetsParams{})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", assets)
}

func TestRPCWalletDBKeyStore(t *testing.T) {
	wallet := prepareRPC(t)

	tree := "my_app"

	result, err := wallet.Store(StoreParams{
		Tree:  tree,
		Key:   "test",
		Value: 100,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", result)

	result2, err := wallet.HasKey(HasKeyParams{
		Tree: tree,
		Key:  "test",
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", result2)

	result3, err := wallet.GetValueFromKey(GetValueFromKeyParams{
		Tree: tree,
		Key:  "test",
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", result3)

	result4, err := wallet.CountMatchingEntries(CountMatchingEntriesParams{
		Tree: tree,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", result4)

	result5, err := wallet.Delete(StoreParams{
		Tree: tree,
		Key:  "test",
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", result5)

	result6, err := wallet.DeleteTreeEntries(DeleteTreeEntriesParams{
		Tree: tree,
	})

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", result6)
}

func TestRPCWalletCountMatchingEntries(t *testing.T) {
	wallet := prepareRPC(t)

	tree := "my_app_2"

	result, err := wallet.Store(StoreParams{
		Tree:  tree,
		Key:   "test",
		Value: 100,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", result)

	equalQuery := &Query{QueryValue: &QueryValue{Equal: 100}}

	queryJson, err := json.Marshal(equalQuery)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(queryJson))

	result2, err := wallet.CountMatchingEntries(CountMatchingEntriesParams{
		Tree:  tree,
		Value: equalQuery,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", result2)

	greaterQuery := &Query{QueryValue: &QueryValue{QueryNumber: &QueryNumber{Greater: 50}}}

	queryJson, err = json.Marshal(greaterQuery)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(queryJson))

	result3, err := wallet.CountMatchingEntries(CountMatchingEntriesParams{
		Tree:  tree,
		Value: greaterQuery,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", result3)
}

func TestRPCWalletQueryDB(t *testing.T) {
	wallet := prepareRPC(t)

	tree := "my_app_3"

	for i := 0; i < 10; i++ {
		_, err := wallet.Store(StoreParams{
			Tree:  tree,
			Key:   fmt.Sprintf("i_%d", i),
			Value: fmt.Sprintf("i_%d", i),
		})
		if err != nil {
			t.Fatal(err)
		}

		_, err = wallet.Store(StoreParams{
			Tree:  tree,
			Key:   i,
			Value: fmt.Sprintf("test_%d", i),
		})
		if err != nil {
			t.Fatal(err)
		}
	}

	result, err := wallet.QueryDB(QueryDBParams{
		Tree: tree,
		Key:  &Query{QueryValue: &QueryValue{ContainsValue: "test_"}},
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", result)
}
