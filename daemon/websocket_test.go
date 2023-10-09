package daemon

import (
	"context"
	"testing"
	"time"

	"github.com/xelis-project/xelis-go-sdk/config"
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

	err := daemon.GetInfo(func(info *GetInfoResult, res *RPCResponse, err error) {
		t.Logf("%+v", info)
		daemon.cancel()
	})

	if err != nil {
		t.Fatal(err)
	}

	err = daemon.HandleListeners()
	if err != nil {
		t.Fatal(err)
	}

	daemon.conn.Close()
}

func TestWSNewBlock(t *testing.T) {
	daemon := setupWebSocket(t)

	var closeListener func() error
	closeListener, err := daemon.OnNewBlock(func(newBlock *NewBlockResult, r *RPCResponse) {
		t.Logf("%+v", newBlock)
		err := closeListener() // testing closeListener but not needed if we simply close the connection
		if err != nil {
			t.Fatal(err)
		}

		daemon.cancel()
	})

	if err != nil {
		t.Fatal(err)
	}

	err = daemon.HandleListeners()
	if err != nil {
		t.Fatal(err)
	}

	daemon.conn.Close()
}

func TestRoutineHandleListener(t *testing.T) {
	daemon := setupWebSocket(t)

	_, err := daemon.OnNewBlock(func(newBlock *NewBlockResult, r *RPCResponse) {
		t.Logf("%+v", newBlock)
	})
	if err != nil {
		t.Fatal(err)
	}

	go daemon.HandleListeners()
	time.Sleep(10 * time.Second)
	daemon.Close()
}

func TestWSMultiSubcribe(t *testing.T) {
	daemon := setupWebSocket(t)

	count := 0
	tryClose := func() {
		count++
		if count > 2 {
			daemon.cancel()
		}
	}

	_, err := daemon.OnListenEvent(NewBlock, func(res *RPCResponse) {
		t.Logf("%+v", res)
		tryClose()
	})
	if err != nil {
		t.Fatal(err)
	}

	_, err = daemon.OnListenEvent(NewBlock, func(res *RPCResponse) {
		t.Logf("%+v", res)
		tryClose()
	})
	if err != nil {
		t.Fatal(err)
	}

	err = daemon.GetInfo(func(info *GetInfoResult, res *RPCResponse, err error) {
		t.Logf("%+v", info)
		tryClose()
	})
	if err != nil {
		t.Fatal(err)
	}

	err = daemon.HandleListeners()
	if err != nil {
		t.Fatal(err)
	}

	daemon.conn.Close()
}
