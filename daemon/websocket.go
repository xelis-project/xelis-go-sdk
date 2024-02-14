package daemon

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/xelis-project/xelis-go-sdk/lib"
)

type WebSocket struct {
	ws *lib.WebSocket
}

func NewWebSocket(ctx context.Context, url string) (*WebSocket, error) {
	ws, err := lib.NewWebSocket(ctx, url, nil)
	if err != nil {
		return nil, err
	}

	daemonWS := &WebSocket{
		ws: ws,
	}

	return daemonWS, nil
}

func (w *WebSocket) Close() error {
	return w.ws.Close()
}

func (w *WebSocket) CloseEvent(event lib.RPCEvent) error {
	return w.ws.CloseEvent(event)
}

func (w *WebSocket) NewBlockFunc(onData func(Block, error)) error {
	return w.ws.ListenEventFunc(NewBlock, func(res lib.RPCResponse) {
		var result Block
		if res.Error != nil {
			onData(result, fmt.Errorf(res.Error.Message))
			return
		}

		err := json.Unmarshal(res.Result, &result)
		onData(result, err)
	})
}

func (w *WebSocket) TransactionAddedInMempoolFunc(onData func(Transaction, error)) error {
	return w.ws.ListenEventFunc(TransactionAddedInMempool, func(res lib.RPCResponse) {
		var result Transaction
		if res.Error != nil {
			onData(result, fmt.Errorf(res.Error.Message))
			return
		}

		err := json.Unmarshal(res.Result, &result)
		onData(result, err)
	})
}

func (w *WebSocket) BlockOrderedFunc(onData func(Block, error)) error {
	return w.ws.ListenEventFunc(BlockOrdered, func(res lib.RPCResponse) {
		var result Block
		if res.Error != nil {
			onData(result, fmt.Errorf(res.Error.Message))
			return
		}

		err := json.Unmarshal(res.Result, &result)
		onData(result, err)
	})
}

func (w *WebSocket) TransactionExecutedFunc(onData func(TransactionExecutedResult, error)) error {
	return w.ws.ListenEventFunc(TransactionExecuted, func(res lib.RPCResponse) {
		var result TransactionExecutedResult
		if res.Error != nil {
			onData(result, fmt.Errorf(res.Error.Message))
			return
		}

		err := json.Unmarshal(res.Result, &result)
		onData(result, err)
	})
}

func (w *WebSocket) PeerConnectedFunc(onData func(Peer, error)) error {
	return w.ws.ListenEventFunc(PeerConnected, func(res lib.RPCResponse) {
		var result Peer
		if res.Error != nil {
			onData(result, fmt.Errorf(res.Error.Message))
			return
		}

		err := json.Unmarshal(res.Result, &result)
		onData(result, err)
	})
}

func (w *WebSocket) PeerDisconnectedFunc(onData func(uint64, error)) error {
	return w.ws.ListenEventFunc(PeerDisconnected, func(res lib.RPCResponse) {
		var id uint64
		if res.Error != nil {
			onData(id, fmt.Errorf(res.Error.Message))
			return
		}

		err := json.Unmarshal(res.Result, &id)
		onData(id, err)
	})
}

func (w *WebSocket) PeerStateUpdatedFunc(onData func(Peer, error)) error {
	return w.ws.ListenEventFunc(PeerStateUpdated, func(res lib.RPCResponse) {
		var result Peer
		if res.Error != nil {
			onData(result, fmt.Errorf(res.Error.Message))
			return
		}

		err := json.Unmarshal(res.Result, &result)
		onData(result, err)
	})
}

func (w *WebSocket) GetVersion() (version string, err error) {
	res, err := w.ws.Call(GetVersion, nil)
	if err != nil {
		return
	}

	if res.Error != nil {
		err = fmt.Errorf(res.Error.Message)
		return
	}

	err = json.Unmarshal(res.Result, &version)
	return
}

func (w *WebSocket) GetInfo() (result GetInfoResult, err error) {
	res, err := w.ws.Call(GetInfo, nil)
	if err != nil {
		return
	}

	if res.Error != nil {
		err = fmt.Errorf(res.Error.Message)
		return
	}

	err = json.Unmarshal(res.Result, &result)
	return
}

