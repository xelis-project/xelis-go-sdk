package daemon

import (
	"github.com/xelis-project/xelis-go-sdk/daemon/events"
	"github.com/xelis-project/xelis-go-sdk/daemon/methods"
	"github.com/xelis-project/xelis-go-sdk/rpc"
)

type WebSocket struct {
	Prefix string
	WS     *rpc.WebSocket
}

func NewWebSocket(endpoint string) (*WebSocket, error) {
	ws, err := rpc.NewWebSocket(endpoint, nil)
	if err != nil {
		return nil, err
	}

	return &WebSocket{
		WS: ws,
	}, nil
}

func (w *WebSocket) BatchCall(requests []rpc.RPCRequest, result []interface{}) (res []rpc.RPCResponse, errs []error) {
	return w.WS.BatchCall(requests, result)
}

func (w *WebSocket) Close() error {
	return w.WS.Close()
}

func (w *WebSocket) CloseEvent(event string) error {
	return w.WS.CloseEvent(event)
}

func (w *WebSocket) ConnectionErr() chan error {
	return w.WS.ConnectionErr
}

func (w *WebSocket) NewBlockChannel() (chan Block, chan error, error) {
	chanResult := make(chan Block)
	chanErr := make(chan error)

	err := w.WS.ListenEventFunc(w.Prefix+events.NewBlock, func(res rpc.RPCResponse) {
		var result Block
		err := rpc.ParseResponseResult(res, &result)
		if err != nil {
			chanErr <- err
		} else {
			chanResult <- result
		}
	})

	return chanResult, chanErr, err
}

func (w *WebSocket) NewBlockFunc(onData func(Block, error)) error {
	return w.WS.ListenEventFunc(w.Prefix+events.NewBlock, func(res rpc.RPCResponse) {
		var result Block
		err := rpc.ParseResponseResult(res, &result)
		onData(result, err)
	})
}

func (w *WebSocket) TransactionAddedInMempoolChannel() (chan Transaction, chan error, error) {
	chanResult := make(chan Transaction)
	chanErr := make(chan error)

	err := w.WS.ListenEventFunc(w.Prefix+events.TransactionAddedInMempool, func(res rpc.RPCResponse) {
		var result Transaction
		err := rpc.ParseResponseResult(res, &result)
		if err != nil {
			chanErr <- err
		} else {
			chanResult <- result
		}
	})

	return chanResult, chanErr, err
}

func (w *WebSocket) TransactionAddedInMempoolFunc(onData func(Transaction, error)) error {
	return w.WS.ListenEventFunc(w.Prefix+events.TransactionAddedInMempool, func(res rpc.RPCResponse) {
		var result Transaction
		err := rpc.ParseResponseResult(res, &result)
		onData(result, err)
	})
}

func (w *WebSocket) BlockOrderedChannel() (chan BlockOrderedEvent, chan error, error) {
	chanResult := make(chan BlockOrderedEvent)
	chanErr := make(chan error)

	err := w.WS.ListenEventFunc(w.Prefix+events.BlockOrdered, func(res rpc.RPCResponse) {
		var result BlockOrderedEvent
		err := rpc.ParseResponseResult(res, &result)
		if err != nil {
			chanErr <- err
		} else {
			chanResult <- result
		}
	})

	return chanResult, chanErr, err
}

func (w *WebSocket) BlockOrderedFunc(onData func(BlockOrderedEvent, error)) error {
	return w.WS.ListenEventFunc(w.Prefix+events.BlockOrdered, func(res rpc.RPCResponse) {
		var result BlockOrderedEvent
		err := rpc.ParseResponseResult(res, &result)
		onData(result, err)
	})
}

func (w *WebSocket) TransactionExecutedChannel() (chan TransactionExecutedEvent, chan error, error) {
	chanResult := make(chan TransactionExecutedEvent)
	chanErr := make(chan error)

	err := w.WS.ListenEventFunc(w.Prefix+events.TransactionExecuted, func(res rpc.RPCResponse) {
		var result TransactionExecutedEvent
		err := rpc.ParseResponseResult(res, &result)
		if err != nil {
			chanErr <- err
		} else {
			chanResult <- result
		}
	})

	return chanResult, chanErr, err
}

