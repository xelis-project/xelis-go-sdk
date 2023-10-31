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

func (d *RPC) GetVersion(ctx context.Context) (version string, err error) {
	err = d.Client.CallResult(ctx, string(GetVersion), nil, &version)
	return
}

func (d *RPC) GetInfo(ctx context.Context) (result GetInfoResult, err error) {
	err = d.Client.CallResult(ctx, string(GetInfo), nil, &result)
	return
}

func (d *RPC) GetHeight(ctx context.Context) (height uint64, err error) {
	err = d.Client.CallResult(ctx, string(GetHeight), nil, &height)
	return
}

func (d *RPC) GetTopoheight(ctx context.Context) (topoheight uint64, err error) {
	err = d.Client.CallResult(ctx, string(GetTopoHeight), nil, &topoheight)
	return
}

func (d *RPC) GetStableheight(ctx context.Context) (stableheight uint64, err error) {
	err = d.Client.CallResult(ctx, string(GetStableHeight), nil, &stableheight)
	return
}

func (d *RPC) GetBlockTemplate(ctx context.Context, addr string) (result GetBlockTemplateResult, err error) {
	params := map[string]string{"address": addr}
	err = d.Client.CallResult(ctx, string(GetBlockTemplate), params, &result)
	return
}

func (d *RPC) GetBlockAtTopoheight(ctx context.Context, params GetBlockAtTopoheightParams) (block Block, err error) {
	err = d.Client.CallResult(ctx, string(GetBlockAtTopoheight), params, &block)
	return
}

func (d *RPC) GetBlocksAtHeight(ctx context.Context, params GetBlocksAtHeightParams) (blocks []Block, err error) {
	err = d.Client.CallResult(ctx, string(GetBlocksAtHeight), params, &blocks)
	return
}

func (d *RPC) GetBlockByHash(ctx context.Context, params GetBlockByHashParams) (block Block, err error) {
	err = d.Client.CallResult(ctx, string(GetBlockByHash), params, &block)
	return
}

func (d *RPC) GetTopBlock(ctx context.Context, params GetTopBlockParams) (block Block, err error) {
	err = d.Client.CallResult(ctx, string(GetTopBlock), params, &block)
	return
}

func (d *RPC) GetNonce(ctx context.Context, addr string) (nonce uint64, err error) {
	params := map[string]string{"address": addr}
	err = d.Client.CallResult(ctx, string(GetNonce), params, &nonce)
	return
}

func (d *RPC) HasNonce(ctx context.Context, addr string) (hasNonce bool, err error) {
	params := map[string]string{"address": addr}
	var result map[string]bool
	err = d.Client.CallResult(ctx, string(HasNonce), params, &result)
	hasNonce = result["exist"]
	return
}

func (d *RPC) GetLastBalance(ctx context.Context, params GetLastBalanceParams) (balance GetLastBalanceResult, err error) {
	err = d.Client.CallResult(ctx, string(GetLastBalance), params, &balance)
	return
}

func (d *RPC) GetBalanceAtTopoheight(ctx context.Context, params GetBalanceAtTopoheightParams) (balance Balance, err error) {
	err = d.Client.CallResult(ctx, string(GetBalanceAtTopoheight), params, &balance)
	return
}

func (d *RPC) GetAssets(ctx context.Context, params GetAssetsParams) (assets []string, err error) {
	err = d.Client.CallResult(ctx, string(GetAssets), params, &assets)
	return
}

func (d *RPC) CountAssets(ctx context.Context) (count uint64, err error) {
	err = d.Client.CallResult(ctx, string(CountAssets), nil, &count)
	return
}

func (d *RPC) CountTransactions(ctx context.Context) (count uint64, err error) {
	err = d.Client.CallResult(ctx, string(CountTransactions), nil, &count)
	return
}

func (d *RPC) CountAccounts(ctx context.Context) (count uint64, err error) {
	err = d.Client.CallResult(ctx, string(CountAccounts), nil, &count)
	return
}

func (d *RPC) GetMempool(ctx context.Context) (txs []Transaction, err error) {
	err = d.Client.CallResult(ctx, string(GetMempool), nil, &txs)
	return
}

func (d *RPC) GetTips(ctx context.Context) (tips []string, err error) {
	err = d.Client.CallResult(ctx, string(GetTips), nil, &tips)
	return
}

func (d *RPC) P2PStatus(ctx context.Context) (status P2PStatusResult, err error) {
	err = d.Client.CallResult(ctx, string(P2PStatus), nil, &status)
	return
}

func (d *RPC) GetBlocksRangeByTopoheight(ctx context.Context, params GetTopoHeightRangeParams) (blocks []Block, err error) {
	err = d.Client.CallResult(ctx, string(GetBlocksRangeByTopoheight), params, &blocks)
	return
}

func (d *RPC) GetBlocksRangeByHeight(ctx context.Context, params GetHeightRangeParams) (blocks []Block, err error) {
	err = d.Client.CallResult(ctx, string(GetBlocksRangeByHeight), params, &blocks)
	return
}

func (d *RPC) GetTransactions(ctx context.Context, params GetTransactionsParams) (txs []Transaction, err error) {
	err = d.Client.CallResult(ctx, string(GetTransactions), params, &txs)
	return
}

func (d *RPC) GetTransaction(ctx context.Context, hash string) (txs []Transaction, err error) {
	params := map[string]string{"hash": hash}
	err = d.Client.CallResult(ctx, string(GetTransaction), params, &txs)
	return
}

func (d *RPC) GetDAGOrder(ctx context.Context, params GetTopoHeightRangeParams) (hashes []string, err error) {
	err = d.Client.CallResult(ctx, string(GetDAGOrder), params, &hashes)
	return
}

func (d *RPC) SubmitBlock(ctx context.Context, blockTemplate string) (result bool, err error) {
	params := map[string]string{"block_template": blockTemplate}
	err = d.Client.CallResult(ctx, string(SubmitBlock), params, &result)
	return
}

func (d *RPC) SubmitTransaction(ctx context.Context, data string) (result bool, err error) {
	params := map[string]string{"data": data}
	err = d.Client.CallResult(ctx, string(SubmitTransaction), params, &result)
	return
}

func (d *RPC) GetAccounts(ctx context.Context, params GetAccountsParams) (addresses []string, err error) {
	err = d.Client.CallResult(ctx, string(GetAccounts), params, &addresses)
	return
}

func (d *RPC) GetAccountHistory(ctx context.Context, addr string) (history AccountHistory, err error) {
	params := map[string]string{"address": addr}
	err = d.Client.CallResult(ctx, string(GetAccountHistory), params, &history)
	return
}

func (d *RPC) GetAccountAssets(ctx context.Context, addr string) (assets []string, err error) {
	params := map[string]string{"address": addr}
	err = d.Client.CallResult(ctx, string(GetAccountAssets), params, &assets)
	return
}

func (d *RPC) GetPeers(ctx context.Context) (peers []Peer, err error) {
	err = d.Client.CallResult(ctx, string(GetPeers), nil, &peers)
	return
}
