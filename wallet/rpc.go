package wallet

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	netUrl "net/url"

	"github.com/creachadair/jrpc2"
	"github.com/creachadair/jrpc2/jhttp"
)

type RPC struct {
	ctx    context.Context
	Client *jrpc2.Client
}

type AuthTransport struct {
	Transport http.RoundTripper
	Username  string
	Password  string
}

func (t *AuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	setAuthHeader(req.Header, t.Username, t.Password)
	return t.Transport.RoundTrip(req)
}

func setAuthHeader(header http.Header, username string, password string) {
	auth := fmt.Sprintf("%s:%s", username, password)
	buf := bytes.Buffer{}
	encoder := base64.NewEncoder(base64.StdEncoding, &buf)
	encoder.Write([]byte(auth))
	encoder.Close()

	header.Set("Authorization", fmt.Sprintf("Basic %s", buf.String()))
}

func NewRPC(ctx context.Context, url string, username string, password string) (*RPC, error) {
	daemonUrl, err := netUrl.Parse(url)
	if err != nil {
		return nil, err
	}

	channel := jhttp.NewChannel(daemonUrl.String(), &jhttp.ChannelOptions{
		Client: &http.Client{
			Transport: &AuthTransport{
				Transport: http.DefaultTransport,
				Username:  username,
				Password:  password,
			},
		},
	})
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

func (d *RPC) GetNetwork() (network string, err error) {
	err = d.Client.CallResult(d.ctx, string(GetNetwork), nil, &network)
	return
}

func (d *RPC) GetNonce() (nonce uint64, err error) {
	err = d.Client.CallResult(d.ctx, string(GetNonce), nil, &nonce)
	return
}

func (d *RPC) GetTopoheight() (topoheight uint64, err error) {
	err = d.Client.CallResult(d.ctx, string(GetTopoheight), nil, &topoheight)
	return
}

func (d *RPC) GetAddress(params GetAddressParams) (address string, err error) {
	err = d.Client.CallResult(d.ctx, string(GetAddress), params, &address)
	return
}

func (d *RPC) SplitAddress(params SplitAddressParams) (result SplitAddressResult, err error) {
	err = d.Client.CallResult(d.ctx, string(SplitAddress), params, &result)
	return
}

func (d *RPC) Rescan(params RescanParams) (success bool, err error) {
	err = d.Client.CallResult(d.ctx, string(Rescan), params, &success)
	return
}

func (d *RPC) GetBalance(params GetBalanceParams) (balance uint64, err error) {
	err = d.Client.CallResult(d.ctx, string(GetBalance), params, &balance)
	return
}

func (d *RPC) HasBalance(params GetBalanceParams) (exists bool, err error) {
	err = d.Client.CallResult(d.ctx, string(HasBalance), params, &exists)
	return
}

func (d *RPC) GetTrackedAssets() (assets []string, err error) {
	err = d.Client.CallResult(d.ctx, string(GetTrackedAssets), nil, &assets)
	return
}

func (d *RPC) GetAssetPrecision(params GetAssetPrecisionParams) (decimals int, err error) {
	err = d.Client.CallResult(d.ctx, string(GetAssetPrecision), params, &decimals)
	return
}

func (d *RPC) GetTransaction(params GetTransactionParams) (transaction TransactionEntry, err error) {
	err = d.Client.CallResult(d.ctx, string(GetTransaction), params, &transaction)
	return
}

func (d *RPC) BuildTransaction(params BuildTransactionParams) (result BuildTransactionResult, err error) {
	if err = checkFeeBuilder(params.Fee); err != nil {
		return
	}

	err = d.Client.CallResult(d.ctx, string(BuildTransaction), params, &result)
	return
}

func (d *RPC) ListTransactions(params ListTransactionsParams) (txs []TransactionEntry, err error) {
	err = d.Client.CallResult(d.ctx, string(ListTransactions), params, &txs)
	return
}

func (d *RPC) IsOnline() (online bool, err error) {
	err = d.Client.CallResult(d.ctx, string(IsOnline), nil, &online)
	return
}

func (d *RPC) SetOnlineMode() (success bool, err error) {
	err = d.Client.CallResult(d.ctx, string(SetOnlineMode), nil, &success)
	return
}

func (d *RPC) SetOfflineMode() (success bool, err error) {
	err = d.Client.CallResult(d.ctx, string(SetOfflineMode), nil, &success)
	return
}

func (d *RPC) SignData(data interface{}) (signature string, err error) {
	err = d.Client.CallResult(d.ctx, string(SignData), data, &signature)
	return
}

func (d *RPC) EstimateFees(params EstimateFeesParams) (amount uint64, err error) {
	err = d.Client.CallResult(d.ctx, string(EstimateFees), params, &amount)
	return
}

func checkFeeBuilder(fee *FeeBuilder) error {
	if fee != nil && fee.Multiplier != nil && fee.Value != nil {
		return fmt.Errorf("you cannot set both Multiplier and Value in FeeBuilder")
	}

	return nil
}