func (w *WebSocket) TransactionExecutedFunc(onData func(TransactionExecutedEvent, error)) error {
	return w.WS.ListenEventFunc(w.Prefix+events.TransactionExecuted, func(res rpc.RPCResponse) {
		var result TransactionExecutedEvent
		err := rpc.ParseResponseResult(res, &result)
		onData(result, err)
	})
}

func (w *WebSocket) PeerConnectedChannel() (chan Peer, chan error, error) {
	chanResult := make(chan Peer)
	chanErr := make(chan error)

	err := w.WS.ListenEventFunc(w.Prefix+events.PeerConnected, func(res rpc.RPCResponse) {
		var result Peer
		err := rpc.ParseResponseResult(res, &result)
		if err != nil {
			chanErr <- err
		} else {
			chanResult <- result
		}
	})

	return chanResult, chanErr, err
}

func (w *WebSocket) PeerConnectedFunc(onData func(Peer, error)) error {
	return w.WS.ListenEventFunc(w.Prefix+events.PeerConnected, func(res rpc.RPCResponse) {
		var result Peer
		err := rpc.ParseResponseResult(res, &result)
		onData(result, err)
	})
}

func (w *WebSocket) PeerDisconnectedChannel() (chan Peer, chan error, error) {
	chanResult := make(chan Peer)
	chanErr := make(chan error)

	err := w.WS.ListenEventFunc(w.Prefix+events.PeerDisconnected, func(res rpc.RPCResponse) {
		var result Peer
		err := rpc.ParseResponseResult(res, &result)
		if err != nil {
			chanErr <- err
		} else {
			chanResult <- result
		}
	})

	return chanResult, chanErr, err
}

func (w *WebSocket) PeerDisconnectedFunc(onData func(Peer, error)) error {
	return w.WS.ListenEventFunc(w.Prefix+events.PeerDisconnected, func(res rpc.RPCResponse) {
		var result Peer
		err := rpc.ParseResponseResult(res, &result)
		onData(result, err)
	})
}

func (w *WebSocket) PeerStateUpdatedChannel() (chan Peer, chan error, error) {
	chanResult := make(chan Peer)
	chanErr := make(chan error)

	err := w.WS.ListenEventFunc(w.Prefix+events.PeerStateUpdated, func(res rpc.RPCResponse) {
		var result Peer
		err := rpc.ParseResponseResult(res, &result)
		if err != nil {
			chanErr <- err
		} else {
			chanResult <- result
		}
	})

	return chanResult, chanErr, err
}

func (w *WebSocket) PeerStateUpdatedFunc(onData func(Peer, error)) error {
	return w.WS.ListenEventFunc(w.Prefix+events.PeerStateUpdated, func(res rpc.RPCResponse) {
		var result Peer
		err := rpc.ParseResponseResult(res, &result)
		onData(result, err)
	})
}

func (w *WebSocket) BlockOrphanedChannel() (chan BlockOrphanedEvent, chan error, error) {
	chanResult := make(chan BlockOrphanedEvent)
	chanErr := make(chan error)

	err := w.WS.ListenEventFunc(w.Prefix+events.BlockOrphaned, func(res rpc.RPCResponse) {
		var result BlockOrphanedEvent
		err := rpc.ParseResponseResult(res, &result)
		if err != nil {
			chanErr <- err
		} else {
			chanResult <- result
		}
	})

	return chanResult, chanErr, err
}

func (w *WebSocket) BlockOrphanedFunc(onData func(BlockOrphanedEvent, error)) error {
	return w.WS.ListenEventFunc(w.Prefix+events.BlockOrphaned, func(res rpc.RPCResponse) {
		var result BlockOrphanedEvent
		err := rpc.ParseResponseResult(res, &result)
		onData(result, err)
	})
}

func (w *WebSocket) StableHeightChangedChannel() (chan StableHeightChangedEvent, chan error, error) {
	chanResult := make(chan StableHeightChangedEvent)
	chanErr := make(chan error)

	err := w.WS.ListenEventFunc(w.Prefix+events.StableHeightChanged, func(res rpc.RPCResponse) {
		var result StableHeightChangedEvent
		err := rpc.ParseResponseResult(res, &result)
		if err != nil {
			chanErr <- err
		} else {
			chanResult <- result
		}
	})

	return chanResult, chanErr, err
}

