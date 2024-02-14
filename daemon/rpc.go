package daemon

import (
	"context"
	netUrl "net/url"

	"github.com/creachadair/jrpc2"
	"github.com/creachadair/jrpc2/jhttp"
)

type RPC struct {
	ctx    context.Context
	Client *jrpc2.Client
}

func NewRPC(ctx context.Context, url string) (*RPC, error) {
	daemonUrl, err := netUrl.Parse(url)
	if err != nil {
		return nil, err
	}

	channel := jhttp.NewChannel(daemonUrl.String(), nil)
	rpcClient := jrpc2.NewClient(channel, nil)

	daemon := &RPC{
		ctx:    ctx,
		Client: rpcClient,
	}

	return daemon, nil
}

func (d *RPC) GetVersion() (version string, err error) {
	err = d.Client.CallResult(d.ctx, string(GetVersion), nil, &version)
	return
}

func (d *RPC) GetInfo() (result GetInfoResult, err error) {
	err = d.Client.CallResult(d.ctx, string(GetInfo), nil, &result)
	return
}

func (d *RPC) GetHeight() (height uint64, err error) {
	err = d.Client.CallResult(d.ctx, string(GetHeight), nil, &height)
	return
}

func (d *RPC) GetTopoheight() (topoheight uint64, err error) {
	err = d.Client.CallResult(d.ctx, string(GetTopoHeight), nil, &topoheight)
	return
}

func (d *RPC) GetStableheight() (stableheight uint64, err error) {
	err = d.Client.CallResult(d.ctx, string(GetStableHeight), nil, &stableheight)
	return
}

func (d *RPC) GetBlockTemplate(addr string) (result GetBlockTemplateResult, err error) {
	params := map[string]string{"address": addr}
	err = d.Client.CallResult(d.ctx, string(GetBlockTemplate), params, &result)
	return
}

func (d *RPC) GetBlockAtTopoheight(params GetBlockAtTopoheightParams) (block Block, err error) {
	err = d.Client.CallResult(d.ctx, string(GetBlockAtTopoheight), params, &block)
	return
}

func (d *RPC) GetBlocksAtHeight(params GetBlocksAtHeightParams) (blocks []Block, err error) {
	err = d.Client.CallResult(d.ctx, string(GetBlocksAtHeight), params, &blocks)
	return
}

func (d *RPC) GetBlockByHash(params GetBlockByHashParams) (block Block, err error) {
	err = d.Client.CallResult(d.ctx, string(GetBlockByHash), params, &block)
	return
}

func (d *RPC) GetTopBlock(params GetTopBlockParams) (block Block, err error) {
	err = d.Client.CallResult(d.ctx, string(GetTopBlock), params, &block)
	return
}

func (d *RPC) GetNonce(addr string) (nonce GetNonceResult, err error) {
	params := map[string]string{"address": addr}
	err = d.Client.CallResult(d.ctx, string(GetNonce), params, &nonce)
	return
}

func (d *RPC) HasNonce(addr string) (hasNonce bool, err error) {
	params := map[string]string{"address": addr}
	var result map[string]bool
	err = d.Client.CallResult(d.ctx, string(HasNonce), params, &result)
	hasNonce = result["exist"]
	return
}

func (d *RPC) GetBalance(params GetBalanceParams) (balance GetBalanceResult, err error) {
	err = d.Client.CallResult(d.ctx, string(GetBalance), params, &balance)
	return
}

func (d *RPC) HasBalance(params GetBalanceParams) (hasBalance bool, err error) {
	var result map[string]bool
	err = d.Client.CallResult(d.ctx, string(HasBalance), params, &result)
	hasBalance = result["exists"]
	return
}

func (d *RPC) GetBalanceAtTopoheight(params GetBalanceAtTopoheightParams) (balance Balance, err error) {
	err = d.Client.CallResult(d.ctx, string(GetBalanceAtTopoheight), params, &balance)
	return
}

func (d *RPC) GetAsset(assetId string) (asset Asset, err error) {
	params := map[string]string{"asset": assetId}
	err = d.Client.CallResult(d.ctx, string(GetAsset), params, &asset)
	return
}

