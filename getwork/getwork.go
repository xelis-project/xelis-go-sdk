package getwork

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/gorilla/websocket"
)

type Getwork struct {
	conn *websocket.Conn

	Job           chan BlockTemplate
	AcceptedBlock chan bool
	RejectedBlock chan string
	Err           chan error
}

func NewGetwork(endpoint, minerAddress, worker string) (*Getwork, error) {
	socketUrl, err := url.Parse(fmt.Sprintf("%s/%s/%s", endpoint, minerAddress, worker))
	if err != nil {
		return nil, err
	}

	conn, _, err := websocket.DefaultDialer.Dial(socketUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	getwork := &Getwork{
		conn:          conn,
		Job:           make(chan BlockTemplate, 1),
		AcceptedBlock: make(chan bool, 1),
		RejectedBlock: make(chan string, 1),
		Err:           make(chan error, 1),
	}

	go func() {
		defer getwork.Close()

		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				getwork.Err <- err
				return
			}

			getwork.handleMessage(msg)
		}
	}()

	return getwork, nil
}

func (w *Getwork) Close() {
	close(w.Job)
	close(w.AcceptedBlock)
	close(w.RejectedBlock)
	close(w.Err)
	w.conn.Close()
}

func (w *Getwork) handleMessage(msg []byte) {
	var res interface{}
	err := json.Unmarshal(msg, &res)
	if err != nil {
		w.Err <- err
		return
	}

	jsonMap, ok := res.(map[string]interface{})
	if ok {
		blockTemplate, ok := jsonMap[NewJob].(map[string]interface{})
		if ok {
			w.Job <- BlockTemplate{
				Difficulty: blockTemplate["difficulty"].(string),
				Height:     uint64(blockTemplate["height"].(float64)),
				Template:   blockTemplate["template"].(string),
			}
			return
		}

		rejected, ok := jsonMap[BlockRejected].(string)
		if ok {
			w.RejectedBlock <- rejected
			return
		}
	}

	value, ok := res.(string)
	if ok {
		if value == BlockAccepted {
			w.AcceptedBlock <- true
			return
		}
	}
}

func (w *Getwork) SubmitBlock(hexData string) (err error) {
	data := map[string]interface{}{"block_template": hexData}
	return w.conn.WriteJSON(data)
}