func (w *WebSocket) StableHeightChangedFunc(onData func(StableHeightChangedEvent, error)) error {
	return w.WS.ListenEventFunc(w.Prefix+events.StableHeightChanged, func(res rpc.RPCResponse) {
		var result StableHeightChangedEvent
		err := rpc.ParseResponseResult(res, &result)
		onData(result, err)
	})
}

func (w *WebSocket) StableTopoheightChangedChannel() (chan StableTopoheightChangedEvent, chan error, error) {
	chanResult := make(chan StableTopoheightChangedEvent)
	chanErr := make(chan error)

	err := w.WS.ListenEventFunc(w.Prefix+events.StableTopoheightChanged, func(res rpc.RPCResponse) {
		var result StableTopoheightChangedEvent
		err := rpc.ParseResponseResult(res, &result)
		if err != nil {
			chanErr <- err
		} else {
			chanResult <- result
		}
	})

	return chanResult, chanErr, err
}

func (w *WebSocket) StableTopoheightChangedFunc(onData func(StableTopoheightChangedEvent, error)) error {
	return w.WS.ListenEventFunc(w.Prefix+events.StableTopoheightChanged, func(res rpc.RPCResponse) {
		var result StableTopoheightChangedEvent
		err := rpc.ParseResponseResult(res, &result)
		onData(result, err)
	})
}

func (w *WebSocket) PeerPeerListUpdatedChannel() (chan PeerPeerListUpdatedEvent, chan error, error) {
	chanResult := make(chan PeerPeerListUpdatedEvent)
	chanErr := make(chan error)

	err := w.WS.ListenEventFunc(w.Prefix+events.PeerPeerListUpdated, func(res rpc.RPCResponse) {
		var result PeerPeerListUpdatedEvent
		err := rpc.ParseResponseResult(res, &result)
		if err != nil {
			chanErr <- err
		} else {
			chanResult <- result
		}
	})

	return chanResult, chanErr, err
}

func (w *WebSocket) PeerPeerListUpdatedFunc(onData func(PeerPeerListUpdatedEvent, error)) error {
	return w.WS.ListenEventFunc(w.Prefix+events.PeerPeerListUpdated, func(res rpc.RPCResponse) {
		var result PeerPeerListUpdatedEvent
		err := rpc.ParseResponseResult(res, &result)
		onData(result, err)
	})
}

func (w *WebSocket) PeerPeerDisconnectedChannel() (chan PeerPeerDisconnectedEvent, chan error, error) {
	chanResult := make(chan PeerPeerDisconnectedEvent)
	chanErr := make(chan error)

	err := w.WS.ListenEventFunc(w.Prefix+events.PeerPeerDisconnected, func(res rpc.RPCResponse) {
		var result PeerPeerDisconnectedEvent
		err := rpc.ParseResponseResult(res, &result)
		if err != nil {
			chanErr <- err
		} else {
			chanResult <- result
		}
	})

	return chanResult, chanErr, err
}

func (w *WebSocket) PeerPeerDisconnectedFunc(onData func(PeerPeerDisconnectedEvent, error)) error {
	return w.WS.ListenEventFunc(w.Prefix+events.PeerPeerDisconnected, func(res rpc.RPCResponse) {
		var result PeerPeerDisconnectedEvent
		err := rpc.ParseResponseResult(res, &result)
		onData(result, err)
	})
}

func (w *WebSocket) TransactionOrphanedChannel() (chan Transaction, chan error, error) {
	chanResult := make(chan Transaction)
	chanErr := make(chan error)

	err := w.WS.ListenEventFunc(w.Prefix+events.TransactionOrphaned, func(res rpc.RPCResponse) {
		var result Transaction
		err := rpc.ParseResponseResult(res, &result)
		if err != nil {
			chanErr <- err
		} else {
			chanResult <- result
		}
	})

	return chanResult, chanErr, err
}

func (w *WebSocket) TransactionOrphanedFunc(onData func(Transaction, error)) error {
	return w.WS.ListenEventFunc(w.Prefix+events.TransactionOrphaned, func(res rpc.RPCResponse) {
		var result Transaction
		err := rpc.ParseResponseResult(res, &result)
		onData(result, err)
	})
}

