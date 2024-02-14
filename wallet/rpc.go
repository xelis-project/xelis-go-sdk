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
	"github.com/xelis-project/xelis-go-sdk/daemon"
)

type RPC struct {
	ctx    context.Context
	Client *jrpc2.Client
}

type HeaderTransport struct {
	Transport http.RoundTripper
	Headers   map[string]string
}

func (t *HeaderTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	for key, value := range t.Headers {
		req.Header.Set(key, value)
	}
	return t.Transport.RoundTrip(req)
}

func NewRPC(ctx context.Context, url string, username string, password string) (*RPC, error) {
	daemonUrl, err := netUrl.Parse(url)
	if err != nil {
		return nil, err
	}

	headers := make(map[string]string)
	auth := fmt.Sprintf("%s:%s", username, password)

	buf := bytes.Buffer{}
	encoder := base64.NewEncoder(base64.StdEncoding, &buf)
	encoder.Write([]byte(auth))
	encoder.Close()

	headers["Authorization"] = fmt.Sprintf("Basic %s", buf.String())

	channel := jhttp.NewChannel(daemonUrl.String(), &jhttp.ChannelOptions{
		Client: &http.Client{
			Transport: &HeaderTransport{
				Transport: http.DefaultTransport,
				Headers:   headers,
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
	err = d.Client.CallResult(d.ctx, string(daemon.GetVersion), nil, &version)
	return
}