func (w *WebSocket) GetHeight() (height uint64, err error) {
	res, err := w.ws.Call(GetHeight, nil)
	if err != nil {
		return
	}

	if res.Error != nil {
		err = fmt.Errorf(res.Error.Message)
		return
	}

	err = json.Unmarshal(res.Result, &height)
	return
}

func (w *WebSocket) GetTopoheight() (topoheight uint64, err error) {
	res, err := w.ws.Call(GetTopoHeight, nil)
	if err != nil {
		return
	}

	if res.Error != nil {
		err = fmt.Errorf(res.Error.Message)
		return
	}

	err = json.Unmarshal(res.Result, &topoheight)
	return
}

func (w *WebSocket) GetStableheight() (stableheight uint64, err error) {
	res, err := w.ws.Call(GetStableHeight, nil)
	if err != nil {
		return
	}

	if res.Error != nil {
		err = fmt.Errorf(res.Error.Message)
		return
	}

	err = json.Unmarshal(res.Result, &stableheight)
	return
}

func (w *WebSocket) GetBlockTemplate(addr string) (result GetBlockTemplateResult, err error) {
	params := map[string]string{"address": addr}
	res, err := w.ws.Call(GetBlockTemplate, params)
	if err != nil {
		return
	}

	if res.Error != nil {
		err = fmt.Errorf(res.Error.Message)
		return
	}

	err = json.Unmarshal(res.Result, &result)
	return
}

func (w *WebSocket) GetBlockAtTopoheight(params GetBlockAtTopoheightParams) (block Block, err error) {
	res, err := w.ws.Call(GetBlockAtTopoheight, params)
	if err != nil {
		return
	}

	if res.Error != nil {
		err = fmt.Errorf(res.Error.Message)
		return
	}

	err = json.Unmarshal(res.Result, &block)
	return
}

func (w *WebSocket) GetBlocksAtHeight(params GetBlockAtTopoheightParams) (blocks []Block, err error) {
	res, err := w.ws.Call(GetBlocksAtHeight, params)
	if err != nil {
		return
	}

	if res.Error != nil {
		err = fmt.Errorf(res.Error.Message)
		return
	}

	err = json.Unmarshal(res.Result, &blocks)
	return
}

func (w *WebSocket) GetBlockByHash(params GetBlockByHashParams) (block Block, err error) {
	res, err := w.ws.Call(GetBlockByHash, params)
	if err != nil {
		return
	}

	if res.Error != nil {
		err = fmt.Errorf(res.Error.Message)
		return
	}

	err = json.Unmarshal(res.Result, &block)
	return
}

func (w *WebSocket) GetTopBlock(params GetTopBlockParams) (block Block, err error) {
	res, err := w.ws.Call(GetTopBlock, params)
	if err != nil {
		return
	}

	if res.Error != nil {
		err = fmt.Errorf(res.Error.Message)
		return
	}

	err = json.Unmarshal(res.Result, &block)
	return
}

func (w *WebSocket) GetNonce(addr string) (nonce GetNonceResult, err error) {
	params := map[string]string{"address": addr}
	res, err := w.ws.Call(GetNonce, params)
	if err != nil {
		return
	}

	if res.Error != nil {
		err = fmt.Errorf(res.Error.Message)
		return
	}

	err = json.Unmarshal(res.Result, &nonce)
	return
}

func (w *WebSocket) HasNonce(addr string) (hasNonce bool, err error) {
	params := map[string]string{"address": addr}
	res, err := w.ws.Call(HasNonce, params)
	if err != nil {
		return
	}

	if res.Error != nil {
		err = fmt.Errorf(res.Error.Message)
		return
	}

	var result map[string]bool
	err = json.Unmarshal(res.Result, &result)
	hasNonce = result["exist"]
	return
}

func (w *WebSocket) GetBalance(params GetBalanceParams) (balance GetBalanceResult, err error) {
	res, err := w.ws.Call(GetBalance, params)
	if err != nil {
		return
	}

	if res.Error != nil {
		err = fmt.Errorf(res.Error.Message)
		return
	}

	err = json.Unmarshal(res.Result, &balance)
	return
}

func (w *WebSocket) HasBalance(params GetBalanceParams) (hasBalance bool, err error) {
	res, err := w.ws.Call(HasBalance, params)
	if err != nil {
		return
	}

	if res.Error != nil {
		err = fmt.Errorf(res.Error.Message)
		return
	}

	var result map[string]bool
	err = json.Unmarshal(res.Result, &result)
	hasBalance = result["exists"]
	return
}