func (w *WebSocket) DeployContractChannel() (chan NewContractEvent, chan error, error) {
	chanResult := make(chan NewContractEvent)
	chanErr := make(chan error)

	err := w.WS.ListenEventFunc(w.Prefix+events.DeployContract, func(res rpc.RPCResponse) {
		var result NewContractEvent
		err := rpc.ParseResponseResult(res, &result)
		if err != nil {
			chanErr <- err
		} else {
			chanResult <- result
		}
	})

	return chanResult, chanErr, err
}

func (w *WebSocket) DeployContractFunc(onData func(NewContractEvent, error)) error {
	return w.WS.ListenEventFunc(w.Prefix+events.DeployContract, func(res rpc.RPCResponse) {
		var result NewContractEvent
		err := rpc.ParseResponseResult(res, &result)
		onData(result, err)
	})
}

func (w *WebSocket) InvokeContractChannel(params InvokeContractEventParams) (chan InvokeContractEvent, chan error, error) {
	chanResult := make(chan InvokeContractEvent)
	chanErr := make(chan error)

	event := rpc.EventParamsWrap(w.Prefix+events.InvokeContract, params)
	err := w.WS.ListenEventFunc(event, func(res rpc.RPCResponse) {
		var result InvokeContractEvent
		err := rpc.ParseResponseResult(res, &result)
		if err != nil {
			chanErr <- err
		} else {
			chanResult <- result
		}
	})

	return chanResult, chanErr, err
}

func (w *WebSocket) InvokeContractFunc(params InvokeContractEventParams, onData func(InvokeContractEvent, error)) error {
	event := rpc.EventParamsWrap(w.Prefix+events.InvokeContract, params)
	return w.WS.ListenEventFunc(event, func(res rpc.RPCResponse) {
		var result InvokeContractEvent
		err := rpc.ParseResponseResult(res, &result)
		onData(result, err)
	})
}

func (w *WebSocket) NewBlockTemplateFunc(onData func(GetBlockTemplateResult, error)) error {
	return w.WS.ListenEventFunc(w.Prefix+events.NewBlockTemplate, func(res rpc.RPCResponse) {
		var result GetBlockTemplateResult
		err := rpc.ParseResponseResult(res, &result)
		onData(result, err)
	})
}

func (w *WebSocket) NewBlockTemplateChannel() (chan GetBlockTemplateResult, chan error, error) {
	chanResult := make(chan GetBlockTemplateResult)
	chanErr := make(chan error)

	err := w.WS.ListenEventFunc(w.Prefix+events.NewBlockTemplate, func(res rpc.RPCResponse) {
		var result GetBlockTemplateResult
		err := rpc.ParseResponseResult(res, &result)
		if err != nil {
			chanErr <- err
		} else {
			chanResult <- result
		}
	})

	return chanResult, chanErr, err
}

func (w *WebSocket) GetVersion() (version string, err error) {
	_, err = w.WS.Call(w.Prefix+methods.GetVersion, nil, &version)
	return
}

func (w *WebSocket) GetInfo() (result GetInfoResult, err error) {
	_, err = w.WS.Call(w.Prefix+methods.GetInfo, nil, &result)
	return
}

func (w *WebSocket) GetHeight() (height uint64, err error) {
	_, err = w.WS.Call(w.Prefix+methods.GetHeight, nil, &height)
	return
}

func (w *WebSocket) GetTopoheight() (topoheight uint64, err error) {
	_, err = w.WS.Call(w.Prefix+methods.GetTopoheight, nil, &topoheight)
	return
}

func (w *WebSocket) GetStableHeight() (stableheight uint64, err error) {
	_, err = w.WS.Call(w.Prefix+methods.GetStableHeight, nil, &stableheight)
	return
}

func (w *WebSocket) GetStableTopoheight() (topoheight uint64, err error) {
	_, err = w.WS.Call(w.Prefix+methods.GetStableTopoheight, nil, &topoheight)
	return
}

func (w *WebSocket) GetStableBalance(params GetBalanceParams) (result GetStableBalanceResult, err error) {
	_, err = w.WS.Call(w.Prefix+methods.GetStableBalance, nil, &result)
	return
}

func (w *WebSocket) GetBlockTemplate(addr string) (result GetBlockTemplateResult, err error) {
	params := map[string]string{"address": addr}
	_, err = w.WS.Call(w.Prefix+methods.GetBlockTemplate, params, &result)
	return
}

