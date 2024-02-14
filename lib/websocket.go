package lib

import (
	"context"
	"encoding/json"
	"fmt"
	netUrl "net/url"
	"time"

	"github.com/gorilla/websocket"
)

type WebSocket struct {
	ctx      context.Context
	cancel   context.CancelFunc
	conn     *websocket.Conn
	channels map[int64]chan RPCResponse
	events   map[RPCEvent]int64
	id       int64
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
		ctx:      ctx,
		cancel:   cancel,
		conn:     conn,
		channels: make(map[int64]chan RPCResponse),
		events:   make(map[RPCEvent]int64),
	}

	go ws.listen()
	return ws, nil
}

func (w *WebSocket) listen() {
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
					channel, ok := w.channels[rpcResponse.ID]
					if ok {
						channel <- rpcResponse
					}
				}
			}
		}
	}()
}

func (w *WebSocket) subscribeEvent(event RPCEvent) (RPCResponse, error) {
	return w.Call("subscribe", map[string]interface{}{
		"notify": event,
	})
}

func (w *WebSocket) unsubscribeEvent(event RPCEvent) (RPCResponse, error) {
	return w.Call("unsubscribe", map[string]interface{}{
		"notify": event,
	})
}

func (w *WebSocket) Close() error {
	// Remove channels and events.
	// We don't need to send unsubscribe event if we just close the connection.
	for id := range w.channels {
		ch := w.channels[id]
		close(ch)
		delete(w.channels, id)
	}

	for event := range w.events {
		delete(w.events, event)
	}

	w.cancel()
	return w.conn.Close()
}

func (w *WebSocket) ListenEventFunc(event RPCEvent, onData func(RPCResponse)) (closeEvent func() error, err error) {
	id, ok := w.events[event]
	if !ok {
		var res RPCResponse
		res, err = w.subscribeEvent(event)
		if err != nil {
			return
		}

		if res.Error != nil {
			err = fmt.Errorf(res.Error.Message)
			return
		}

		id = res.ID
		w.events[event] = id
	}

	ch := w.channels[id] // We don't need to make a channel because subscribeEvent already created it. It receives the ws messages on the same id.

	go func() {
		for res := range ch {
			onData(res)
		}
	}()

	closeEvent = func() error {
		_, ok := w.events[event]
		if ok {
			res, err := w.unsubscribeEvent(event)

			if err != nil {
				return err
			}

			if res.Error != nil {
				return fmt.Errorf(res.Error.Message)
			}

			close(ch)
			delete(w.channels, id)
			delete(w.events, event)
		}

		return nil
	}

	return
}

func (w *WebSocket) Call(method RPCMethod, params map[string]interface{}) (RPCResponse, error) {
	var res RPCResponse
	var err error

	w.id++
	rpcRequest := RPCRequest{ID: w.id, JSONRPC: "2.0", Method: method, Params: params}
	msg, err := json.Marshal(rpcRequest)
	if err != nil {
		return res, err
	}

	ch := make(chan RPCResponse)
	w.channels[w.id] = ch

	timer := time.AfterFunc(10*time.Second, func() {
		close(ch)
		delete(w.channels, w.id)
		err = fmt.Errorf("timeout waiting for response")
	})

	err = w.conn.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		return res, err
	}

	res = <-ch
	timer.Stop()
	return res, err
}