func (w *WebSocket) GetBalanceAtTopoheight(params GetBalanceAtTopoheightParams) (balance Balance, err error) {
	res, err := w.ws.Call(GetBalanceAtTopoheight, params)
	if err != nil {
		return
	}

	if res.Error != nil {
		err = fmt.Errorf(res.Error.Message)
		return
	}

	err = json.Unmarshal(res.Result, &balance)
	return
}

func (w *WebSocket) GetAsset(assetId string) (asset Asset, err error) {
	params := map[string]string{"asset": assetId}
	res, err := w.ws.Call(GetAsset, params)
	if err != nil {
		return
	}

	if res.Error != nil {
		err = fmt.Errorf(res.Error.Message)
		return
	}

	err = json.Unmarshal(res.Result, &asset)
	return
}

func (w *WebSocket) GetAssets(params GetAssetsParams) (assets []string, err error) {
	res, err := w.ws.Call(GetAssets, params)
	if err != nil {
		return
	}

	if res.Error != nil {
		err = fmt.Errorf(res.Error.Message)
		return
	}

	err = json.Unmarshal(res.Result, &assets)
	return
}

func (w *WebSocket) CountAssets() (count uint64, err error) {
	res, err := w.ws.Call(CountAssets, nil)
	if err != nil {
		return
	}

	if res.Error != nil {
		err = fmt.Errorf(res.Error.Message)
		return
	}

	err = json.Unmarshal(res.Result, &count)
	return
}

func (w *WebSocket) CountTransactions() (count uint64, err error) {
	res, err := w.ws.Call(CountTransactions, nil)
	if err != nil {
		return
	}

	if res.Error != nil {
		err = fmt.Errorf(res.Error.Message)
		return
	}

	err = json.Unmarshal(res.Result, &count)
	return
}

func (w *WebSocket) CountAccounts() (count uint64, err error) {
	res, err := w.ws.Call(CountAccounts, nil)
	if err != nil {
		return
	}

	if res.Error != nil {
		err = fmt.Errorf(res.Error.Message)
		return
	}

	err = json.Unmarshal(res.Result, &count)
	return
}

func (w *WebSocket) GetTips() (tips []string, err error) {
	res, err := w.ws.Call(GetTips, nil)
	if err != nil {
		return
	}

	if res.Error != nil {
		err = fmt.Errorf(res.Error.Message)
		return
	}

	err = json.Unmarshal(res.Result, &tips)
	return
}

func (w *WebSocket) P2PStatus() (status P2PStatusResult, err error) {
	res, err := w.ws.Call(P2PStatus, nil)
	if err != nil {
		return
	}

	if res.Error != nil {
		err = fmt.Errorf(res.Error.Message)
		return
	}

	err = json.Unmarshal(res.Result, &status)
	return
}

func (w *WebSocket) GetDAGOrder(params GetTopoheightRangeParams) (hashes []string, err error) {
	res, err := w.ws.Call(GetDAGOrder, params)
	if err != nil {
		return
	}

	if res.Error != nil {
		err = fmt.Errorf(res.Error.Message)
		return
	}

	err = json.Unmarshal(res.Result, &hashes)
	return
}

func (w *WebSocket) SubmitBlock(hexData string) (result bool, err error) {
	params := map[string]string{"block_template": hexData}
	res, err := w.ws.Call(SubmitBlock, params)
	if err != nil {
		return
	}

	if res.Error != nil {
		err = fmt.Errorf(res.Error.Message)
		return
	}

	err = json.Unmarshal(res.Result, &result)
	return
}

func (w *WebSocket) SubmitTransaction(hexData string) (result bool, err error) {
	params := map[string]string{"data": hexData}
	res, err := w.ws.Call(SubmitTransaction, params)
	if err != nil {
		return
	}

	if res.Error != nil {
		err = fmt.Errorf(res.Error.Message)
		return
	}

	err = json.Unmarshal(res.Result, &result)
	return
}

func (w *WebSocket) GetMempool() (txs []Transaction, err error) {
	res, err := w.ws.Call(GetMempool, nil)
	if err != nil {
		return
	}

	if res.Error != nil {
		err = fmt.Errorf(res.Error.Message)
		return
	}

	err = json.Unmarshal(res.Result, &txs)
	return
}