func (w *WebSocket) GetBlockAtTopoheight(params GetBlockAtTopoheightParams) (block Block, err error) {
	_, err = w.WS.Call(w.Prefix+methods.GetBlockAtTopoheight, params, &block)
	return
}

func (w *WebSocket) GetBlocksAtHeight(params GetBlockAtTopoheightParams) (blocks []Block, err error) {
	_, err = w.WS.Call(w.Prefix+methods.GetBlocksAtHeight, params, &blocks)
	return
}

func (w *WebSocket) GetBlockByHash(params GetBlockByHashParams) (block Block, err error) {
	_, err = w.WS.Call(w.Prefix+methods.GetBlockByHash, params, &block)
	return
}

func (w *WebSocket) GetTopBlock(params GetTopBlockParams) (block Block, err error) {
	_, err = w.WS.Call(w.Prefix+methods.GetTopBlock, params, &block)
	return
}

func (w *WebSocket) GetNonce(addr string) (nonce GetNonceResult, err error) {
	params := map[string]string{"address": addr}
	_, err = w.WS.Call(w.Prefix+methods.GetNonce, params, &nonce)
	return
}

func (w *WebSocket) GetNonceAtTopoheight(params GetNonceAtTopoheightParams) (nonce VersionedNonce, err error) {
	_, err = w.WS.Call(w.Prefix+methods.GetNonceAtTopoheight, params, &nonce)
	return
}

func (w *WebSocket) HasNonce(addr string) (hasNonce bool, err error) {
	params := map[string]string{"address": addr}
	_, err = w.WS.Call(w.Prefix+methods.HasNonce, params, &hasNonce)
	return
}

func (w *WebSocket) GetBalance(params GetBalanceParams) (balance GetBalanceResult, err error) {
	_, err = w.WS.Call(w.Prefix+methods.GetBalance, params, &balance)
	return
}

func (w *WebSocket) HasBalance(params GetBalanceParams) (hasBalance bool, err error) {
	_, err = w.WS.Call(w.Prefix+methods.HasBalance, params, &hasBalance)
	return
}

func (w *WebSocket) GetBalanceAtTopoheight(params GetBalanceAtTopoheightParams) (balance VersionedBalance, err error) {
	_, err = w.WS.Call(w.Prefix+methods.GetBalanceAtTopoheight, params, &balance)
	return
}

func (w *WebSocket) GetAsset(assetId string) (asset AssetData, err error) {
	params := map[string]string{"asset": assetId}
	_, err = w.WS.Call(w.Prefix+methods.GetAsset, params, &asset)
	return
}

func (w *WebSocket) GetAssets(params GetAssetsParams) (assets []AssetData, err error) {
	_, err = w.WS.Call(w.Prefix+methods.GetAssets, params, &assets)
	return
}

func (w *WebSocket) CountAssets() (count uint64, err error) {
	_, err = w.WS.Call(w.Prefix+methods.CountAssets, nil, &count)
	return
}

func (w *WebSocket) CountTransactions() (count uint64, err error) {
	_, err = w.WS.Call(w.Prefix+methods.CountTransactions, nil, &count)
	return
}

func (w *WebSocket) CountAccounts() (count uint64, err error) {
	_, err = w.WS.Call(w.Prefix+methods.CountAccounts, nil, &count)
	return
}

func (w *WebSocket) GetTips() (tips []string, err error) {
	_, err = w.WS.Call(w.Prefix+methods.GetTips, nil, &tips)
	return
}

func (w *WebSocket) P2PStatus() (status P2PStatusResult, err error) {
	_, err = w.WS.Call(w.Prefix+methods.P2PStatus, nil, &status)
	return
}

func (w *WebSocket) GetDAGOrder(params GetTopoheightRangeParams) (hashes []string, err error) {
	_, err = w.WS.Call(w.Prefix+methods.GetDAGOrder, params, &hashes)
	return
}

func (w *WebSocket) SubmitBlock(params SubmitBlockParams) (result bool, err error) {
	_, err = w.WS.Call(w.Prefix+methods.SubmitBlock, params, &result)
	return
}

