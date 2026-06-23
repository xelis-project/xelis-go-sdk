package wallet

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/xelis-project/xelis-go-sdk/data"
	"github.com/xelis-project/xelis-go-sdk/rpc"
	"github.com/xelis-project/xelis-go-sdk/wallet/methods"
)

type RPC struct {
	http *rpc.Http
}

func setAuthHeader(header http.Header, username string, password string) {
	auth := fmt.Sprintf("%s:%s", username, password)
	buf := bytes.Buffer{}
	encoder := base64.NewEncoder(base64.StdEncoding, &buf)
	encoder.Write([]byte(auth))
	encoder.Close()

	header.Set("Authorization", fmt.Sprintf("Basic %s", buf.String()))
}

func NewRPC(url string, username string, password string) (*RPC, error) {
	header := http.Header{}
	setAuthHeader(header, username, password)
	http, err := rpc.NewHttp(url, header)
	if err != nil {
		return nil, err
	}

	daemon := &RPC{
		http,
	}

	return daemon, nil
}

func (d *RPC) BatchRequest(requests []rpc.RPCRequest, result []interface{}) (res *http.Response, errs []error) {
	return d.http.BatchRequest(requests, result)
}

func (d *RPC) Request(method string, params interface{}, result interface{}) (res *http.Response, err error) {
	return d.http.Request(method, params, result)
}

func (d *RPC) GetVersion() (version string, err error) {
	_, err = d.Request(methods.GetVersion, nil, &version)
	return
}

func (d *RPC) GetNetwork() (network string, err error) {
	_, err = d.Request(methods.GetNetwork, nil, &network)
	return
}

func (d *RPC) GetNonce() (nonce uint64, err error) {
	_, err = d.Request(methods.GetNonce, nil, &nonce)
	return
}

func (d *RPC) GetTopoheight() (topoheight uint64, err error) {
	_, err = d.Request(methods.GetTopoheight, nil, &topoheight)
	return
}

func (d *RPC) GetAddress(params GetAddressParams) (address string, err error) {
	_, err = d.Request(methods.GetAddress, params, &address)
	return
}

func (d *RPC) SplitAddress(params SplitAddressParams) (result SplitAddressResult, err error) {
	_, err = d.Request(methods.SplitAddress, params, &result)
	return
}

func (d *RPC) Rescan(params RescanParams) (success bool, err error) {
	_, err = d.Request(methods.Rescan, params, &success)
	return
}

func (d *RPC) GetBalance(params GetBalanceParams) (balance uint64, err error) {
	_, err = d.Request(methods.GetBalance, params, &balance)
	return
}

func (d *RPC) HasBalance(params GetBalanceParams) (exists bool, err error) {
	_, err = d.Request(methods.HasBalance, params, &exists)
	return
}

func (d *RPC) GetTrackedAssets(params GetAssetsParams) (assets []string, err error) {
	_, err = d.Request(methods.GetTrackedAssets, params, &assets)
	return
}

func (d *RPC) IsAssetTracked(params IsAssetTrackedParams) (tracked bool, err error) {
	_, err = d.Request(methods.IsAssetTracked, params, &tracked)
	return
}

func (d *RPC) TrackAsset(params TrackAssetParams) (tracked bool, err error) {
	_, err = d.Request(methods.TrackAsset, params, &tracked)
	return
}

func (d *RPC) UntrackAsset(params TrackAssetParams) (untracked bool, err error) {
	_, err = d.Request(methods.UntrackAsset, params, &untracked)
	return
}

func (d *RPC) GetAssetPrecision(params GetAssetPrecisionParams) (decimals int, err error) {
	_, err = d.Request(methods.GetAssetPrecision, params, &decimals)
	return
}

func (d *RPC) GetAssets(params GetAssetsParams) (assets []GetAssetsEntry, err error) {
	_, err = d.Request(methods.GetAssets, params, &assets)
	return
}

func (d *RPC) GetAsset(params GetAssetParams) (asset Asset, err error) {
	_, err = d.Request(methods.GetAsset, params, &asset)
	return
}

func (d *RPC) GetTransaction(params GetTransactionParams) (transaction TransactionEntry, err error) {
	_, err = d.Request(methods.GetTransaction, params, &transaction)
	return
}

func (d *RPC) SearchTransaction(params SearchTransactionParams) (result SearchTransactionResult, err error) {
	_, err = d.Request(methods.SearchTransaction, params, &result)
	return
}

func (d *RPC) DumpTransaction(params GetTransactionParams) (tx string, err error) {
	_, err = d.Request(methods.DumpTransaction, params, &tx)
	return
}

func (d *RPC) BuildTransaction(params BuildTransactionParams) (result TransactionResponse, err error) {
	if err = checkFeeBuilder(params.Fee); err != nil {
		return
	}

	_, err = d.Request(methods.BuildTransaction, params, &result)
	return
}

func (d *RPC) BuildTransactionOffline(params BuildTransactionOfflineParams) (result TransactionResponse, err error) {
	_, err = d.Request(methods.BuildTransactionOffline, params, &result)
	return
}

func (d *RPC) BuildUnsignedTransaction(params BuildUnsignedTransactionParams) (result UnsignedTransactionResponse, err error) {
	_, err = d.Request(methods.BuildUnsignedTransaction, params, &result)
	return
}

func (d *RPC) SignUnsignedTransaction(params SignUnsignedTransactionParams) (result SignatureId, err error) {
	_, err = d.Request(methods.SignUnsignedTransaction, params, &result)
	return
}