func (w *WebSocket) GetTransaction(hash string) (tx Transaction, err error) {
	params := map[string]string{"hash": hash}
	res, err := w.ws.Call(GetTransaction, params)
	if err != nil {
		return
	}

	if res.Error != nil {
		err = fmt.Errorf(res.Error.Message)
		return
	}

	err = json.Unmarshal(res.Result, &tx)
	return
}

func (w *WebSocket) GetTransactions(params GetTransactionsParams) (txs []Transaction, err error) {
	res, err := w.ws.Call(GetTransactions, params)
	if err != nil {
		return
	}

	if res.Error != nil {
		err = fmt.Errorf(res.Error.Message)
		return
	}

	err = json.Unmarshal(res.Result, &txs)
	return
}

func (w *WebSocket) GetBlocksRangeByTopoheight(params GetTopoheightRangeParams) (blocks []Block, err error) {
	res, err := w.ws.Call(GetBlocksRangeByTopoheight, params)
	if err != nil {
		return
	}

	if res.Error != nil {
		err = fmt.Errorf(res.Error.Message)
		return
	}

	err = json.Unmarshal(res.Result, &blocks)
	return
}

func (w *WebSocket) GetBlocksRangeByHeight(params GetHeightRangeParams) (blocks []Block, err error) {
	res, err := w.ws.Call(GetBlocksRangeByHeight, params)
	if err != nil {
		return
	}

	if res.Error != nil {
		err = fmt.Errorf(res.Error.Message)
		return
	}

	err = json.Unmarshal(res.Result, &blocks)
	return
}

func (w *WebSocket) GetAccounts(params GetAccountsParams) (addresses []string, err error) {
	res, err := w.ws.Call(GetAccounts, params)
	if err != nil {
		return
	}

	if res.Error != nil {
		err = fmt.Errorf(res.Error.Message)
		return
	}

	err = json.Unmarshal(res.Result, &addresses)
	return
}

func (w *WebSocket) GetAccountHistory(addr string) (history AccountHistory, err error) {
	params := map[string]string{"address": addr}
	res, err := w.ws.Call(GetAccountHistory, params)
	if err != nil {
		return
	}

	if res.Error != nil {
		err = fmt.Errorf(res.Error.Message)
		return
	}

	err = json.Unmarshal(res.Result, &history)
	return
}

func (w *WebSocket) GetAccountAssets(addr string) (assets []string, err error) {
	params := map[string]string{"address": addr}
	res, err := w.ws.Call(GetAccountAssets, params)
	if err != nil {
		return
	}

	if res.Error != nil {
		err = fmt.Errorf(res.Error.Message)
		return
	}

	err = json.Unmarshal(res.Result, &assets)
	return
}

func (w *WebSocket) GetPeers() (peers []Peer, err error) {
	res, err := w.ws.Call(GetPeers, nil)
	if err != nil {
		return
	}

	if res.Error != nil {
		err = fmt.Errorf(res.Error.Message)
		return
	}

	err = json.Unmarshal(res.Result, &peers)
	return
}

func (w *WebSocket) GetDevFeeThresholds() (fees []Fee, err error) {
	res, err := w.ws.Call(GetDevFeeThresholds, nil)
	if err != nil {
		return
	}

	if res.Error != nil {
		err = fmt.Errorf(res.Error.Message)
		return
	}

	err = json.Unmarshal(res.Result, &fees)
	return
}

func (w *WebSocket) GetSizeOnDisk() (sizeOnDisk SizeOnDisk, err error) {
	res, err := w.ws.Call(GetSizeOnDisk, nil)
	if err != nil {
		return
	}

	if res.Error != nil {
		err = fmt.Errorf(res.Error.Message)
		return
	}

	err = json.Unmarshal(res.Result, &sizeOnDisk)
	return
}

func (w *WebSocket) IsTxExecutedInBlock(params IsTxExecutedInBlockParams) (executed bool, err error) {
	res, err := w.ws.Call(IsTxExecutedInBlock, params)
	if err != nil {
		return
	}

	if res.Error != nil {
		err = fmt.Errorf(res.Error.Message)
		return
	}

	err = json.Unmarshal(res.Result, &executed)
	return
}
