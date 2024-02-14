package xswd

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/xelis-project/xelis-go-sdk/daemon"
	"github.com/xelis-project/xelis-go-sdk/lib"
	"github.com/xelis-project/xelis-go-sdk/wallet"
)

type Permission int

var (
	Ask          Permission = 0
	AcceptAlways Permission = 1
	DenyAlways   Permission = 2
)

type ApplicationData struct {
	ID          string                `json:"id"`
	Name        string                `json:"name"`
	Description string                `json:"description"`
	Url         string                `json:"url,omitempty"`
	Permissions map[string]Permission `json:"permissions"`
	Signature   string                `json:"signature,omitempty"`
}

type XSWD struct {
	WS     *lib.WebSocket
	Daemon *daemon.WebSocket
	Wallet *wallet.WebSocket
}

func NewXSWD(ctx context.Context, endpoint string) (*XSWD, error) {
	ws, err := lib.NewWebSocket(ctx, endpoint, nil)
	if err != nil {
		return nil, err
	}

	ws.CallTimeout = 0 // Not timeout, because we have to wait for user input.

	daemon := &daemon.WebSocket{
		Prefix: "node.",
		WS:     ws,
	}

	wallet := &wallet.WebSocket{
		Prefix: "wallet.",
		WS:     ws,
	}

	return &XSWD{
		WS:     ws,
		Daemon: daemon,
		Wallet: wallet,
	}, nil
}

func (x *XSWD) Close() error {
	return x.WS.Close()
}

func (x *XSWD) Authorize(app ApplicationData) (res lib.RPCResponse, err error) {
	data, err := json.Marshal(app)
	if err != nil {
		return
	}

	// id of 0 is reserved and not use in Call().
	res, err = x.WS.RawCall(0, data)
	if err != nil {
		return
	}

	if res.Error != nil {
		err = fmt.Errorf(res.Error.Message)
		return
	}

	return
}
