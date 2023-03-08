package daemon

import (
	"context"
	"encoding/json"
	"fmt"
	netUrl "net/url"
	"reflect"
	"time"

	"github.com/gorilla/websocket"
)

type WebSocket struct {
	ctx       context.Context
	cancel    context.CancelFunc
	conn      *websocket.Conn
	listeners map[int64][]func(*RPCResponse)
	events    map[RPCEvent]int64
}

func funcUniqueId(f interface{}) string {
	v := reflect.ValueOf(f).Pointer()
	return fmt.Sprintf("%d", v)
}

func createRPCRequest(method RPCMethod, params map[string]interface{}) ([]byte, int64, error) {
	id := time.Now().Unix()
	rpcRequest := RPCRequest{ID: id, JSONRPC: "2.0", Method: method, Params: params}

	msg, err := json.Marshal(rpcRequest)
	return []byte(msg), id, err
}

func NewWebSocket(ctx context.Context, url string) (*WebSocket, error) {
	daemonUrl, err := netUrl.Parse(url)
	if err != nil {
		return nil, err
	}

	conn, _, err := websocket.DefaultDialer.Dial(daemonUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(ctx)
	ws := &WebSocket{
		ctx:       ctx,
		cancel:    cancel,
		conn:      conn,
		listeners: make(map[int64][]func(*RPCResponse)),
		events:    make(map[RPCEvent]int64),
	}

	go ws.readLoop()
	return ws, nil
}

func (w *WebSocket) readLoop() {
	go func() {
		select {
		case <-w.ctx.Done():
			return
		default:
			for {
				msgType, msg, err := w.conn.ReadMessage()
				if err != nil {
					w.cancel()
					return
				}

				if msgType == websocket.TextMessage {
					var rpcResponse RPCResponse
					json.Unmarshal(msg, &rpcResponse)
					for id, listeners := range w.listeners {
						if rpcResponse.ID == id {
							for _, listener := range listeners {
								go listener(&rpcResponse)
							}
						}
					}
				}
			}
		}
	}()
}

func (w *WebSocket) HandleListeners() error {
	for {
		select {
		case <-w.ctx.Done():
			return nil
		default:
			err := w.ctx.Err()
			if err != nil {
				return err
			}

			time.Sleep(time.Millisecond)
		}
	}
}

func (w *WebSocket) subscribeEvent(event RPCEvent) (int64, error) {
	wait := make(chan struct{})
	var resError error
	var id int64
	err := w.Call("subscribe", map[string]interface{}{
		"notify": event,
	}, func(res *RPCResponse, err error) {
		if err != nil {
			resError = err
		} else if res.Error != nil {
			resError = fmt.Errorf(res.Error.Message)
		} else {
			id = res.ID
		}

		close(wait)
	})
	if err != nil {
		return 0, err
	}

	<-wait
	if resError != nil {
		return 0, resError
	}

	return id, nil
}

func (w *WebSocket) unsubscribeEvent(event RPCEvent) error {
	wait := make(chan struct{})
	var resError error
	err := w.Call("unsubscribe", map[string]interface{}{
		"notify": event,
	}, func(res *RPCResponse, err error) {
		if err != nil {
			resError = err
		} else if res.Error != nil {
			resError = fmt.Errorf(res.Error.Message)
		}

		close(wait)
	})

	if err != nil {
		return err
	}
	<-wait
	if resError != nil {
		return resError
	}

	return nil
}

func (w *WebSocket) clearListeners() {
	for id := range w.listeners {
		delete(w.listeners, id)
	}

	for event := range w.events {
		delete(w.events, event)
	}
}

func (w *WebSocket) Close() {
	// just remove listeners and events
	// we don't need to send unsubscribe event if we just close the connection
	w.clearListeners()
	w.cancel()
	w.conn.Close()
}

func (w *WebSocket) OnListenEvent(event RPCEvent, onData func(*RPCResponse)) (func() error, error) {
	id, ok := w.events[event]
	if !ok {
		newId, err := w.subscribeEvent(event)
		if err != nil {
			return nil, err
		}

		id = newId
		w.events[event] = id
	}

	w.listeners[id] = append(w.listeners[id], onData)

	closeListen := func() error {
		if len(w.listeners[id]) == 1 {
			err := w.unsubscribeEvent(event)
			if err != nil {
				return err
			}
		}

		var newListeners []func(*RPCResponse)
		funcId := funcUniqueId(onData)
		for _, f := range w.listeners[id] {
			if funcUniqueId(f) != funcId {
				newListeners = append(newListeners, f)
			}
		}

		w.listeners[id] = newListeners
		if len(w.listeners[id]) == 0 {
			delete(w.listeners, id)
			delete(w.events, event)
		}

		return nil
	}

	return closeListen, nil
}

func (w *WebSocket) OnNewBlock(onData func(*NewBlockResult, *RPCResponse)) (func() error, error) {
	return w.OnListenEvent(NewBlock, func(res *RPCResponse) {
		var result NewBlockResult
		json.Unmarshal(res.Result, &result)
		onData(&result, res)
	})
}

func (w *WebSocket) Call(method RPCMethod, params map[string]interface{}, onData func(*RPCResponse, error)) error {
	msg, id, err := createRPCRequest(method, params)
	if err != nil {
		return err
	}

	timer := time.AfterFunc(10*time.Second, func() {
		delete(w.listeners, id)
		onData(nil, fmt.Errorf("timeout waiting for response"))
	})

	w.listeners[id] = append(w.listeners[id], func(res *RPCResponse) {
		timer.Stop()
		delete(w.listeners, id)
		onData(res, nil)
	})

	err = w.conn.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		return err
	}

	return nil
}

func (w *WebSocket) GetInfo(onData func(*GetInfoResult, *RPCResponse, error)) error {
	return w.Call(GetInfo, nil, func(res *RPCResponse, err error) {
		if err != nil {
			onData(nil, nil, err)
			return
		}

		var result GetInfoResult
		err = json.Unmarshal(res.Result, &result)
		onData(&result, res, err)
	})
}
