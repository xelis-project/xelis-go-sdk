package daemon

import (
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
		err := lib.JsonFormatResponse(res, nil, &result)
		onData(result, err)
	})
}

func (w *WebSocket) TransactionAddedInMempoolFunc(onData func(Transaction, error)) error {
	return w.WS.ListenEventFunc(TransactionAddedInMempool, func(res lib.RPCResponse) {
		var result Transaction
		err := lib.JsonFormatResponse(res, nil, &result)
		onData(result, err)
	})
}

func (w *WebSocket) BlockOrderedFunc(onData func(Block, error)) error {
	return w.WS.ListenEventFunc(BlockOrdered, func(res lib.RPCResponse) {
		var result Block
		err := lib.JsonFormatResponse(res, nil, &result)
		onData(result, err)
	})
}

func (w *WebSocket) TransactionExecutedFunc(onData func(TransactionExecutedResult, error)) error {
	return w.WS.ListenEventFunc(TransactionExecuted, func(res lib.RPCResponse) {
		var result TransactionExecutedResult
		err := lib.JsonFormatResponse(res, nil, &result)
		onData(result, err)
	})
}

func (w *WebSocket) PeerConnectedFunc(onData func(Peer, error)) error {
	return w.WS.ListenEventFunc(PeerConnected, func(res lib.RPCResponse) {
		var result Peer
		err := lib.JsonFormatResponse(res, nil, &result)
		onData(result, err)
	})
}

func (w *WebSocket) PeerDisconnectedFunc(onData func(uint64, error)) error {
	return w.WS.ListenEventFunc(PeerDisconnected, func(res lib.RPCResponse) {
		var id uint64
		err := lib.JsonFormatResponse(res, nil, &id)
		onData(id, err)
	})
}

func (w *WebSocket) PeerStateUpdatedFunc(onData func(Peer, error)) error {
	return w.WS.ListenEventFunc(PeerStateUpdated, func(res lib.RPCResponse) {
		var result Peer
		err := lib.JsonFormatResponse(res, nil, &result)
		onData(result, err)
	})
}

func (w *WebSocket) GetVersion() (version string, err error) {
	res, err := w.WS.Call(w.Prefix+GetVersion, nil)
	err = lib.JsonFormatResponse(res, err, &version)
	return
}

func (w *WebSocket) GetInfo() (result GetInfoResult, err error) {
	res, err := w.WS.Call(w.Prefix+GetInfo, nil)
	err = lib.JsonFormatResponse(res, err, &result)
	return
}

func (w *WebSocket) GetHeight() (height uint64, err error) {
	res, err := w.WS.Call(w.Prefix+GetHeight, nil)
	err = lib.JsonFormatResponse(res, err, &height)
	return
}

func (w *WebSocket) GetTopoheight() (topoheight uint64, err error) {
	res, err := w.WS.Call(w.Prefix+GetTopoHeight, nil)
	err = lib.JsonFormatResponse(res, err, &topoheight)
	return
}

func (w *WebSocket) GetStableheight() (stableheight uint64, err error) {
	res, err := w.WS.Call(w.Prefix+GetStableHeight, nil)
	err = lib.JsonFormatResponse(res, err, &stableheight)
	return
}

func (w *WebSocket) GetBlockTemplate(addr string) (result GetBlockTemplateResult, err error) {
	params := map[string]string{"address": addr}
	res, err := w.WS.Call(w.Prefix+GetBlockTemplate, params)
	err = lib.JsonFormatResponse(res, err, &result)
	return
}

func (w *WebSocket) GetBlockAtTopoheight(params GetBlockAtTopoheightParams) (block Block, err error) {
	res, err := w.WS.Call(w.Prefix+GetBlockAtTopoheight, params)
	err = lib.JsonFormatResponse(res, err, &block)
	return
}

func (w *WebSocket) GetBlocksAtHeight(params GetBlockAtTopoheightParams) (blocks []Block, err error) {
	res, err := w.WS.Call(w.Prefix+GetBlocksAtHeight, params)
	err = lib.JsonFormatResponse(res, err, &blocks)
	return
}

