package daemon

import (
	"context"
	"sync"
	"testing"

	"github.com/xelis-project/xelis-go-sdk/config"
	"github.com/xelis-project/xelis-go-sdk/lib"
)

func setupWebSocket(t *testing.T) (daemon *WebSocket) {
	ctx := context.Background()
	daemon, err := NewWebSocket(ctx, config.TESTNET_NODE_WS)
	if err != nil {
		t.Fatal(err)
	}

	return
}

func TestWSGetInfo(t *testing.T) {
	daemon := setupWebSocket(t)

	info, _, err := daemon.GetInfo()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", info)
	daemon.Close()
}

func TestWSNewBlock(t *testing.T) {
	daemon := setupWebSocket(t)
	var wg sync.WaitGroup
	wg.Add(1)
	_, err := daemon.NewBlockFunc(func(newBlock NewBlockResult, res lib.RPCResponse) {
		t.Logf("%+v", newBlock)
		wg.Done()
	})

	if err != nil {
		t.Fatal(err)
	}

	wg.Wait()
	daemon.ws.Close()
}

func TestWSUnsubscribe(t *testing.T) {
	daemon := setupWebSocket(t)

	closeEvent, err := daemon.NewBlockFunc(func(block NewBlockResult, res lib.RPCResponse) {
		t.Logf("%+v", res)
	})

	if err != nil {
		t.Fatal(err)
	}

	err = closeEvent()
	if err != nil {
		t.Fatal(err)
	}

	daemon.Close()
}

func TestWSCallAndMultiSubscribe(t *testing.T) {
	daemon := setupWebSocket(t)
	var wg sync.WaitGroup

	wg.Add(1)
	_, err := daemon.ws.ListenEventFunc(NewBlock, func(res lib.RPCResponse) {
		t.Logf("%+v", res)
		wg.Done()
	})
	if err != nil {
		t.Fatal(err)
	}

	wg.Add(1)
	_, err = daemon.ws.ListenEventFunc(NewBlock, func(res lib.RPCResponse) {
		t.Logf("%+v", res)
		wg.Done()
	})
	if err != nil {
		t.Fatal(err)
	}

	info, _, err := daemon.GetInfo()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", info)

	wg.Wait()
	daemon.Close()
}
