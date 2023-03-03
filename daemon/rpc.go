package daemon

import (
	"context"
	netUrl "net/url"

	"github.com/creachadair/jrpc2"
	"github.com/creachadair/jrpc2/jhttp"
)

type RPC struct {
	Client *jrpc2.Client
}

func NewRPC(url string) (*RPC, error) {
	daemonUrl, err := netUrl.Parse(url)
	if err != nil {
		return nil, err
	}

	channel := jhttp.NewChannel(daemonUrl.String(), nil)
	rpcClient := jrpc2.NewClient(channel, nil)
	daemon := &RPC{
		Client: rpcClient,
	}

	return daemon, nil
}

func (d *RPC) GetInfo(ctx context.Context) (GetInfoResult, error) {
	var result GetInfoResult
	err := d.Client.CallResult(ctx, "get_info", nil, &result)
	return result, err
}

func (d *RPC) GetTopoHeight(ctx context.Context) (uint64, error) {
	var topoHeight uint64
	err := d.Client.CallResult(ctx, "get_topoheight", nil, &topoHeight)
	return topoHeight, err
}
