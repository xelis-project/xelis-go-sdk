package wallet

import (
	"net/http"

	"github.com/xelis-project/xelis-go-sdk/daemon"
	"github.com/xelis-project/xelis-go-sdk/rpc"
)

type WebSocket struct {
	Prefix string
	WS     *rpc.WebSocket
}

func NewWebSocket(endpoint string, username string, password string) (*WebSocket, error) {
	header := make(http.Header)
	setAuthHeader(header, username, password)
	ws, err := rpc.NewWebSocket(endpoint, header)
	if err != nil {
		return nil, err
	}

	daemonWS := &WebSocket{
		WS: ws,
	}

	return daemonWS, nil
}

func (w *WebSocket) Close() error {
	return w.WS.Close()
}

func (w *WebSocket) CloseEvent(event string) error {
	return w.WS.CloseEvent(event)
}

func (w *WebSocket) NewTopoheightChannel() (chan uint64, chan error, error) {
	chanTopoheight := make(chan uint64)
	chanErr := make(chan error)

	err := w.WS.ListenEventFunc(NewTopoheight, func(res rpc.RPCResponse) {
		var result map[string]interface{}
		err := rpc.JsonFormatResponse(res, nil, &result)
		if err != nil {
			chanErr <- err
		} else {
			chanTopoheight <- uint64(result["topoheight"].(float64))
		}
	})

	return chanTopoheight, chanErr, err
}

func (w *WebSocket) NewTopoheightFunc(onData func(uint64, error)) error {
	return w.WS.ListenEventFunc(NewTopoheight, func(res rpc.RPCResponse) {
		var result map[string]interface{}
		err := rpc.JsonFormatResponse(res, nil, &result)
		topoheight := uint64(result["topoheight"].(float64))
		onData(topoheight, err)
	})
}

func (w *WebSocket) NewAssetChannel() (chan daemon.AssetWithData, chan error, error) {
	chanAssetWithData := make(chan daemon.AssetWithData)
	chanErr := make(chan error)

	err := w.WS.ListenEventFunc(NewAsset, func(res rpc.RPCResponse) {
		var result daemon.AssetWithData
		err := rpc.JsonFormatResponse(res, nil, &result)
		if err != nil {
			chanErr <- err
		} else {
			chanAssetWithData <- result
		}
	})

	return chanAssetWithData, chanErr, err
}

func (w *WebSocket) NewAssetFunc(onData func(daemon.AssetWithData, error)) error {
	return w.WS.ListenEventFunc(NewAsset, func(res rpc.RPCResponse) {
		var result daemon.AssetWithData
		err := rpc.JsonFormatResponse(res, nil, &result)
		onData(result, err)
	})
}

func (w *WebSocket) NewTransactionChannel() (chan TransactionEntry, chan error, error) {
	chanTransactionEntry := make(chan TransactionEntry)
	chanErr := make(chan error)

	err := w.WS.ListenEventFunc(NewTransaction, func(res rpc.RPCResponse) {
		var result TransactionEntry
		err := rpc.JsonFormatResponse(res, nil, &result)
		if err != nil {
			chanErr <- err
		} else {
			chanTransactionEntry <- result
		}
	})

	return chanTransactionEntry, chanErr, err
}

func (w *WebSocket) NewTransactionFunc(onData func(TransactionEntry, error)) error {
	return w.WS.ListenEventFunc(NewTransaction, func(res rpc.RPCResponse) {
		var result TransactionEntry
		err := rpc.JsonFormatResponse(res, nil, &result)
		onData(result, err)
	})
}

func (w *WebSocket) BalanceChangedChannel() (chan BalanceChangedResult, chan error, error) {
	chanBalanceChangedResult := make(chan BalanceChangedResult)
	chanErr := make(chan error)

	err := w.WS.ListenEventFunc(BalanceChanged, func(res rpc.RPCResponse) {
		var result BalanceChangedResult
		err := rpc.JsonFormatResponse(res, nil, &result)
		if err != nil {
			chanErr <- err
		} else {
			chanBalanceChangedResult <- result
		}
	})

	return chanBalanceChangedResult, chanErr, err
}

func (w *WebSocket) BalanceChangedFunc(onData func(BalanceChangedResult, error)) error {
	return w.WS.ListenEventFunc(BalanceChanged, func(res rpc.RPCResponse) {
		var result BalanceChangedResult
		err := rpc.JsonFormatResponse(res, nil, &result)
		onData(result, err)
	})
}

func (w *WebSocket) RescanChannel() (chan uint64, chan error, error) {
	chanStartTopoheight := make(chan uint64)
	chanErr := make(chan error)

	err := w.WS.ListenEventFunc(Rescan, func(res rpc.RPCResponse) {
		var result map[string]interface{}
		err := rpc.JsonFormatResponse(res, nil, &result)
		if err != nil {
			chanErr <- err
		} else {
			chanStartTopoheight <- uint64(result["start_topoheight"].(float64))
		}
	})

	return chanStartTopoheight, chanErr, err
}

func (w *WebSocket) RescanFunc(onData func(uint64, error)) error {
	return w.WS.ListenEventFunc(Rescan, func(res rpc.RPCResponse) {
		var result map[string]interface{}
		err := rpc.JsonFormatResponse(res, nil, &result)
		startTopoheight := uint64(result["start_topoheight"].(float64))
		onData(startTopoheight, err)
	})
}

