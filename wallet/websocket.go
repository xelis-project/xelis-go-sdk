package wallet

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/xelis-project/xelis-go-sdk/daemon"
	"github.com/xelis-project/xelis-go-sdk/lib"
)

type WebSocket struct {
	ws *lib.WebSocket
}

func NewWebSocket(ctx context.Context, url string, username string, password string) (*WebSocket, error) {
	header := make(http.Header)
	setAuthHeader(header, username, password)
	ws, err := lib.NewWebSocket(ctx, url, header)
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

func (w *WebSocket) GetNetwork() (network string, err error) {
	res, err := w.ws.Call(GetNetwork, nil)
	if err != nil {
		return
	}

	if res.Error != nil {
		err = fmt.Errorf(res.Error.Message)
		return
	}

	err = json.Unmarshal(res.Result, &network)
	return
}

func (w *WebSocket) GetNonce() (nonce uint64, err error) {
	res, err := w.ws.Call(GetNonce, nil)
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

func (w *WebSocket) GetTopoheight() (topoheight uint64, err error) {
	res, err := w.ws.Call(GetTopoheight, nil)
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

func (w *WebSocket) GetAddress(params GetAddressParams) (address uint64, err error) {
	res, err := w.ws.Call(GetAddress, params)
	if err != nil {
		return
	}

	if res.Error != nil {
		err = fmt.Errorf(res.Error.Message)
		return
	}

	err = json.Unmarshal(res.Result, &address)
	return
}

func (w *WebSocket) SplitAddress(params SplitAddressParams) (result SplitAddressResult, err error) {
	res, err := w.ws.Call(SplitAddress, params)
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

func (w *WebSocket) Rescan(params RescanParams) (success bool, err error) {
	res, err := w.ws.Call(Rescan, params)
	if err != nil {
		return
	}

	if res.Error != nil {
		err = fmt.Errorf(res.Error.Message)
		return
	}

	err = json.Unmarshal(res.Result, &success)
	return
}

func (w *WebSocket) GetBalance(params GetBalanceParams) (balance uint64, err error) {
	res, err := w.ws.Call(Rescan, params)
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

func (w *WebSocket) GetTrackedAssets() (assets []string, err error) {
	res, err := w.ws.Call(GetTrackedAssets, nil)
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

func (w *WebSocket) GetAssetPrecision(params GetAssetPrecisionParams) (decimals int, err error) {
	res, err := w.ws.Call(GetAssetPrecision, nil)
	if err != nil {
		return
	}

	if res.Error != nil {
		err = fmt.Errorf(res.Error.Message)
		return
	}

	err = json.Unmarshal(res.Result, &decimals)
	return
}

func (w *WebSocket) GetTransaction(params GetTransactionParams) (transaction daemon.Transaction, err error) {
	res, err := w.ws.Call(GetTransaction, params)
	if err != nil {
		return
	}

	if res.Error != nil {
		err = fmt.Errorf(res.Error.Message)
		return
	}

	err = json.Unmarshal(res.Result, &transaction)
	return
}

func (w *WebSocket) BuildTransaction(params BuildTransactionParams) (result BuildTransactionResult, err error) {
	res, err := w.ws.Call(BuildTransaction, params)
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

func (w *WebSocket) ListTransactions(params ListTransactionsParams) (txs []daemon.Transaction, err error) {
	res, err := w.ws.Call(ListTransactions, params)
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

func (w *WebSocket) IsOnline() (online bool, err error) {
	res, err := w.ws.Call(IsOnline, nil)
	if err != nil {
		return
	}

	if res.Error != nil {
		err = fmt.Errorf(res.Error.Message)
		return
	}

	err = json.Unmarshal(res.Result, &online)
	return
}

func (w *WebSocket) SignData(data interface{}) (signature string, err error) {
	res, err := w.ws.Call(SignData, data)
	if err != nil {
		return
	}

	if res.Error != nil {
		err = fmt.Errorf(res.Error.Message)
		return
	}

	err = json.Unmarshal(res.Result, &signature)
	return
}

func (w *WebSocket) EstimateFees(params EstimateFeesParams) (amount uint64, err error) {
	res, err := w.ws.Call(EstimateFees, params)
	if err != nil {
		return
	}

	if res.Error != nil {
		err = fmt.Errorf(res.Error.Message)
		return
	}

	err = json.Unmarshal(res.Result, &amount)
	return
}