func (w *WebSocket) GetBlockByHash(params GetBlockByHashParams) (block Block, err error) {
	res, err := w.WS.Call(w.Prefix+GetBlockByHash, params)
	err = lib.JsonFormatResponse(res, err, &block)
	return
}

func (w *WebSocket) GetTopBlock(params GetTopBlockParams) (block Block, err error) {
	res, err := w.WS.Call(w.Prefix+GetTopBlock, params)
	err = lib.JsonFormatResponse(res, err, &block)
	return
}

func (w *WebSocket) GetNonce(addr string) (nonce GetNonceResult, err error) {
	params := map[string]string{"address": addr}
	res, err := w.WS.Call(w.Prefix+GetNonce, params)
	err = lib.JsonFormatResponse(res, err, &nonce)
	return
}

func (w *WebSocket) GetNonceAtTopoheight(params GetNonceAtTopoheightParams) (nonce GetNonceResult, err error) {
	res, err := w.WS.Call(w.Prefix+GetNonceAtTopoheight, params)
	err = lib.JsonFormatResponse(res, err, &nonce)
	return
}

func (w *WebSocket) HasNonce(addr string) (hasNonce bool, err error) {
	params := map[string]string{"address": addr}
	res, err := w.WS.Call(w.Prefix+HasNonce, params)
	err = lib.JsonFormatResponse(res, err, &hasNonce)
	return
}

func (w *WebSocket) GetBalance(params GetBalanceParams) (balance GetBalanceResult, err error) {
	res, err := w.WS.Call(w.Prefix+GetBalance, params)
	err = lib.JsonFormatResponse(res, err, &balance)
	return
}

func (w *WebSocket) HasBalance(params GetBalanceParams) (hasBalance bool, err error) {
	res, err := w.WS.Call(w.Prefix+HasBalance, params)
	err = lib.JsonFormatResponse(res, err, &hasBalance)
	return
}

func (w *WebSocket) GetBalanceAtTopoheight(params GetBalanceAtTopoheightParams) (balance VersionedBalance, err error) {
	res, err := w.WS.Call(w.Prefix+GetBalanceAtTopoheight, params)
	err = lib.JsonFormatResponse(res, err, &balance)
	return
}

func (w *WebSocket) GetAsset(assetId string) (asset Asset, err error) {
	params := map[string]string{"asset": assetId}
	res, err := w.WS.Call(w.Prefix+GetAsset, params)
	err = lib.JsonFormatResponse(res, err, &asset)
	return
}

func (w *WebSocket) GetAssets(params GetAssetsParams) (assets []AssetWithData, err error) {
	res, err := w.WS.Call(w.Prefix+GetAssets, params)
	err = lib.JsonFormatResponse(res, err, &assets)
	return
}

func (w *WebSocket) CountAssets() (count uint64, err error) {
	res, err := w.WS.Call(w.Prefix+CountAssets, nil)
	err = lib.JsonFormatResponse(res, err, &count)
	return
}

func (w *WebSocket) CountTransactions() (count uint64, err error) {
	res, err := w.WS.Call(w.Prefix+CountTransactions, nil)
	err = lib.JsonFormatResponse(res, err, &count)
	return
}

func (w *WebSocket) CountAccounts() (count uint64, err error) {
	res, err := w.WS.Call(w.Prefix+CountAccounts, nil)
	err = lib.JsonFormatResponse(res, err, &count)
	return
}

func (w *WebSocket) GetTips() (tips []string, err error) {
	res, err := w.WS.Call(w.Prefix+GetTips, nil)
	err = lib.JsonFormatResponse(res, err, &tips)
	return
}

func (w *WebSocket) P2PStatus() (status P2PStatusResult, err error) {
	res, err := w.WS.Call(w.Prefix+P2PStatus, nil)
	err = lib.JsonFormatResponse(res, err, &status)
	return
}

