package xelis

import (
	"context"
	netUrl "net/url"

	"github.com/creachadair/jrpc2"
	"github.com/creachadair/jrpc2/jhttp"
)

type DaemonRPC struct {
	Client *jrpc2.Client
}

func NewDaemonRPC(url string) (*DaemonRPC, error) {
	daemonUrl, err := netUrl.Parse(url)
	if err != nil {
		return nil, err
	}

	channel := jhttp.NewChannel(daemonUrl.String(), nil)
	rpcClient := jrpc2.NewClient(channel, nil)
	daemon := &DaemonRPC{
		Client: rpcClient,
	}

	return daemon, nil
}

type GetInfoResult struct {
	BlockTimeTarget uint64 `json:"block_time_target"`
	Difficulty      uint64 `json:"difficulty"`
	Height          uint64 `json:"height"`
	MempoolSize     uint64 `json:"mempool_size"`
	NativeSupply    uint64 `json:"native_supply"`
	StableHeight    uint64 `json:"stableheight"`
	TopHash         string `json:"top_hash"`
	Version         string `json:"version"`
	Network         string `json:"network"`
	TopoHeight      uint64 `json:"topoheight"`
}

func (d *DaemonRPC) GetInfo(ctx context.Context) (GetInfoResult, error) {
	var result GetInfoResult
	err := d.Client.CallResult(ctx, "get_info", nil, &result)
	return result, err
}

func (d *DaemonRPC) GetTopoHeight(ctx context.Context) (uint64, error) {
	var topoHeight uint64
	err := d.Client.CallResult(ctx, "get_topoheight", nil, &topoHeight)
	return topoHeight, err
}
