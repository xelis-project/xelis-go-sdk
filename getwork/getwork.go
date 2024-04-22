package getwork

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/xelis-project/xelis-go-sdk/lib"
)

type Getwork struct {
	WS *lib.WebSocket

	Jobs           chan BlockTemplate
	AcceptedBlocks chan string
	RejectedBlocks chan string
}

func NewGetwork(endpoint, minerAddress, worker string) (*Getwork, error) {
	if len(endpoint) < 1 {
		return nil, errors.New("invalid endpoint")
	}

	ws, err := lib.NewWebSocket(endpoint+"getwork/"+minerAddress+"/"+worker, nil)
	if err != nil {
		return nil, err
	}

	jobs := make(chan BlockTemplate, 1)
	acceptedBlocks := make(chan string, 1)
	rejectedBlocks := make(chan string, 1)

	go func() {
		for {
			msg := <-ws.Notifications

			if msg == nil {
				fmt.Println("msg is nil")
				continue
			}

			fmt.Println("received a message from websocket:", string(msg))

			var rpcResponse map[string]json.RawMessage
			err = json.Unmarshal(msg, &rpcResponse)
			if err != nil {
				fmt.Println(err)
				return
			}

			for i, v := range rpcResponse {
				switch i {
				case NewJob:
					data := BlockTemplate{}
					err := json.Unmarshal(v, &data)
					if err != nil {
						fmt.Println(err)
						return
					}

					jobs <- data

				case BlockAccepted:

					acceptedBlocks <- string(v)

				case BlockRejected:

					rejectedBlocks <- string(v)
				}
			}
		}
	}()

	return &Getwork{
		WS:             ws,
		Jobs:           jobs,
		AcceptedBlocks: acceptedBlocks,
		RejectedBlocks: rejectedBlocks,
	}, nil
}

func (w *Getwork) Close() error {
	return w.WS.Close()
}

func (w *Getwork) CloseEvent(event string) error {
	return w.WS.CloseEvent(event)
}

func (w *Getwork) Submit(blockminer string) error {
	d := map[string]any{}
	d[SubmitBlock] = blockminer

	err := w.WS.Write(d)
	return err
}

func (w *Getwork) SubmitBlock(data string) (err error) {
	return w.WS.GetConn().WriteJSON(map[string]any{
		"block_template": data,
	})
}