func (w *WebSocket) GetDAGOrder(params GetTopoheightRangeParams) (hashes []string, err error) {
	res, err := w.WS.Call(w.Prefix+GetDAGOrder, params)
	err = lib.JsonFormatResponse(res, err, &hashes)
	return
}

func (w *WebSocket) SubmitBlock(hexData string) (result bool, err error) {
	params := map[string]string{"block_template": hexData}
	res, err := w.WS.Call(w.Prefix+SubmitBlock, params)
	err = lib.JsonFormatResponse(res, err, &result)
	return
}

func (w *WebSocket) SubmitTransaction(hexData string) (result bool, err error) {
	params := map[string]string{"data": hexData}
	res, err := w.WS.Call(w.Prefix+SubmitTransaction, params)
	err = lib.JsonFormatResponse(res, err, &hexData)
	return
}

func (w *WebSocket) GetMempool() (txs []Transaction, err error) {
	res, err := w.WS.Call(w.Prefix+GetMempool, nil)
	err = lib.JsonFormatResponse(res, err, &txs)
	return
}

func (w *WebSocket) GetTransaction(hash string) (tx Transaction, err error) {
	params := map[string]string{"hash": hash}
	res, err := w.WS.Call(w.Prefix+GetTransaction, params)
	err = lib.JsonFormatResponse(res, err, &tx)
	return
}

func (w *WebSocket) GetTransactions(params GetTransactionsParams) (txs []Transaction, err error) {
	res, err := w.WS.Call(w.Prefix+GetTransactions, params)
	err = lib.JsonFormatResponse(res, err, &txs)
	return
}

func (w *WebSocket) GetBlocksRangeByTopoheight(params GetTopoheightRangeParams) (blocks []Block, err error) {
	res, err := w.WS.Call(w.Prefix+GetBlocksRangeByTopoheight, params)
	err = lib.JsonFormatResponse(res, err, &blocks)
	return
}

func (w *WebSocket) GetBlocksRangeByHeight(params GetHeightRangeParams) (blocks []Block, err error) {
	res, err := w.WS.Call(w.Prefix+GetBlocksRangeByHeight, params)
	err = lib.JsonFormatResponse(res, err, &blocks)
	return
}

func (w *WebSocket) GetAccounts(params GetAccountsParams) (addresses []string, err error) {
	res, err := w.WS.Call(w.Prefix+GetAccounts, params)
	err = lib.JsonFormatResponse(res, err, &addresses)
	return
}

func (w *WebSocket) GetAccountHistory(addr string) (history []AccountHistory, err error) {
	params := map[string]string{"address": addr}
	res, err := w.WS.Call(w.Prefix+GetAccountHistory, params)
	err = lib.JsonFormatResponse(res, err, &history)
	return
}

func (w *WebSocket) GetAccountAssets(addr string) (assets []string, err error) {
	params := map[string]string{"address": addr}
	res, err := w.WS.Call(w.Prefix+GetAccountAssets, params)
	err = lib.JsonFormatResponse(res, err, &assets)
	return
}

func (w *WebSocket) GetPeers() (result GetPeersResult, err error) {
	res, err := w.WS.Call(w.Prefix+GetPeers, nil)
	err = lib.JsonFormatResponse(res, err, &result)
	return
}

func (w *WebSocket) GetDevFeeThresholds() (fees []Fee, err error) {
	res, err := w.WS.Call(w.Prefix+GetDevFeeThresholds, nil)
	err = lib.JsonFormatResponse(res, err, &fees)
	return
}

func (w *WebSocket) GetSizeOnDisk() (sizeOnDisk SizeOnDisk, err error) {
	res, err := w.WS.Call(w.Prefix+GetSizeOnDisk, nil)
	err = lib.JsonFormatResponse(res, err, &sizeOnDisk)
	return
}

func (w *WebSocket) IsTxExecutedInBlock(params IsTxExecutedInBlockParams) (executed bool, err error) {
	res, err := w.WS.Call(w.Prefix+IsTxExecutedInBlock, params)
	err = lib.JsonFormatResponse(res, err, &executed)
	return
}
