package daemon

import (
	"context"
	"encoding/json"

	"github.com/xelis-project/xelis-go-sdk/lib"
)

type WebSocket struct {
	ws *lib.WebSocket
}

func NewWebSocket(ctx context.Context, url string) (*WebSocket, error) {
	ws, err := lib.NewWebSocket(ctx, url)
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

func (w *WebSocket) NewBlockFunc(onData func(NewBlockResult, lib.RPCResponse)) (func() error, error) {
	return w.ws.ListenEventFunc(NewBlock, func(res lib.RPCResponse) {
		var result NewBlockResult
		json.Unmarshal(res.Result, &result)
		onData(result, res)
	})
}

func (w *WebSocket) GetInfo() (GetInfoResult, lib.RPCResponse, error) {
	var result GetInfoResult

	res, err := w.ws.Call(GetInfo, nil)
	if err != nil {
		return result, res, err
	}

	err = json.Unmarshal(res.Result, &result)
	return result, res, err
}