func (d *RPC) FinalizeUnsignedTransaction(params FinalizeUnsignedTransactionParams) (result TransactionResponse, err error) {
	_, err = d.Request(methods.FinalizeUnsignedTransaction, params, &result)
	return
}

func (d *RPC) GetPendingTransactions() (txs []TransactionPending, err error) {
	_, err = d.Request(methods.GetPendingTransactions, nil, &txs)
	return
}

func (d *RPC) ClearTxCache() (result bool, err error) {
	_, err = d.Request(methods.ClearTxCache, nil, &result)
	return
}

func (d *RPC) ListTransactions(params ListTransactionsParams) (txs []TransactionEntry, err error) {
	_, err = d.Request(methods.ListTransactions, params, &txs)
	return
}

func (d *RPC) IsOnline() (online bool, err error) {
	_, err = d.Request(methods.IsOnline, nil, &online)
	return
}

func (d *RPC) SetOnlineMode(params SetOnlineModeParams) (success bool, err error) {
	_, err = d.Request(methods.SetOnlineMode, params, &success)
	return
}

func (d *RPC) SetOfflineMode() (success bool, err error) {
	_, err = d.Request(methods.SetOfflineMode, nil, &success)
	return
}

func (d *RPC) SignData(data data.Element) (signature string, err error) {
	// note: SignData parse params null value to {}
	_, err = d.Request(methods.SignData, data, &signature)
	return
}

func (d *RPC) VerifySignedData(params VerifySignedDataParams) (valid bool, err error) {
	_, err = d.Request(methods.VerifySignedData, params, &valid)
	return
}

func (d *RPC) EstimateFees(params EstimateFeesParams) (amount uint64, err error) {
	_, err = d.Request(methods.EstimateFees, params, &amount)
	return
}

func (d *RPC) EstimateExtraDataSize(params EstimateExtraDataSizeParams) (result EstimateExtraDataSizeResult, err error) {
	_, err = d.Request(methods.EstimateExtraDataSize, params, &result)
	return
}

func (d *RPC) NetworkInfo() (result NetworkInfoResult, err error) {
	_, err = d.Request(methods.NetworkInfo, nil, &result)
	return
}

func (d *RPC) DecryptExtraData(params DecryptExtraDataParams) (result PlaintextExtraData, err error) {
	_, err = d.Request(methods.DecryptExtraData, params, &result)
	return
}

func (d *RPC) DecryptCiphertext(params DecryptCiphertextParams) (result *uint64, err error) {
	_, err = d.Request(methods.DecryptCiphertext, params, &result)
	return
}

func (d *RPC) CreateOwnershipProof(params CreateOwnershipProofParams) (result interface{}, err error) {
	_, err = d.Request(methods.CreateOwnershipProof, params, &result)
	return
}

func (d *RPC) CreateBalanceProof(params CreateBalanceProofParams) (result interface{}, err error) {
	_, err = d.Request(methods.CreateBalanceProof, params, &result)
	return
}

func (d *RPC) VerifyHumanReadableProof(params VerifyHumanReadableProofParams) (valid bool, err error) {
	_, err = d.Request(methods.VerifyHumanReadableProof, params, &valid)
	return
}

func (d *RPC) GetMatchingKeys(params GetMatchingKeysParams) (result []interface{}, err error) {
	_, err = d.Request(methods.GetMatchingKeys, params, &result)
	return
}

func (d *RPC) CountMatchingEntries(params CountMatchingEntriesParams) (result uint64, err error) {
	_, err = d.Request(methods.CountMatchingEntries, params, &result)
	return
}

func (d *RPC) GetValueFromKey(params GetValueFromKeyParams) (result interface{}, err error) {
	_, err = d.Request(methods.GetValueFromKey, params, &result)
	return
}

func (d *RPC) Store(params StoreParams) (result bool, err error) {
	_, err = d.Request(methods.Store, params, &result)
	return
}

func (d *RPC) Delete(params interface{}) (result bool, err error) {
	_, err = d.Request(methods.Delete, params, &result)
	return
}

func (d *RPC) DeleteTreeEntries(params DeleteTreeEntriesParams) (result bool, err error) {
	_, err = d.Request(methods.DeleteTreeEntries, params, &result)
	return
}

func (d *RPC) HasKey(params HasKeyParams) (result bool, err error) {
	_, err = d.Request(methods.HasKey, params, &result)
	return
}

func (d *RPC) QueryDB(params QueryDBParams) (result QueryResult, err error) {
	_, err = d.Request(methods.QueryDB, params, &result)
	return
}

func checkFeeBuilder(fee *FeeBuilder) error {
	if fee != nil {
		fixedModes := 0
		for _, isSet := range []bool{fee.Fixed != nil, fee.Value != nil} {
			if isSet {
				fixedModes++
			}
		}
		extraModes := 0
		for _, isSet := range []bool{fee.Extra != nil, fee.Multiplier != nil} {
			if isSet {
				extraModes++
			}
		}
		if fixedModes > 1 || extraModes > 1 || (fixedModes > 0 && extraModes > 0) {
			return fmt.Errorf("you cannot set multiple fee modes in FeeBuilder")
		}
		if fee.Extra != nil && fee.Extra.Tip != nil && fee.Extra.Multiplier != nil {
			return fmt.Errorf("you cannot set both Tip and Multiplier in ExtraFeeMode")
		}
	}

	return nil
}
