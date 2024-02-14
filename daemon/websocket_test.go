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

	info, err := daemon.GetInfo()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", info)

	fees, err := daemon.GetDevFeeThresholds()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", fees)

	size, err := daemon.GetSizeOnDisk()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", size)

	daemon.Close()
}

func TestWSNewBlock(t *testing.T) {
	daemon := setupWebSocket(t)
	var wg sync.WaitGroup
	wg.Add(1)
	err := daemon.NewBlockFunc(func(newBlock Block, err error) {
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

	err := daemon.NewBlockFunc(func(block Block, err error) {
		t.Logf("%+v", block)
	})

	if err != nil {
		t.Fatal(err)
	}

	err = daemon.CloseEvent(NewBlock)
	if err != nil {
		t.Fatal(err)
	}

	daemon.Close()
}

func TestWSCallAndMultiSubscribe(t *testing.T) {
	daemon := setupWebSocket(t)
	var wg sync.WaitGroup

	wg.Add(1)
	err := daemon.ws.ListenEventFunc(NewBlock, func(res lib.RPCResponse) {
		t.Logf("%+v", res)
		wg.Done()
	})
	if err != nil {
		t.Fatal(err)
	}

	wg.Add(1)
	err = daemon.ws.ListenEventFunc(NewBlock, func(res lib.RPCResponse) {
		t.Logf("%+v", res)
		wg.Done()
	})
	if err != nil {
		t.Fatal(err)
	}

	info, err := daemon.GetInfo()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", info)

	wg.Wait()
	daemon.Close()
}

func TestWSPeers(t *testing.T) {
	daemon := setupWebSocket(t)
	var wg sync.WaitGroup

	wg.Add(1)
	daemon.PeerConnectedFunc(func(p Peer, err error) {
		t.Logf("%+v", p)
		wg.Done()
	})

	wg.Add(1)
	daemon.PeerDisconnectedFunc(func(id uint64, err error) {
		t.Logf("%d", id)
		wg.Done()
	})

	wg.Wait()
	daemon.Close()
}

func TestWSPeerUpdated(t *testing.T) {
	daemon := setupWebSocket(t)
	var wg sync.WaitGroup

	wg.Add(1)
	daemon.PeerStateUpdatedFunc(func(p Peer, err error) {
		t.Logf("%+v", p)
		wg.Done()
	})

	wg.Wait()
	daemon.Close()
}
