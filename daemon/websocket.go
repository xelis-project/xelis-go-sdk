package daemon

import (
	"encoding/json"
	"fmt"

	"github.com/xelis-project/xelis-go-sdk/lib"
)

type WebSocket struct {
	Prefix string
	WS     *lib.WebSocket
}

func NewWebSocket(endpoint string) (*WebSocket, error) {
	ws, err := lib.NewWebSocket(endpoint, nil)
	if err != nil {
		return nil, err
	}

	return &WebSocket{
		WS: ws,
	}, nil
}

func (w *WebSocket) Close() error {
	return w.WS.Close()
}

func (w *WebSocket) CloseEvent(event string) error {
	return w.WS.CloseEvent(event)
}

func (w *WebSocket) NewBlockFunc(onData func(Block, error)) error {
	return w.WS.ListenEventFunc(NewBlock, func(res lib.RPCResponse) {
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
	return w.WS.ListenEventFunc(TransactionAddedInMempool, func(res lib.RPCResponse) {
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
	return w.WS.ListenEventFunc(BlockOrdered, func(res lib.RPCResponse) {
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
	return w.WS.ListenEventFunc(TransactionExecuted, func(res lib.RPCResponse) {
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
	return w.WS.ListenEventFunc(PeerConnected, func(res lib.RPCResponse) {
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
	return w.WS.ListenEventFunc(PeerDisconnected, func(res lib.RPCResponse) {
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
	return w.WS.ListenEventFunc(PeerStateUpdated, func(res lib.RPCResponse) {
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
	res, err := w.WS.Call(w.Prefix+GetVersion, nil)
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
	res, err := w.WS.Call(w.Prefix+GetInfo, nil)
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
	res, err := w.WS.Call(w.Prefix+GetHeight, nil)
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
	res, err := w.WS.Call(w.Prefix+GetTopoHeight, nil)
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
	res, err := w.WS.Call(w.Prefix+GetStableHeight, nil)
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
	res, err := w.WS.Call(w.Prefix+GetBlockTemplate, params)
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
	res, err := w.WS.Call(w.Prefix+GetBlockAtTopoheight, params)
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
	res, err := w.WS.Call(w.Prefix+GetBlocksAtHeight, params)
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
	res, err := w.WS.Call(w.Prefix+GetBlockByHash, params)
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
	res, err := w.WS.Call(w.Prefix+GetTopBlock, params)
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
	res, err := w.WS.Call(w.Prefix+GetNonce, params)
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
	res, err := w.WS.Call(w.Prefix+HasNonce, params)
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
	res, err := w.WS.Call(w.Prefix+GetBalance, params)
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
	res, err := w.WS.Call(w.Prefix+HasBalance, params)
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
	res, err := w.WS.Call(w.Prefix+GetBalanceAtTopoheight, params)
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
	res, err := w.WS.Call(w.Prefix+GetAsset, params)
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
	res, err := w.WS.Call(w.Prefix+GetAssets, params)
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
	res, err := w.WS.Call(w.Prefix+CountAssets, nil)
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
	res, err := w.WS.Call(w.Prefix+CountTransactions, nil)
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
	res, err := w.WS.Call(w.Prefix+CountAccounts, nil)
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
	res, err := w.WS.Call(w.Prefix+GetTips, nil)
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
	res, err := w.WS.Call(w.Prefix+P2PStatus, nil)
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
	res, err := w.WS.Call(w.Prefix+GetDAGOrder, params)
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
	res, err := w.WS.Call(w.Prefix+SubmitBlock, params)
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
	res, err := w.WS.Call(w.Prefix+SubmitTransaction, params)
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
	res, err := w.WS.Call(w.Prefix+GetMempool, nil)
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
	res, err := w.WS.Call(w.Prefix+GetTransaction, params)
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
	res, err := w.WS.Call(w.Prefix+GetTransactions, params)
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
	res, err := w.WS.Call(w.Prefix+GetBlocksRangeByTopoheight, params)
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
	res, err := w.WS.Call(w.Prefix+GetBlocksRangeByHeight, params)
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
	res, err := w.WS.Call(w.Prefix+GetAccounts, params)
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
	res, err := w.WS.Call(w.Prefix+GetAccountHistory, params)
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
	res, err := w.WS.Call(w.Prefix+GetAccountAssets, params)
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
	res, err := w.WS.Call(w.Prefix+GetPeers, nil)
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
	res, err := w.WS.Call(w.Prefix+GetDevFeeThresholds, nil)
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
	res, err := w.WS.Call(w.Prefix+GetSizeOnDisk, nil)
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
	res, err := w.WS.Call(w.Prefix+IsTxExecutedInBlock, params)
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