func (w *WebSocket) OnlineChannel() (chan bool, chan error, error) {
	chanOnline := make(chan bool)
	chanErr := make(chan error)

	err := w.WS.ListenEventFunc(Online, func(res rpc.RPCResponse) {
		chanOnline <- true
	})

	return chanOnline, chanErr, err
}

func (w *WebSocket) OnlineFunc(onData func()) error {
	return w.WS.ListenEventFunc(Online, func(res rpc.RPCResponse) {
		onData()
	})
}

func (w *WebSocket) OfflineChannel() (chan bool, chan error, error) {
	chanOffline := make(chan bool)
	chanErr := make(chan error)

	err := w.WS.ListenEventFunc(Offline, func(res rpc.RPCResponse) {
		chanOffline <- true
	})

	return chanOffline, chanErr, err
}

func (w *WebSocket) OfflineFunc(onData func()) error {
	return w.WS.ListenEventFunc(Offline, func(res rpc.RPCResponse) {
		onData()
	})
}

func (w *WebSocket) GetVersion() (version string, err error) {
	res, err := w.WS.Call(w.Prefix+GetVersion, nil)
	err = rpc.JsonFormatResponse(res, err, &version)
	return
}

func (w *WebSocket) GetNetwork() (network string, err error) {
	res, err := w.WS.Call(w.Prefix+GetNetwork, nil)
	err = rpc.JsonFormatResponse(res, err, &network)
	return
}

func (w *WebSocket) GetNonce() (nonce uint64, err error) {
	res, err := w.WS.Call(w.Prefix+GetNonce, nil)
	err = rpc.JsonFormatResponse(res, err, &nonce)
	return
}

func (w *WebSocket) GetTopoheight() (topoheight uint64, err error) {
	res, err := w.WS.Call(w.Prefix+GetTopoheight, nil)
	err = rpc.JsonFormatResponse(res, err, &topoheight)
	return
}

func (w *WebSocket) GetAddress(params GetAddressParams) (address string, err error) {
	res, err := w.WS.Call(w.Prefix+GetAddress, params)
	err = rpc.JsonFormatResponse(res, err, &address)
	return
}

func (w *WebSocket) SplitAddress(params SplitAddressParams) (result SplitAddressResult, err error) {
	res, err := w.WS.Call(w.Prefix+SplitAddress, params)
	err = rpc.JsonFormatResponse(res, err, &result)
	return
}

func (w *WebSocket) Rescan(params RescanParams) (success bool, err error) {
	res, err := w.WS.Call(w.Prefix+Rescan, params)
	err = rpc.JsonFormatResponse(res, err, &success)
	return
}

func (w *WebSocket) GetBalance(params GetBalanceParams) (balance uint64, err error) {
	res, err := w.WS.Call(w.Prefix+Rescan, params)
	err = rpc.JsonFormatResponse(res, err, &balance)
	return
}

func (w *WebSocket) HasBalance(params GetBalanceParams) (exists bool, err error) {
	res, err := w.WS.Call(w.Prefix+HasBalance, params)
	err = rpc.JsonFormatResponse(res, err, &exists)
	return
}

func (w *WebSocket) GetTrackedAssets() (assets []string, err error) {
	res, err := w.WS.Call(w.Prefix+GetTrackedAssets, nil)
	err = rpc.JsonFormatResponse(res, err, &assets)
	return
}

func (w *WebSocket) GetAssetPrecision(params GetAssetPrecisionParams) (decimals int, err error) {
	res, err := w.WS.Call(w.Prefix+GetAssetPrecision, nil)
	err = rpc.JsonFormatResponse(res, err, &decimals)
	return
}

func (w *WebSocket) GetTransaction(params GetTransactionParams) (transaction TransactionEntry, err error) {
	res, err := w.WS.Call(w.Prefix+GetTransaction, params)
	err = rpc.JsonFormatResponse(res, err, &transaction)
	return
}

func (w *WebSocket) BuildTransaction(params BuildTransactionParams) (result BuildTransactionResult, err error) {
	res, err := w.WS.Call(w.Prefix+BuildTransaction, params)
	err = rpc.JsonFormatResponse(res, err, &result)
	return
}

func (w *WebSocket) ListTransactions(params ListTransactionsParams) (txs []TransactionEntry, err error) {
	res, err := w.WS.Call(w.Prefix+ListTransactions, params)
	err = rpc.JsonFormatResponse(res, err, &txs)
	return
}

func (w *WebSocket) IsOnline() (online bool, err error) {
	res, err := w.WS.Call(w.Prefix+IsOnline, nil)
	err = rpc.JsonFormatResponse(res, err, &online)
	return
}

func (w *WebSocket) SetOnlineMode() (success bool, err error) {
	res, err := w.WS.Call(w.Prefix+SetOnlineMode, nil)
	err = rpc.JsonFormatResponse(res, err, &success)
	return
}

func (w *WebSocket) SetOfflineMode() (success bool, err error) {
	res, err := w.WS.Call(w.Prefix+SetOfflineMode, nil)
	err = rpc.JsonFormatResponse(res, err, &success)
	return
}

func (w *WebSocket) SignData(data interface{}) (signature string, err error) {
	res, err := w.WS.Call(w.Prefix+SignData, data)
	err = rpc.JsonFormatResponse(res, err, &signature)
	return
}

func (w *WebSocket) EstimateFees(params EstimateFeesParams) (amount uint64, err error) {
	res, err := w.WS.Call(w.Prefix+EstimateFees, params)
	err = rpc.JsonFormatResponse(res, err, &amount)
	return
}