func (d *RPC) GetAssets(params GetAssetsParams) (assets []string, err error) {
	err = d.Client.CallResult(d.ctx, string(GetAssets), params, &assets)
	return
}

func (d *RPC) CountAssets() (count uint64, err error) {
	err = d.Client.CallResult(d.ctx, string(CountAssets), nil, &count)
	return
}

func (d *RPC) CountTransactions() (count uint64, err error) {
	err = d.Client.CallResult(d.ctx, string(CountTransactions), nil, &count)
	return
}

func (d *RPC) CountAccounts() (count uint64, err error) {
	err = d.Client.CallResult(d.ctx, string(CountAccounts), nil, &count)
	return
}

func (d *RPC) GetTips() (tips []string, err error) {
	err = d.Client.CallResult(d.ctx, string(GetTips), nil, &tips)
	return
}

func (d *RPC) P2PStatus() (status P2PStatusResult, err error) {
	err = d.Client.CallResult(d.ctx, string(P2PStatus), nil, &status)
	return
}

func (d *RPC) GetDAGOrder(params GetTopoheightRangeParams) (hashes []string, err error) {
	err = d.Client.CallResult(d.ctx, string(GetDAGOrder), params, &hashes)
	return
}

func (d *RPC) SubmitBlock(blockTemplate string) (result bool, err error) {
	params := map[string]string{"block_template": blockTemplate}
	err = d.Client.CallResult(d.ctx, string(SubmitBlock), params, &result)
	return
}

func (d *RPC) SubmitTransaction(data string) (result bool, err error) {
	params := map[string]string{"data": data}
	err = d.Client.CallResult(d.ctx, string(SubmitTransaction), params, &result)
	return
}

func (d *RPC) GetMempool() (txs []Transaction, err error) {
	err = d.Client.CallResult(d.ctx, string(GetMempool), nil, &txs)
	return
}

func (d *RPC) GetTransaction(hash string) (tx Transaction, err error) {
	params := map[string]string{"hash": hash}
	err = d.Client.CallResult(d.ctx, string(GetTransaction), params, &tx)
	return
}

func (d *RPC) GetTransactions(params GetTransactionsParams) (txs []Transaction, err error) {
	err = d.Client.CallResult(d.ctx, string(GetTransactions), params, &txs)
	return
}

func (d *RPC) GetBlocksRangeByTopoheight(params GetTopoheightRangeParams) (blocks []Block, err error) {
	err = d.Client.CallResult(d.ctx, string(GetBlocksRangeByTopoheight), params, &blocks)
	return
}

func (d *RPC) GetBlocksRangeByHeight(params GetHeightRangeParams) (blocks []Block, err error) {
	err = d.Client.CallResult(d.ctx, string(GetBlocksRangeByHeight), params, &blocks)
	return
}

func (d *RPC) GetAccounts(params GetAccountsParams) (addresses []string, err error) {
	err = d.Client.CallResult(d.ctx, string(GetAccounts), params, &addresses)
	return
}

func (d *RPC) GetAccountHistory(addr string) (history AccountHistory, err error) {
	params := map[string]string{"address": addr}
	err = d.Client.CallResult(d.ctx, string(GetAccountHistory), params, &history)
	return
}

func (d *RPC) GetAccountAssets(addr string) (assets []string, err error) {
	params := map[string]string{"address": addr}
	err = d.Client.CallResult(d.ctx, string(GetAccountAssets), params, &assets)
	return
}

func (d *RPC) GetPeers() (peers []Peer, err error) {
	err = d.Client.CallResult(d.ctx, string(GetPeers), nil, &peers)
	return
}

func (d *RPC) GetDevFeeThresholds() (fees []Fee, err error) {
	err = d.Client.CallResult(d.ctx, string(GetDevFeeThresholds), nil, &fees)
	return
}

func (d *RPC) GetSizeOnDisk() (sizeOnDisk SizeOnDisk, err error) {
	err = d.Client.CallResult(d.ctx, string(GetSizeOnDisk), nil, &sizeOnDisk)
	return
}

func (d *RPC) IsTxExecutedInBlock(params IsTxExecutedInBlockParams) (executed bool, err error) {
	err = d.Client.CallResult(d.ctx, string(IsTxExecutedInBlock), params, &executed)
	return
}
