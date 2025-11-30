package daemon

import (
	"net/http"

	"github.com/xelis-project/xelis-go-sdk/daemon/methods"
	"github.com/xelis-project/xelis-go-sdk/rpc"
)

type RPC struct {
	http *rpc.Http
}

func NewRPC(url string) (*RPC, error) {
	http, err := rpc.NewHttp(url, nil)
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

func (d *RPC) GetInfo() (result GetInfoResult, err error) {
	_, err = d.Request(methods.GetInfo, nil, &result)
	return
}

func (d *RPC) GetHeight() (height uint64, err error) {
	_, err = d.Request(methods.GetHeight, nil, &height)
	return
}

func (d *RPC) GetTopoheight() (topoheight uint64, err error) {
	_, err = d.Request(methods.GetTopoheight, nil, &topoheight)
	return
}

func (d *RPC) GetStableHeight() (stableheight uint64, err error) {
	_, err = d.Request(methods.GetStableHeight, nil, &stableheight)
	return
}

func (d *RPC) GetStableTopoheight() (topoheight uint64, err error) {
	_, err = d.Request(methods.GetStableTopoheight, nil, &topoheight)
	return
}

func (d *RPC) GetStableBalance(params GetBalanceParams) (result GetStableBalanceResult, err error) {
	_, err = d.Request(methods.GetStableBalance, params, &result)
	return
}

func (d *RPC) GetBlockTemplate(addr string) (result GetBlockTemplateResult, err error) {
	params := map[string]string{"address": addr}
	_, err = d.Request(methods.GetBlockTemplate, params, &result)
	return
}

func (d *RPC) GetBlockAtTopoheight(params GetBlockAtTopoheightParams) (block Block, err error) {
	_, err = d.Request(methods.GetBlockAtTopoheight, params, &block)
	return
}

func (d *RPC) GetBlocksAtHeight(params GetBlocksAtHeightParams) (blocks []Block, err error) {
	_, err = d.Request(methods.GetBlocksAtHeight, params, &blocks)
	return
}

func (d *RPC) GetBlockByHash(params GetBlockByHashParams) (block Block, err error) {
	_, err = d.Request(methods.GetBlockByHash, params, &block)
	return
}

func (d *RPC) GetTopBlock(params GetTopBlockParams) (block Block, err error) {
	_, err = d.Request(methods.GetTopBlock, params, &block)
	return
}

func (d *RPC) GetNonce(addr string) (nonce GetNonceResult, err error) {
	params := map[string]string{"address": addr}
	_, err = d.Request(methods.GetNonce, params, &nonce)
	return
}

func (d *RPC) HasNonce(addr string) (hasNonce bool, err error) {
	params := map[string]string{"address": addr}
	var result map[string]bool
	_, err = d.Request(methods.HasNonce, params, &result)
	hasNonce = result["exist"]
	return
}

func (d *RPC) GetNonceAtTopoheight(params GetNonceAtTopoheightParams) (nonce VersionedNonce, err error) {
	_, err = d.Request(methods.GetNonceAtTopoheight, params, &nonce)
	return
}

func (d *RPC) GetBalance(params GetBalanceParams) (balance GetBalanceResult, err error) {
	_, err = d.Request(methods.GetBalance, params, &balance)
	return
}

func (d *RPC) HasBalance(params GetBalanceParams) (hasBalance bool, err error) {
	var result map[string]bool
	_, err = d.Request(methods.HasBalance, params, &result)
	hasBalance = result["exists"]
	return
}

func (d *RPC) GetBalanceAtTopoheight(params GetBalanceAtTopoheightParams) (balance VersionedBalance, err error) {
	_, err = d.Request(methods.GetBalanceAtTopoheight, params, &balance)
	return
}

func (d *RPC) GetAsset(assetId string) (asset AssetData, err error) {
	params := map[string]string{"asset": assetId}
	_, err = d.Request(methods.GetAsset, params, &asset)
	return
}

func (d *RPC) GetAssets(params GetAssetsParams) (assets []AssetData, err error) {
	_, err = d.Request(methods.GetAssets, params, &assets)
	return
}

func (d *RPC) CountAssets() (count uint64, err error) {
	_, err = d.Request(methods.CountAssets, nil, &count)
	return
}

func (d *RPC) CountTransactions() (count uint64, err error) {
	_, err = d.Request(methods.CountTransactions, nil, &count)
	return
}

func (d *RPC) CountAccounts() (count uint64, err error) {
	_, err = d.Request(methods.CountAccounts, nil, &count)
	return
}

func (d *RPC) GetTips() (tips []string, err error) {
	_, err = d.Request(methods.GetTips, nil, &tips)
	return
}

func (d *RPC) P2PStatus() (status P2PStatusResult, err error) {
	_, err = d.Request(methods.P2PStatus, nil, &status)
	return
}

func (d *RPC) GetDAGOrder(params GetTopoheightRangeParams) (hashes []string, err error) {
	_, err = d.Request(methods.GetDAGOrder, params, &hashes)
	return
}

func (d *RPC) SubmitBlock(params SubmitBlockParams) (result bool, err error) {
	_, err = d.Request(methods.SubmitBlock, params, &result)
	return
}

func (d *RPC) SubmitTransaction(data string) (result bool, err error) {
	params := map[string]string{"data": data}
	_, err = d.Request(methods.SubmitTransaction, params, &result)
	return
}

func (d *RPC) GetMempool() (result GetMempoolResult, err error) {
	_, err = d.Request(methods.GetMempool, nil, &result)
	return
}

func (d *RPC) GetMempoolCache(params GetMempoolCacheParams) (result GetMempoolCacheResult, err error) {
	_, err = d.Request(methods.GetMempoolCache, params, &result)
	return
}

func (d *RPC) GetTransaction(hash string) (tx TransactionResponse, err error) {
	params := map[string]string{"hash": hash}
	_, err = d.Request(methods.GetTransaction, params, &tx)
	return
}

func (d *RPC) GetTransactions(params GetTransactionsParams) (txs []TransactionResponse, err error) {
	_, err = d.Request(methods.GetTransactions, params, &txs)
	return
}

func (d *RPC) GetBlocksRangeByTopoheight(params GetTopoheightRangeParams) (blocks []Block, err error) {
	_, err = d.Request(methods.GetBlocksRangeByTopoheight, params, &blocks)
	return
}

func (d *RPC) GetBlocksRangeByHeight(params GetHeightRangeParams) (blocks []Block, err error) {
	_, err = d.Request(methods.GetBlocksRangeByHeight, params, &blocks)
	return
}

func (d *RPC) GetAccounts(params GetAccountsParams) (addresses []string, err error) {
	_, err = d.Request(methods.GetAccounts, params, &addresses)
	return
}

func (d *RPC) GetAccountHistory(addr string) (history []AccountHistory, err error) {
	params := map[string]string{"address": addr}
	_, err = d.Request(methods.GetAccountHistory, params, &history)
	return
}

func (d *RPC) GetAccountAssets(addr string) (assets []string, err error) {
	params := map[string]string{"address": addr}
	_, err = d.Request(methods.GetAccountAssets, params, &assets)
	return
}

func (d *RPC) GetPeers() (result GetPeersResult, err error) {
	_, err = d.Request(methods.GetPeers, nil, &result)
	return
}

func (d *RPC) GetDevFeeThresholds() (fees []Fee, err error) {
	_, err = d.Request(methods.GetDevFeeThresholds, nil, &fees)
	return
}

func (d *RPC) GetSizeOnDisk() (sizeOnDisk SizeOnDisk, err error) {
	_, err = d.Request(methods.GetSizeOnDisk, nil, &sizeOnDisk)
	return
}

func (d *RPC) IsTxExecutedInBlock(params IsTxExecutedInBlockParams) (executed bool, err error) {
	_, err = d.Request(methods.IsTxExecutedInBlock, params, &executed)
	return
}

func (d *RPC) GetAccountRegistrationTopoheight(addr string) (topoheight uint64, err error) {
	params := map[string]string{"address": addr}
	_, err = d.Request(methods.GetAccountRegistrationTopoheight, params, &topoheight)
	return
}

func (d *RPC) IsAccountRegistered(params IsAccountRegisteredParams) (exists bool, err error) {
	_, err = d.Request(methods.IsAccountRegistered, params, &exists)
	return
}

func (d *RPC) GetDifficulty() (result GetDifficultyResult, err error) {
	_, err = d.Request(methods.GetDifficulty, nil, &result)
	return
}

func (d *RPC) ValidateAddress(params ValidateAddressParams) (result ValidateAddressResult, err error) {
	_, err = d.Request(methods.ValidateAddress, params, &result)
	return
}

func (d *RPC) ExtractKeyFromAddress(params ExtractKeyFromAddressParams) (key ExtractKeyFromAddressResult, err error) {
	_, err = d.Request(methods.ExtractKeyFromAddress, params, &key)
	return
}

func (d *RPC) GetMinerWork(params GetMinerWorkParams) (result GetMinerWorkResult, err error) {
	_, err = d.Request(methods.GetMinerWork, params, &result)
	return
}

func (d *RPC) SplitAddress(params SplitAddressParams) (result SplitAddressResult, err error) {
	_, err = d.Request(methods.SplitAddress, params, &result)
	return
}

func (d *RPC) GetHardForks() (result []HardFork, err error) {
	_, err = d.Request(methods.GetHardForks, nil, &result)
	return
}

func (d *RPC) GetEstimatedFeeRates() (result FeeRatesEstimated, err error) {
	_, err = d.Request(methods.GetEstimatedFeeRates, nil, &result)
	return
}

func (d *RPC) GetPrunedTopoheight() (result uint64, err error) {
	_, err = d.Request(methods.GetPrunedTopoheight, nil, &result)
	return
}

func (d *RPC) GetTransactionExecutor(params GetTransactionExecutorParams) (result GetTransactionExecutorResult, err error) {
	_, err = d.Request(methods.GetTransactionExecutor, params, &result)
	return
}

func (d *RPC) HasMultisigAtTopoheight(params HasMultisigAtTopoheightParams) (result bool, err error) {
	_, err = d.Request(methods.HasMultisigAtTopoheight, params, &result)
	return
}

func (d *RPC) GetMultisigAtTopoheight(params GetMultisigAtTopoheightParams) (result GetMultisigAtTopoheightResult, err error) {
	_, err = d.Request(methods.GetMultisigAtTopoheight, params, &result)
	return
}

func (d *RPC) GetMultisig(params GetMultisigParams) (result GetMultisigResult, err error) {
	_, err = d.Request(methods.GetMultisig, params, &result)
	return
}

func (d *RPC) HasMultisig(params HasMultisigParams) (result bool, err error) {
	_, err = d.Request(methods.HasMultisig, params, &result)
	return
}

func (d *RPC) GetContractOutputs(params GetContractOutputsParams) (result []ContractOutput, err error) {
	var outputs []interface{}
	_, err = d.Request(methods.GetContractOutputs, params, &outputs)
	result = parseContractOutputs(outputs)
	return
}

func (d *RPC) GetContractModule(params GetContractModuleParams) (result GetContractModuleResult, err error) {
	_, err = d.Request(methods.GetContractModule, params, &result)
	return
}

func (d *RPC) GetContractData(params GetContractDataParams) (result interface{}, err error) {
	_, err = d.Request(methods.GetContractData, params, &result)
	return
}

func (d *RPC) GetContractDataAtTopoheight(params GetContractDataAtTopoheightParams) (result interface{}, err error) {
	_, err = d.Request(methods.GetContractDataAtTopoheight, params, &result)
	return
}

func (d *RPC) GetContractBalance(params GetContractBalanceParams) (result GetContractBalanceResult, err error) {
	_, err = d.Request(methods.GetContractBalance, params, &result)
	return
}

func (d *RPC) GetContractBalanceAtTopoheight(params GetContractBalanceAtTopoheightParams) (result GetContractBalanceAtTopoheightResult, err error) {
	_, err = d.Request(methods.GetContractBalanceAtTopoheight, params, &result)
	return
}

func (d *RPC) CountContracts() (result uint64, err error) {
	_, err = d.Request(methods.CountAccounts, nil, &result)
	return
}

func (d *RPC) MakeIntegratedAddress(params MakeIntegratedAddressParams) (result string, err error) {
	_, err = d.Request(methods.MakeIntegratedAddress, params, &result)
	return
}

func (d *RPC) DecryptExtraData(params DecryptExtraDataParams) (result interface{}, err error) {
	_, err = d.Request(methods.DecryptExtraData, params, &result)
	return
}

func parseContractOutputs(outputs []interface{}) (result []ContractOutput) {
	for _, output := range outputs {
		switch out := output.(type) {
		case map[string]interface{}:
			for key, value := range out {
				switch key {
				case "exit_code":
					// if the value is nil it means the code failed and did not execute
					exit_code, ok := value.(float64)
					if !ok {
						break
					}

					result = append(result, ContractOutputExitCode{
						ExitCode: uint64(exit_code),
					})
				case "refund_gas":
					refund_gas, ok := value.(map[string]interface{})
					if !ok {
						break
					}

					amount, ok := refund_gas["amount"].(float64)
					if !ok {
						break
					}

					result = append(result, ContractOutputRefundGas{
						Amount: uint64(amount),
					})
				case "transfer":
					transfer, ok := value.(map[string]interface{})
					if !ok {
						break
					}

					amount, ok := transfer["amount"].(float64)
					if !ok {
						break
					}

					asset, ok := transfer["asset"].(string)
					if !ok {
						break
					}

					destination, ok := transfer["destination"].(string)
					if !ok {
						break
					}

					result = append(result, ContractOutputTransfer{
						Amount:      uint64(amount),
						Asset:       asset,
						Destination: destination,
					})
				}
			}
		case string:
			switch out {
			case "refund_deposits":
				result = append(result, ContractOutputRefundDeposits{})
			}
		}
	}

	return
}
