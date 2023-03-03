package daemon

import (
	"encoding/json"
	"fmt"
	netUrl "net/url"
	"time"

	"github.com/gorilla/websocket"
)

type WebSocket struct {
	Conn *websocket.Conn
}

func createRPCRequest(method string, params map[string]interface{}) ([]byte, int64, error) {
	id := time.Now().UnixMicro()
	rpcRequest := RPCRequest{ID: id, JSONRPC: "2.0", Method: method, Params: params}

	msg, err := json.Marshal(rpcRequest)
	return []byte(msg), id, err
}

func NewWebSocket(url string) (*WebSocket, error) {
	daemonUrl, err := netUrl.Parse(url)
	if err != nil {
		return nil, err
	}

	conn, _, err := websocket.DefaultDialer.Dial(daemonUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	ws := &WebSocket{
		Conn: conn,
	}

	return ws, nil
}

func (w *WebSocket) Close() {
	w.Conn.Close()
}

func (w *WebSocket) Listen(onData func(int, []byte, error), done chan struct{}) {
	go func() {
		select {
		case <-done:
			return
		default:
			for {
				msgType, msg, err := w.Conn.ReadMessage()
				onData(msgType, msg, err)
			}
		}
	}()
}

func (w *WebSocket) Subscribe(event string, onData func(RPCResponse, error), done chan struct{}) error {
	subscribeResponse, err := w.Call("subscribe", map[string]interface{}{
		"notify": event,
	})
	if err != nil {
		return err
	}

	w.Listen(func(msgType int, msg []byte, err error) {
		if err != nil {
			onData(RPCResponse{}, err)
		} else if msgType == websocket.TextMessage {
			var rpcResponse RPCResponse
			err := json.Unmarshal(msg, &rpcResponse)

			if subscribeResponse.ID == rpcResponse.ID {
				onData(rpcResponse, err)
			}
		}
	}, done)

	return err
}

func (w *WebSocket) NewBlock(onData func(NewBlockResult, error), done chan struct{}) error {
	return w.Subscribe("NewBlock", func(response RPCResponse, err error) {
		var result NewBlockResult
		if err != nil {
			onData(result, err)
		}

		data, err := json.Marshal(response.Result)
		if err != nil {
			onData(result, err)
			return
		}

		err = json.Unmarshal(data, &result)
		onData(result, err)
	}, done)
}

func (w *WebSocket) Call(method string, params map[string]interface{}) (response RPCResponse, err error) {
	msg, id, err := createRPCRequest(method, params)
	if err != nil {
		return response, err
	}

	err = w.Conn.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		return response, err
	}

	for {
		select {
		case <-time.After(5 * time.Second):
			return response, fmt.Errorf("timeout waiting for response")
		default:
			msgType, msg, err := w.Conn.ReadMessage()
			if err != nil {
				return response, err
			}

			if msgType == websocket.TextMessage {
				err := json.Unmarshal(msg, &response)
				if err != nil {
					return response, err
				}

				if response.ID == id {
					return response, nil
				}
			}

			time.Sleep(10 * time.Millisecond)
		}
	}
}

func (w *WebSocket) CallGetInfo() (result GetInfoResult, err error) {
	response, err := w.Call("get_info", nil)
	if err != nil {
		return result, err
	}

	data, err := json.Marshal(response.Result)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(data, &result)
	return result, err
}