func (w *WebSocket) SubmitTransaction(hexData string) (result bool, err error) {
	params := map[string]string{"data": hexData}
	_, err = w.WS.Call(w.Prefix+methods.SubmitTransaction, params, &hexData)
	return
}

func (w *WebSocket) GetMempool() (result GetMempoolResult, err error) {
	_, err = w.WS.Call(w.Prefix+methods.GetMempool, nil, &result)
	return
}

func (w *WebSocket) GetMempoolCache(params GetMempoolCacheParams) (result GetMempoolCacheResult, err error) {
	_, err = w.WS.Call(w.Prefix+methods.GetMempool, params, &result)
	return
}

func (w *WebSocket) GetTransaction(hash string) (tx Transaction, err error) {
	params := map[string]string{"hash": hash}
	_, err = w.WS.Call(w.Prefix+methods.GetTransaction, params, &tx)
	return
}

func (w *WebSocket) GetTransactions(params GetTransactionsParams) (txs []Transaction, err error) {
	_, err = w.WS.Call(w.Prefix+methods.GetTransactions, params, &txs)
	return
}

func (w *WebSocket) GetBlocksRangeByTopoheight(params GetTopoheightRangeParams) (blocks []Block, err error) {
	_, err = w.WS.Call(w.Prefix+methods.GetBlocksRangeByTopoheight, params, &blocks)
	return
}

func (w *WebSocket) GetBlocksRangeByHeight(params GetHeightRangeParams) (blocks []Block, err error) {
	_, err = w.WS.Call(w.Prefix+methods.GetBlocksRangeByHeight, params, &blocks)
	return
}

func (w *WebSocket) GetAccounts(params GetAccountsParams) (addresses []string, err error) {
	_, err = w.WS.Call(w.Prefix+methods.GetAccounts, params, &addresses)
	return
}

func (w *WebSocket) GetAccountHistory(addr string) (history []AccountHistory, err error) {
	params := map[string]string{"address": addr}
	_, err = w.WS.Call(w.Prefix+methods.GetAccountHistory, params, &history)
	return
}

func (w *WebSocket) GetAccountAssets(addr string) (assets []string, err error) {
	params := map[string]string{"address": addr}
	_, err = w.WS.Call(w.Prefix+methods.GetAccountAssets, params, &assets)
	return
}

func (w *WebSocket) GetPeers() (result GetPeersResult, err error) {
	_, err = w.WS.Call(w.Prefix+methods.GetPeers, nil, &result)
	return
}

func (w *WebSocket) GetDevFeeThresholds() (fees []Fee, err error) {
	_, err = w.WS.Call(w.Prefix+methods.GetDevFeeThresholds, nil, &fees)
	return
}

func (w *WebSocket) GetSizeOnDisk() (sizeOnDisk SizeOnDisk, err error) {
	_, err = w.WS.Call(w.Prefix+methods.GetSizeOnDisk, nil, &sizeOnDisk)
	return
}

func (w *WebSocket) IsTxExecutedInBlock(params IsTxExecutedInBlockParams) (executed bool, err error) {
	_, err = w.WS.Call(w.Prefix+methods.IsTxExecutedInBlock, params, &executed)
	return
}

func (w *WebSocket) GetAccountRegistrationTopoheight(addr string) (topoheight uint64, err error) {
	params := map[string]string{"address": addr}
	_, err = w.WS.Call(w.Prefix+methods.GetAccountRegistrationTopoheight, params, &topoheight)
	return
}

func (w *WebSocket) IsAccountRegistered(params IsAccountRegisteredParams) (exists bool, err error) {
	_, err = w.WS.Call(w.Prefix+methods.IsAccountRegistered, params, &exists)
	return
}

func (w *WebSocket) GetDifficulty() (result GetDifficultyResult, err error) {
	_, err = w.WS.Call(w.Prefix+methods.GetDifficulty, nil, &result)
	return
}

func (w *WebSocket) ValidateAddress(params ValidateAddressParams) (result ValidateAddressResult, err error) {
	_, err = w.WS.Call(w.Prefix+methods.ValidateAddress, params, &result)
	return
}

func (w *WebSocket) ExtractKeyFromAddress(params ExtractKeyFromAddressParams) (key interface{}, err error) {
	_, err = w.WS.Call(w.Prefix+methods.ExtractKeyFromAddress, params, &key)
	return
}

