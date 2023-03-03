package xelis

import (
	"encoding/json"
	"fmt"
	netUrl "net/url"
	"time"

	"github.com/gorilla/websocket"
)

type DaemonWS struct {
	Conn *websocket.Conn
}

type RPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type RPCRequest struct {
	ID      int64                  `json:"id"`
	JSONRPC string                 `json:"jsonrpc"`
	Method  string                 `json:"method"`
	Params  map[string]interface{} `json:"params,omitempty"`
}

type RPCResponse struct {
	ID     int64       `json:"id"`
	Result interface{} `json:"result,omitempty"`
	Error  RPCError    `json:"error,omitempty"`
}

func createRPCRequest(method string, params map[string]interface{}) ([]byte, int64, error) {
	id := time.Now().UnixMicro()
	rpcRequest := RPCRequest{ID: id, JSONRPC: "2.0", Method: method, Params: params}

	msg, err := json.Marshal(rpcRequest)
	return []byte(msg), id, err
}

func NewDaemonWS(url string) (*DaemonWS, error) {
	daemonUrl, err := netUrl.Parse(url)
	if err != nil {
		return nil, err
	}

	conn, _, err := websocket.DefaultDialer.Dial(daemonUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	daemonWS := &DaemonWS{
		Conn: conn,
	}

	return daemonWS, nil
}

func (d *DaemonWS) Close() {
	d.Conn.Close()
}

func (d *DaemonWS) Listen(onData func(int, []byte, error), done chan struct{}) {
	go func() {
		select {
		case <-done:
			return
		default:
			for {
				msgType, msg, err := d.Conn.ReadMessage()
				onData(msgType, msg, err)
			}
		}
	}()
}

func (d *DaemonWS) Subscribe(event string, onData func(RPCResponse, error), done chan struct{}) error {
	subscribeResponse, err := d.Call("subscribe", map[string]interface{}{
		"notify": event,
	})
	if err != nil {
		return err
	}

	d.Listen(func(msgType int, msg []byte, err error) {
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

type NewBlockResult struct {
	BlockType            string   `json:"block_type"`
	CumulativeDifficulty uint64   `json:"cumulative_difficulty"`
	Difficulty           uint64   `json:"difficulty"`
	Event                string   `json:"event"`
	ExtraNonce           string   `json:"extra_nonce"`
	Hash                 string   `json:"hash"`
	Height               uint64   `json:"height"`
	Miner                string   `json:"miner"`
	Nonce                uint64   `json:"nonce"`
	Reward               uint64   `json:"reward"`
	Supply               uint64   `json:"supply"`
	Timestamp            uint64   `json:"timestamp"`
	Tips                 []string `json:"tips"`
	Topoheight           uint64   `json:"topoheight"`
	TotalFees            uint64   `json:"total_fees"`
	TotalSizeInBytes     uint64   `json:"total_size_in_bytes"`
	TxsHashes            []string `json:"txs_hashes"`
}

func (d *DaemonWS) NewBlock(onData func(NewBlockResult, error), done chan struct{}) error {
	return d.Subscribe("NewBlock", func(response RPCResponse, err error) {
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

func (d *DaemonWS) Call(method string, params map[string]interface{}) (response RPCResponse, err error) {
	msg, id, err := createRPCRequest(method, params)
	if err != nil {
		return response, err
	}

	err = d.Conn.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		return response, err
	}

	for {
		select {
		case <-time.After(5 * time.Second):
			return response, fmt.Errorf("timeout waiting for response")
		default:
			msgType, msg, err := d.Conn.ReadMessage()
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

func (d *DaemonWS) CallGetInfo() (result GetInfoResult, err error) {
	response, err := d.Call("get_info", nil)
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
