package wallet

import (
	"net/http"

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
