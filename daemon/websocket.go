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

func createRPCRequest(method RPCMethod, params map[string]interface{}) ([]byte, int64, error) {
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

func (w *WebSocket) Subscribe(event RPCEvent, onData func(*RPCResponse, error), done chan struct{}) error {
	subRes, err := w.Call("subscribe", map[string]interface{}{
		"notify": event,
	})
	if err != nil {
		return err
	}

	if subRes.Error != nil {
		return fmt.Errorf(subRes.Error.Message)
	}

	w.Listen(func(msgType int, msg []byte, err error) {
		if err != nil {
			onData(nil, err)
		} else if msgType == websocket.TextMessage {
			var rpcResponse RPCResponse
			err := json.Unmarshal(msg, &rpcResponse)
			if err != nil {
				onData(nil, err)
			} else if subRes.ID == rpcResponse.ID {
				onData(&rpcResponse, nil)
			}
		}
	}, done)

	return err
}

func (w *WebSocket) OnNewBlock(onData func(*NewBlockResult, error), done chan struct{}) error {
	return w.Subscribe(NewBlock, func(response *RPCResponse, err error) {
		var result NewBlockResult
		if err != nil {
			onData(nil, err)
		}

		err = json.Unmarshal(response.Result, &result)
		onData(&result, err)
	}, done)
}

func (w *WebSocket) Call(method RPCMethod, params map[string]interface{}) (*RPCResponse, error) {
	msg, id, err := createRPCRequest(method, params)
	if err != nil {
		return nil, err
	}

	err = w.Conn.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		return nil, err
	}

	for {
		select {
		case <-time.After(5 * time.Second):
			return nil, fmt.Errorf("timeout waiting for response")
		default:
			msgType, msg, err := w.Conn.ReadMessage()
			if err != nil {
				return nil, err
			}

			if msgType == websocket.TextMessage {
				var response RPCResponse
				err := json.Unmarshal(msg, &response)
				if err != nil {
					return nil, err
				}

				if response.ID == id {
					return &response, nil
				}
			}

			time.Sleep(10 * time.Millisecond)
		}
	}
}

func (w *WebSocket) GetInfo() (*GetInfoResult, error) {
	response, err := w.Call(GetInfo, nil)
	if err != nil {
		return nil, err
	}

	var result GetInfoResult
	err = json.Unmarshal(response.Result, &result)
	return &result, err
}