func (w *WebSocket) GetMinerWork(params GetMinerWorkParams) (result GetMinerWorkResult, err error) {
	_, err = w.WS.Call(w.Prefix+methods.GetMinerWork, params, &result)
	return
}

func (w *WebSocket) SplitAddress(params SplitAddressParams) (result SplitAddressResult, err error) {
	_, err = w.WS.Call(w.Prefix+methods.SplitAddress, params, &result)
	return
}

func (w *WebSocket) GetHardForks() (result []HardFork, err error) {
	_, err = w.WS.Call(w.Prefix+methods.GetHardForks, nil, &result)
	return
}

func (w *WebSocket) GetEstimatedFeeRates() (result []FeeRatesEstimated, err error) {
	_, err = w.WS.Call(w.Prefix+methods.GetEstimatedFeeRates, nil, &result)
	return
}

func (w *WebSocket) GetPrunedTopoheight() (result uint64, err error) {
	_, err = w.WS.Call(w.Prefix+methods.GetPrunedTopoheight, nil, &result)
	return
}

func (w *WebSocket) GetTransactionExecutor(params GetTransactionExecutorParams) (result GetTransactionExecutorResult, err error) {
	_, err = w.WS.Call(w.Prefix+methods.GetTransactionExecutor, params, &result)
	return
}

func (w *WebSocket) HasMultisigAtTopoheight(params HasMultisigAtTopoheightParams) (result bool, err error) {
	_, err = w.WS.Call(w.Prefix+methods.HasMultisigAtTopoheight, params, &result)
	return
}

func (w *WebSocket) GetMultisigAtTopoheight(params GetMultisigAtTopoheightParams) (result GetMultisigAtTopoheightResult, err error) {
	_, err = w.WS.Call(w.Prefix+methods.GetMultisigAtTopoheight, params, &result)
	return
}

func (w *WebSocket) GetMultisig(params GetMultisigParams) (result GetMultisigResult, err error) {
	_, err = w.WS.Call(w.Prefix+methods.GetMultisig, params, &result)
	return
}

func (w *WebSocket) HasMultisig(params HasMultisigParams) (result bool, err error) {
	_, err = w.WS.Call(w.Prefix+methods.HasMultisig, params, &result)
	return
}

func (w *WebSocket) GetContractOutputs(params GetContractOutputsParams) (result []ContractOutput, err error) {
	var outputs []interface{}
	_, err = w.WS.Call(w.Prefix+methods.GetContractOutputs, params, &outputs)
	result = parseContractOutputs(outputs)
	return
}

func (w *WebSocket) GetContractModule(params GetContractModuleParams) (result GetContractModuleResult, err error) {
	_, err = w.WS.Call(w.Prefix+methods.GetContractModule, params, &result)
	return
}

func (w *WebSocket) GetContractData(params GetContractDataParams) (result interface{}, err error) {
	_, err = w.WS.Call(w.Prefix+methods.GetContractData, params, &result)
	return
}

func (w *WebSocket) GetContractDataAtTopoheight(params GetContractDataAtTopoheightParams) (result interface{}, err error) {
	_, err = w.WS.Call(w.Prefix+methods.GetContractDataAtTopoheight, params, &result)
	return
}

func (w *WebSocket) GetContractBalance(params GetContractBalanceParams) (result interface{}, err error) {
	_, err = w.WS.Call(w.Prefix+methods.GetContractBalance, params, &result)
	return
}

func (w *WebSocket) GetContractBalanceAtTopoheight(params GetContractBalanceAtTopoheightParams) (result interface{}, err error) {
	_, err = w.WS.Call(w.Prefix+methods.GetContractBalanceAtTopoheight, params, &result)
	return
}

func (w *WebSocket) CountContracts() (result uint64, err error) {
	_, err = w.WS.Call(w.Prefix+methods.CountContracts, nil, &result)
	return
}

func (w *WebSocket) MakeIntegratedAddress(params MakeIntegratedAddressParams) (result string, err error) {
	_, err = w.WS.Call(w.Prefix+methods.MakeIntegratedAddress, params, &result)
	return
}

func (w *WebSocket) DecryptExtraData(params DecryptExtraDataParams) (result interface{}, err error) {
	_, err = w.WS.Call(w.Prefix+methods.DecryptExtraData, params, &result)
	return
}
