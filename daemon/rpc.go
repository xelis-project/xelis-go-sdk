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

func (d *RPC) GetInfo(ctx context.Context) (*GetInfoResult, error) {
	var result GetInfoResult
	err := d.Client.CallResult(ctx, string(GetInfo), nil, &result)
	return &result, err
}

func (d *RPC) GetTopoHeight(ctx context.Context) (uint64, error) {
	var topoHeight uint64
	err := d.Client.CallResult(ctx, string(GetTopoHeight), nil, &topoHeight)
	return topoHeight, err
}

func (d *RPC) GetStableHeight(ctx context.Context) (uint64, error) {
	var stableHeight uint64
	err := d.Client.CallResult(ctx, string(GetStableHeight), nil, &stableHeight)
	return stableHeight, err
}

func (d *RPC) GetBlocks(ctx context.Context, params *GetRangeParams) ([]Block, error) {
	var blocks []Block
	err := d.Client.CallResult(ctx, string(GetBlocks), params, &blocks)
	return blocks, err
}

func (d *RPC) GetTransactions(ctx context.Context, params *GetTransactionsParams) ([]Transaction, error) {
	var txs []Transaction
	err := d.Client.CallResult(ctx, string(GetTransactions), params, &txs)
	return txs, err
}
